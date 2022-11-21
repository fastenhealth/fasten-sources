package models

import (
	"os"
)

//go:generate mockgen -source=source_client.go -destination=mock/mock_source_client.go
type SourceClient interface {
	GetUsCoreResources() []string
	GetRequest(resourceSubpath string, decodeModelPtr interface{}) error
	SyncAll(db DatabaseRepository) (UpsertSummary, error)
	SyncAllByResourceName(db DatabaseRepository, resourceNames []string) (UpsertSummary, error)

	//Manual client ONLY functions
	SyncAllBundle(db DatabaseRepository, bundleFile *os.File) (UpsertSummary, error)
}
