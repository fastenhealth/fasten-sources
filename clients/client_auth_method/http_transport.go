package client_auth_method

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

var debugMode = false

type debugLoggingTransport struct{}

func (t *debugLoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// modify req here
	//req.Header.Set("User-Agent", "oauth-client/0.0")
	//req.Header.Del("Authorization")

	// Dump request headers
	req_dump, _ := httputil.DumpRequest(req, true)
	fmt.Printf("REQ:\n%s\n", req_dump)

	// Call default rounttrip
	response, err := http.DefaultTransport.RoundTrip(req)

	// Dump response headers
	res_dump, _ := httputil.DumpResponse(response, true)
	fmt.Printf("RES:\n%s\n", res_dump)

	// return result of default roundtrip
	return response, err
}
