package models

import "context"

//go:generate mockgen -source=source_credential_repository.go -destination=mock/mock_source_credential_repository.go
type SourceCredentialRepository interface {
	// this is used to update tokens after refreshing (which can happen in multiple places within a source client)
	StoreTokens(ctx context.Context, sourceCredentials SourceCredential) error
}
