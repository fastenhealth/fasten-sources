package models

import (
	"context"
)

//go:generate mockgen -source=database_repository.go -destination=mock/mock_database_repository.go
type DatabaseRepository interface {
	WrapRawResource(rawResourcePtr interface{})
	UpsertRawResource(ctx context.Context, sourceCredentials SourceCredential, rawResource ResourceInterface) error
}
