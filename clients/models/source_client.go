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

	// RefreshAccessToken will refresh the Access Token using the OAuth Token endpoint
	// https://build.fhir.org/ig/HL7/smart-app-launch/example-app-launch-asymmetric-auth.html#refresh-access-token
	// https://build.fhir.org/ig/HL7/smart-app-launch/example-app-launch-symmetric-auth.html#refresh-access-token
	// Some EHRs implement Rolling Refresh Tokens, so the Refresh Token may also be updated during this process. You must
	// store the new Refresh Token if provided. (This will be done automatically using the SourceCredentialRepository.StoreTokens method)
	RefreshAccessToken(options ...func(*SourceClientRefreshOptions)) error

	// IntrospectToken will provide information about the token using the OAuth Introspection endpoint if available
	// https://build.fhir.org/ig/HL7/smart-app-launch/token-introspection.html
	// Generally this method is used to validate if a Token is still valid or has expired/revoked
	// Most EHR servers will require authentication via client credentials or JWT to access the introspection endpoint
	// Note: This function *may* require a valid AccessToken to be set on the client depending on the source implementation
	IntrospectToken(tokenType TokenIntrospectTokenType) (*TokenIntrospectResponse, error)
}
