package models

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	"os"
)

//go:generate mockgen -source=source_client.go -destination=mock/mock_source_client.go
type SourceClient interface {
	GetResourceTypesUsCore() []string
	GetResourceTypesAllowList() []string
	GetRequest(resourceSubpath string, decodeModelPtr interface{}) (string, error)
	GetResourceBundle(relativeResourcePath string) (interface{}, error)
	SyncAll(db StorageRepository) (UpsertSummary, error)
	SyncAllByResourceName(db StorageRepository, resourceNames []string) (UpsertSummary, error)
	SyncAllByPatientEverythingBundle(db StorageRepository, bundleModel interface{}) (UpsertSummary, error)

	//Manual client ONLY functions
	SyncAllBundle(db StorageRepository, bundleFile *os.File, bundleFhirVersion pkg.FhirVersion) (UpsertSummary, error)
	ExtractPatientId(bundleFile *os.File) (string, pkg.FhirVersion, error)

	GetSourceCredential() SourceCredential
	RefreshAccessToken(options ...func(*SourceClientRefreshOptions)) error

	IntrospectToken(tokenType TokenIntrospectTokenType) (*TokenIntrospectResponse, error)
}
