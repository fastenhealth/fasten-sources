package models

import (
	"context"
)

//go:generate mockgen -source=storage_repository.go -destination=mock/mock_storage_repository.go
type StorageRepository interface {
	UpsertRawResource(ctx context.Context, sourceCredentials SourceCredential, rawResource RawResourceFhir) (bool, error)
	UpsertRawResourceAssociation(
		ctx context.Context,
		sourceId string,
		sourceResourceType string,
		sourceResourceId string,
		targetSourceId string,
		targetResourceType string,
		targetResourceId string,
	) error
	BackgroundJobCheckpoint(ctx context.Context, checkpointData map[string]interface{}, errorData map[string]interface{})
}
