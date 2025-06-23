package models

import (
	"context"
	"net/http"

	"github.com/tink-crypto/tink-go/v2/keyset"
	"golang.org/x/oauth2"
)

type SourceClientOptions struct {
	TestHttpClient *http.Client //test http client, only used for testing, if not provided, default http client will be used
	//When test mode is enabled, tokens will not be refreshed, and Http client provided will be used (usually go-vcr for playback)
	TestMode bool

	//Client related options/overrides. Some of these will already be provided via the Endpoint or the SourceCredential.
	ClientID              string
	ClientSecret          string
	ClientJWTKeysetHandle *keyset.Handle //required for JWT Bearer Token refresh
	RedirectURL           string
	Scopes                []string

	ResourceTypesAllowList []string //list of resource types that are allowed to be fetched from this source. Default empty (USCDI Core List)

	SourceClientRefreshOptions []func(*SourceClientRefreshOptions)

	Context     context.Context
	Concurrency int
}

func WithHttpClient(customHttpClient *http.Client) func(*SourceClientOptions) {
	// we need to be able to override the http client (set Timeouts, or SOCKS Proxy, etc)
	// TODO: this function should be thouroughly tested. See first GH issue.
	// https://github.com/golang/oauth2/issues/324#issuecomment-1537546747
	// context.WithValue(context.Background(), oauth2.HTTPClient, cleanhttp.DefaultClient())
	// https://github.com/hashicorp/go-gcp-common/blob/main/gcputil/credentials.go#L172C9-L172C94
	return func(s *SourceClientOptions) {
		s.Context = context.WithValue(s.Context, oauth2.HTTPClient, customHttpClient)
	}
}

func WithTestHttpClient(httpClient *http.Client) func(*SourceClientOptions) {
	//TODO: should this be moved to WithHttpClient pattern?
	return func(s *SourceClientOptions) {
		s.TestHttpClient = httpClient
		s.TestMode = true
	}
}

func WithClientID(clientID string) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.ClientID = clientID
	}
}

func WithClientSecret(clientSecret string) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.ClientSecret = clientSecret
	}
}

func WithClientJWTKeysetHandle(clientJWTKeysetHandle *keyset.Handle) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.ClientJWTKeysetHandle = clientJWTKeysetHandle
	}
}

func WithRedirectURL(redirectURL string) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.RedirectURL = redirectURL
	}
}

func WithScopes(scopes []string) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.Scopes = scopes
	}
}

func WithResourceTypeAllowList(resourceTypeAllowList []string) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.ResourceTypesAllowList = resourceTypeAllowList
	}
}

func WithSourceClientRefreshOptions(options ...func(*SourceClientRefreshOptions)) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.SourceClientRefreshOptions = options
	}
}

func WithConcurrency(Concurrency int) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.Concurrency = Concurrency
	}
}
