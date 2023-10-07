package models

import (
	"context"
)

//go:generate mockgen -source=database_repository.go -destination=mock/mock_database_repository.go
type DatabaseRepository interface {
	UpsertRawResource(ctx context.Context, sourceCredentials SourceCredential, rawResource RawResourceFhir) (bool, error)
	BackgroundJobCheckpoint(ctx context.Context, checkpointData map[string]interface{}, errorData map[string]interface{})
}
