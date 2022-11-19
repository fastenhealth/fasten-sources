package models

//go:generate mockgen -source=source_credential.go -destination=mock/mock_source_credential.go
type SourceCredential interface {
	GetClientId() string
	GetPatientId() string
	GetOauthAuthorizationEndpoint() string
	GetOauthTokenEndpoint() string
	GetApiEndpointBaseUrl() string
	GetRefreshToken() string
	GetAccessToken() string
	GetExpiresAt() int64

	RefreshTokens(accessToken string, refreshTokens string, expiresAt int64)
}
