package models

import "github.com/fastenhealth/fasten-sources/pkg"

//go:generate mockgen -source=source_credential.go -destination=mock/mock_source_credential.go
//this is actually an interface to a pointer receiver
type SourceCredential interface {
	GetSourceType() pkg.SourceType
	GetClientId() string
	GetPatientId() string
	GetOauthAuthorizationEndpoint() string
	GetOauthTokenEndpoint() string
	GetApiEndpointBaseUrl() string
	GetRefreshToken() string
	GetAccessToken() string
	GetExpiresAt() int64

	SetTokens(accessToken string, refreshTokens string, expiresAt int64)
	IsDynamicClient() bool
	RefreshDynamicClientAccessToken() error
}
