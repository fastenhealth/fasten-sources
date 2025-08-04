package utils

import (
	"github.com/hashicorp/go-retryablehttp"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// parseRetryAfterHeader parses the Retry-After header and returns the
// delay duration according to the spec: https://httpwg.org/specs/rfc7231.html#header.retry-after
// The bool returned will be true if the header was successfully parsed.
// Otherwise, the header was either not present, or was not parseable according to the spec.
//
// Retry-After headers come in two flavors: Seconds or HTTP-Date
//
// Examples:
// * Retry-After: Fri, 31 Dec 1999 23:59:59 GMT
// * Retry-After: 120
// See: https://github.com/hashicorp/go-retryablehttp/blob/73a996c6390330f4a8fac5766d9fbcbe445591c6/client.go#L578
func parseRetryAfterHeader(headers []string) (time.Duration, bool) {
	if len(headers) == 0 || headers[0] == "" {
		return 0, false
	}
	header := headers[0]
	// Retry-After: 120
	if sleep, err := strconv.ParseInt(header, 10, 64); err == nil {
		if sleep < 0 { // a negative sleep doesn't make sense
			return 0, false
		}
		return time.Second * time.Duration(sleep), true
	}

	// Retry-After: Fri, 31 Dec 1999 23:59:59 GMT
	retryTime, err := time.Parse(time.RFC1123, header)
	if err != nil {
		return 0, false
	}
	if until := retryTime.Sub(time.Now()); until > 0 {
		return until, true
	}
	// date is in the past
	return 0, true
}

// Some APIs implement rate limiting using the X-RateLimit headers, not the `Retry-After` header.
// This function will parse the X-RateLimit headers and return a duration to wait before retrying the request.
// X-RateLimit-Limit-Minute: Number of allowed requests in the given minute
// X-RateLimit-Remaining: Number of requests remaining in the current window
// X-RateLimit-Reset: Time left until the current window resets
// X-RateLimit-Limit: Maximum number of requests per minute
// See:https://github.com/hashicorp/go-retryablehttp/blob/73a996c6390330f4a8fac5766d9fbcbe445591c6/client.go#L647
func XRateLimitLinearJitterBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			//print all rate limit headers for debugging
			for k, v := range resp.Header {
				if strings.HasPrefix(strings.ToLower(k), "x-ratelimit-") {
					log.Printf("X-RateLimit Header: %s: %s\n", k, v)
				}
			}

			//make sure we use the canonical header name
			if sleep, ok := parseRetryAfterHeader(resp.Header["X-Ratelimit-Reset"]); ok {
				return sleep
			}
		}
	}
	return retryablehttp.RateLimitLinearJitterBackoff(min, max, attemptNum, resp)
}
