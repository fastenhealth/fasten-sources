package models

import "net/http"

type SourceClientOptions struct {
	TestHttpClient *http.Client //test http client, only used for testing, if not provided, default http client will be used

	SourceClientRefreshOptions []func(*SourceClientRefreshOptions)
}

func WithTestHttpClient(httpClient *http.Client) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.TestHttpClient = httpClient
	}
}

func WithSourceClientRefreshOptions(options ...func(*SourceClientRefreshOptions)) func(*SourceClientOptions) {
	return func(s *SourceClientOptions) {
		s.SourceClientRefreshOptions = options
	}
}
