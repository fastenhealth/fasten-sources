package models

import "net/http"

type SourceClientOptions struct {
	TestHttpClient *http.Client //test http client, only used for testing, if not provided, default http client will be used
	//When test mode is enabled, tokens will not be refreshed, and Http client provided will be used (usually go-vcr for playback)
	TestMode bool

	//Client related options/overrides. Some of these will already be provided via the Endpoint or the SourceCredential.
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string

	SourceClientRefreshOptions []func(*SourceClientRefreshOptions)
}

func WithTestHttpClient(httpClient *http.Client) func(*SourceClientOptions) {
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

func WithSourceClientRefreshOptions(options ...func(*SourceClientRefreshOptions)) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.SourceClientRefreshOptions = options
	}
}
