package models

import (
	"github.com/fastenhealth/fasten-sources/pkg"
)

// this is actually an interface to a pointer receiver
//
//go:generate mockgen -source=source_credential.go -destination=mock/mock_source_credential.go
type SourceCredential interface {
	GetSourceId() string
	GetSourceCredentialType() pkg.SourceCredentialType

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
}
