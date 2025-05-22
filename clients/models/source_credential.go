package models

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	"net/http"
)

// this is actually an interface to a pointer receiver
//
//go:generate mockgen -source=source_credential.go -destination=mock/mock_source_credential.go
type SourceCredential interface {
	GetSourceId() string

	GetEndpointId() string
	GetPortalId() string
	GetBrandId() string
	GetPlatformType() pkg.PlatformType

	GetClientId() string
	GetPatientId() string
	GetRefreshToken() string
	GetAccessToken() string
	GetExpiresAt() int64

	SetTokens(accessToken string, refreshTokens string, expiresAt int64)

	//this is used to determine how we should refresh the access token (either using client token
	ClientAuthenticationMethodType() pkg.ClientAuthenticationMethodType
	RefreshPrivateKeyJwtToken(testHttpClient ...*http.Client) error
}
