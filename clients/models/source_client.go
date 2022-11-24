package models

import (
	"os"
)

//go:generate mockgen -source=source_client.go -destination=mock/mock_source_client.go
type SourceClient interface {
	GetUsCoreResources() []string
	GetRequest(resourceSubpath string, decodeModelPtr interface{}) error
	GetResourceBundle(relativeResourcePath string) (interface{}, error)
	SyncAll(db DatabaseRepository) (UpsertSummary, error)
	SyncAllByResourceName(db DatabaseRepository, resourceNames []string) (UpsertSummary, error)
	SyncAllByPatientEverythingBundle(db DatabaseRepository, bundleModel interface{}) (UpsertSummary, error)

	//Manual client ONLY functions
	SyncAllBundle(db DatabaseRepository, bundleFile *os.File, bundleType string) (UpsertSummary, error)
	ExtractPatientId(bundleFile *os.File) (string, string, error)
}
