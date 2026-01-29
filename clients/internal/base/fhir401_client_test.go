package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/fastenhealth/fasten-sources/clients/testutils"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"

	"github.com/fastenhealth/fasten-sources/clients/models"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/definitions"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir401"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

// RoundTripper helper for unit tests; lets us stub HTTP responses without network
type rtFunc func(*http.Request) (*http.Response, error)

func (f rtFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestNewFHIR401Client(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	sc.EXPECT().GetAccessToken().Return("test-access-token")
	sc.EXPECT().GetRefreshToken().Return("test-refresh-token")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	endpointDefinition := &definitionsModels.LighthouseSourceDefinition{}

	//test
	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, endpointDefinition, models.WithTestHttpClient(&http.Client{}))

	//assert
	require.NoError(t, err)
	require.Equal(t, client.SourceCredential.GetAccessToken(), "test-access-token")
	require.Equal(t, client.SourceCredential.GetRefreshToken(), "test-refresh-token")
}

func TestFHIR401Client_ProcessBundle_Cigna(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	//sc.EXPECT().GetAccessToken().Return("test-access-token")
	//sc.EXPECT().GetRefreshToken().Return("test-refresh-token")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cignaSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("6c0454af-1631-4c4d-905d-5710439df983"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cignaSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)

	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/bundle/cigna_syntheticuser05-everything.json")
	require.NoError(t, err)
	var bundle fhir401.Bundle
	err = json.Unmarshal(jsonBytes, &bundle)
	require.NoError(t, err)

	// test
	wrappedResourceModels, _, err := client.ProcessBundle(bundle)
	//log.Printf("%v", wrappedResourceModels)

	//assert
	require.NoError(t, err)
	require.Equal(t, 11, len(wrappedResourceModels))
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}

func TestFHIR401Client_ProcessBundle_Cerner(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	//sc.EXPECT().GetAccessToken().Return("test-access-token")
	//sc.EXPECT().GetRefreshToken().Return("test-refresh-token")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)

	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/bundle/cerner_open_12724067_DocumentReference.json")
	require.NoError(t, err)
	var bundle fhir401.Bundle
	err = json.Unmarshal(jsonBytes, &bundle)
	require.NoError(t, err)

	// test
	wrappedResourceModels, _, err := client.ProcessBundle(bundle)
	//log.Printf("%v", wrappedResourceModels)

	//assert
	require.NoError(t, err)
	require.Equal(t, 10, len(wrappedResourceModels))
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}

func TestFhir401Client_ProcessResource(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	db := mock_models.NewMockStorageRepository(mockCtrl)
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/document_reference/cerner_document_reference_206130480.json")
	require.NoError(t, err)
	referencedResourcesLookup := &sync.Map{}
	internalFragmentReferenceLookup := map[string]string{}
	summary := models.UpsertSummary{}

	rawResource := models.RawResourceFhir{
		SourceResourceID:   "206130480",
		SourceResourceType: "DocumentReference",
		ResourceRaw:        jsonBytes,
	}

	// test
	_, err = client.ProcessResource(db, rawResource, referencedResourcesLookup, internalFragmentReferenceLookup, &summary)
	referencedResourcesLookupMap := testutils.ToMap[string, bool](referencedResourcesLookup)
	//assert
	require.NoError(t, err)
	require.Equal(t, 8, len(referencedResourcesLookupMap))
	require.Equal(t, map[string]bool{
		"DiagnosticReport/206130480":  false,
		"DocumentReference/206130480": true,
		"Encounter/97953480":          false,
		"Organization/675844":         false,
		"Patient/12724067":            false,
		"Practitioner/12742069":       false,
		"https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XML-206130480": false,
		"https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-206130480":  false,
	}, referencedResourcesLookupMap)
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}

func TestFhir401Client_ProcessEncounterResource_WhichContainsCapitalizedStatusEnum_ShouldParseCorrectly(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	db := mock_models.NewMockStorageRepository(mockCtrl)
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/document_reference/encounter_resource_broken_parse.json")
	require.NoError(t, err)
	referencedResourcesLookup := &sync.Map{}
	internalFragmentReferenceLookup := map[string]string{}
	summary := models.UpsertSummary{}

	rawResource := models.RawResourceFhir{
		SourceResourceID:   "e7vTGSIu3VmPxDxo8hbtSKQ3",
		SourceResourceType: "Encounter",
		ResourceRaw:        jsonBytes,
	}

	// test
	_, err = client.ProcessResource(db, rawResource, referencedResourcesLookup, internalFragmentReferenceLookup, &summary)

	referencedResourcesLookupMap := testutils.ToMap[string, bool](referencedResourcesLookup)

	//assert
	require.NoError(t, err)
	require.Equal(t, 4, len(referencedResourcesLookupMap))
	require.Equal(t, map[string]bool{
		"Encounter/e7vTGSIu3VmPxDxo8hbtSKQ3":                      true,
		"Location/eq0Zb8k23Hg.GX9aMfavyZg3":                       false,
		"Patient/etMDDQeIicnivVsdRrZcEQGPqfFlXIbJH5wMZqQ7ZNldGo3": false,
		"Practitioner/e7y0GLCXTb0rdKZlNSuQhww3":                   false,
	}, referencedResourcesLookupMap)
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}

func TestFhir401Client_ProcessObservationResource_WhichContainsUnicodeCharacters_ShouldParseCorrectly(t *testing.T) {
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	db := mock_models.NewMockStorageRepository(mockCtrl)
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/document_reference/observation_broken_parse.json")
	require.NoError(t, err)
	referencedResourcesLookup := &sync.Map{}
	internalFragmentReferenceLookup := map[string]string{}
	summary := models.UpsertSummary{}

	rawResource := models.RawResourceFhir{
		SourceResourceID:   "e7vTGSIu3VmPxDxo8hbtSKQ3",
		SourceResourceType: "Encounter",
		ResourceRaw:        jsonBytes,
	}

	// test
	_, err = client.ProcessResource(db, rawResource, referencedResourcesLookup, internalFragmentReferenceLookup, &summary)
	referencedResourcesLookupMap := testutils.ToMap[string, bool](referencedResourcesLookup)

	//assert
	require.NoError(t, err)
	//require.Equal(t, 4, len(referencedResourcesLookupMap))
	require.Equal(t, map[string]bool{
		"Encounter/e7vTGSIu3VmPxDxo8hbtSKQ3":                  true,
		"Encounter/exxx3In3deguCac1YMmg3":                     false,
		"Patient/1234":                                        false,
		"ServiceRequest/klsdk3-PBKXHFeY0Q7KMDFz92TYK2nbYKsI3": false,
		"Specimen/xxxxwer-YwClO-82tNgDDI43tkGp2UsGcRs3":       false,
	}, referencedResourcesLookupMap)
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}

func TestFhir401Client_ProcessResourceWithContainedResources(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	db := mock_models.NewMockStorageRepository(mockCtrl)
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	medicareSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("6ae6c14e-b927-4ce0-862f-91123cb8d774"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, medicareSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/document_reference/medicare-eob.json")
	require.NoError(t, err)
	referencedResourcesLookup := &sync.Map{}
	internalFragmentReferenceLookup := map[string]string{}
	summary := models.UpsertSummary{}

	rawResource := models.RawResourceFhir{
		SourceResourceID:   "carrier--10000930037921",
		SourceResourceType: "ExplanationOfBenefit",
		ResourceRaw:        jsonBytes,
	}

	// test
	_, err = client.ProcessResource(db, rawResource, referencedResourcesLookup, internalFragmentReferenceLookup, &summary)
	referencedResourcesLookupMap := testutils.ToMap[string, bool](referencedResourcesLookup)

	//assert
	require.NoError(t, err)
	require.Equal(t, 23, len(referencedResourcesLookupMap))
	//notice how the contained resources are tagged as completed in the referencedResourcesLookup
	require.Equal(t, map[string]bool{
		"Coverage/part-b--10000010254618":              false,
		"ExplanationOfBenefit/carrier--10000930037921": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi00":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi01":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi02":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi03":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi04":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi05":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0x":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xMA==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xMQ==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xMg==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xMw==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xNA==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xNQ==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xNg==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xNw==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xOA==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0xOQ==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0y":     true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0yMA==": true,
		"Observation/RXhwbGFuYXRpb25PZkJlbmVmaXQvY2Fycmllci0tMTAwMDA5MzAwMzc5MjEjbGluZS1vYnNlcnZhdGlvbi0z":     true,
		"Patient/-10000010254618": false,
	}, referencedResourcesLookupMap)
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}

func TestFhir401Client_ProcessPendingResource_Limit5(t *testing.T) {
	t.Parallel()

	// Setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)

	access_token := "token here"
	httpClient := testutils.OAuthVcrSetup(t, false, testutils.WithVcrAccessToken(access_token))

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(
		definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"),
	)
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		testLogger,
		fakeSourceCredential,
		mockSourceCredentialRepository,
		cernerSandboxDefinition,
		models.WithTestHttpClient(httpClient),
	)
	require.NoError(t, err)

	// Read lookup_references.json and convert to sync.Map
	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/bundle/cerner_pending_resources.json")
	require.NoError(t, err)

	var refMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &refMap)
	require.NoError(t, err)

	referencedResourcesLookup := &sync.Map{}
	for k, v := range refMap {
		referencedResourcesLookup.Store(k, v)
	}
	client.SourceClientOptions.Concurrency = 5
	summary := models.UpsertSummary{
		UpdatedResources: []string{},
	}
	// Test
	_, syncErrors := client.ProcessPendingResources(fakeDatabase, &summary, referencedResourcesLookup, map[string]error{})
	for key, err := range syncErrors {
		if err != nil && !strings.Contains(err.Error(), "Skipping") {
			require.NoError(t, err, "error occurred while processing resource %s", key)
		}
	}

	// Assert
	// Assert
	require.Equal(t, 193, len(summary.UpdatedResources), "should have updated 193 resources")
	require.Equal(t, 194, summary.TotalResources, "should have processed 194 resources total")

}

func TestFhir401Client_ProcessPendingResource_Limit1(t *testing.T) {
	t.Parallel()

	// Setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	access_token := "token here"
	httpClient := testutils.OAuthVcrSetup(t, false, testutils.WithVcrAccessToken(access_token))

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(
		definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"),
	)
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		testLogger,
		fakeSourceCredential,
		mockSourceCredentialRepository,
		cernerSandboxDefinition,
		models.WithTestHttpClient(httpClient),
	)
	require.NoError(t, err)

	// Read lookup_references.json and convert to sync.Map
	jsonBytes, err := testutils.ReadTestFixture("testdata/fixtures/401-R4/bundle/cerner_pending_resources.json")
	require.NoError(t, err)

	var refMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &refMap)
	require.NoError(t, err)

	referencedResourcesLookup := &sync.Map{}
	for k, v := range refMap {
		referencedResourcesLookup.Store(k, v)
	}
	client.SourceClientOptions.Concurrency = 1
	summary := models.UpsertSummary{
		UpdatedResources: []string{},
	}

	// Test
	_, syncErrors := client.ProcessPendingResources(fakeDatabase, &summary, referencedResourcesLookup, map[string]error{})
	for key, err := range syncErrors {
		if err != nil && !strings.Contains(err.Error(), "Skipping") {
			require.NoError(t, err, "error occurred while processing resource %s", key)
		}
	}

	// Assert
	require.Equal(t, 193, len(summary.UpdatedResources), "should have updated 193 resources")
	require.Equal(t, 194, summary.TotalResources, "should have processed 194 resources total")

}

// Stub HTTP to return 403 so GetPatient error path runs.
// When scope lacks patient/* and patient/Patient.*, expect ErrScopePatientMissing.
func TestFHIR401_GetPatient_ReturnsScopeMissing_WhenScopeLacksPatient(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Stub a 403 response for any request (no network)
	roundTripper := rtFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 403,
			Status:     "403 Forbidden",
			Body:       io.NopCloser(strings.NewReader(`{"resourceType":"OperationOutcome"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Request:    r,
		}, nil
	})

	sourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	sourceCredential.EXPECT().GetAccessToken().AnyTimes().Return("acc")
	sourceCredential.EXPECT().GetRefreshToken().AnyTimes().Return("ref")
	// Make token appear valid so refresh isn't attempted
	sourceCredential.EXPECT().GetExpiresAt().AnyTimes().Return(time.Now().Add(30 * time.Minute).Unix())
	// Scope missing patient/* and patient/Patient.*
	sourceCredential.EXPECT().GetScope().AnyTimes().Return("openid profile launch/patient offline_access")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	logger := logrus.WithField("test", "401-scope-missing")
	defn, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		logger,
		sourceCredential,
		mockSourceCredentialRepository,
		defn,
		models.WithTestHttpClient(&http.Client{Transport: roundTripper}),
	)
	require.NoError(t, err)

	_, gotErr := client.GetPatient("some-patient-id")
	require.Error(t, gotErr)
	require.ErrorIs(t, gotErr, pkg.ErrScopePatientMissing)
}

// Stub HTTP to return 403 so GetPatient error path runs.
// When scope includes patient/Patient.read, expect ErrResourcePatientFailure.
func TestFHIR401_GetPatient_ReturnsPatientFailure_WhenScopeIncludesPatient(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roundTripper := rtFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 403,
			Status:     "403 Forbidden",
			Body:       io.NopCloser(strings.NewReader(`{"resourceType":"OperationOutcome"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Request:    r,
		}, nil
	})

	sourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	sourceCredential.EXPECT().GetAccessToken().AnyTimes().Return("acc")
	sourceCredential.EXPECT().GetRefreshToken().AnyTimes().Return("ref")
	sourceCredential.EXPECT().GetExpiresAt().AnyTimes().Return(time.Now().Add(30 * time.Minute).Unix())
	// Includes patient/Patient.read
	sourceCredential.EXPECT().GetScope().AnyTimes().Return("openid profile patient/Patient.read")
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	logger := logrus.WithField("test", "401-patient-failure")
	defn, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		logger,
		sourceCredential,
		mockSourceCredentialRepository,
		defn,
		models.WithTestHttpClient(&http.Client{Transport: roundTripper}),
	)
	require.NoError(t, err)

	_, gotErr := client.GetPatient("some-patient-id")
	require.Error(t, gotErr)
	require.ErrorIs(t, gotErr, pkg.ErrResourcePatientFailure)
}

// Stub HTTP to return 403 so GetPatient error path runs.
// When scope includes patient/*.read, expect ErrResourcePatientFailure.
func TestFHIR401_GetPatient_ReturnsPatientFailure_WhenScopeIncludesPatientWildcardRead(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roundTripper := rtFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 403,
			Status:     "403 Forbidden",
			Body:       io.NopCloser(strings.NewReader(`{"resourceType":"OperationOutcome"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Request:    r,
		}, nil
	})

	sourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	sourceCredential.EXPECT().GetAccessToken().AnyTimes().Return("acc")
	sourceCredential.EXPECT().GetRefreshToken().AnyTimes().Return("ref")
	sourceCredential.EXPECT().GetExpiresAt().AnyTimes().Return(time.Now().Add(30 * time.Minute).Unix())
	// Includes patient/*.read
	sourceCredential.EXPECT().GetScope().AnyTimes().Return("openid profile patient/*.read")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	logger := logrus.WithField("test", "401-patient-wildcard-read")
	defn, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		logger,
		sourceCredential,
		mockSourceCredentialRepository,
		defn,
		models.WithTestHttpClient(&http.Client{Transport: roundTripper}),
	)
	require.NoError(t, err)

	_, gotErr := client.GetPatient("some-patient-id")
	require.Error(t, gotErr)
	require.ErrorIs(t, gotErr, pkg.ErrResourcePatientFailure)
}

// Stub HTTP to return 403 so GetPatient error path runs.
// When scope includes patient/Patient.r (short form), expect ErrResourcePatientFailure.
func TestFHIR401_GetPatient_ReturnsPatientFailure_WhenScopeIncludesPatientShortR(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roundTripper := rtFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 403,
			Status:     "403 Forbidden",
			Body:       io.NopCloser(strings.NewReader(`{"resourceType":"OperationOutcome"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Request:    r,
		}, nil
	})

	sourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	sourceCredential.EXPECT().GetAccessToken().AnyTimes().Return("acc")
	sourceCredential.EXPECT().GetRefreshToken().AnyTimes().Return("ref")
	sourceCredential.EXPECT().GetExpiresAt().AnyTimes().Return(time.Now().Add(30 * time.Minute).Unix())
	// Includes patient/Patient.r
	sourceCredential.EXPECT().GetScope().AnyTimes().Return("openid profile patient/Patient.r")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	logger := logrus.WithField("test", "401-patient-short-r")
	defn, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		logger,
		sourceCredential,
		mockSourceCredentialRepository,
		defn,
		models.WithTestHttpClient(&http.Client{Transport: roundTripper}),
	)
	require.NoError(t, err)

	_, gotErr := client.GetPatient("some-patient-id")
	require.Error(t, gotErr)
	require.ErrorIs(t, gotErr, pkg.ErrResourcePatientFailure)
}

func TestGetRequest_NonJSON_ReturnsBinary(t *testing.T) {
	rawData := []byte("lorem fake data")

	// test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(rawData)
	}))
	defer ts.Close()

	// client under test
	client := &SourceClientBase{
		EndpointDefinition: &definitionsModels.LighthouseSourceDefinition{
			PatientAccessEndpoint: &catalog.PatientAccessEndpoint{
				Url: ts.URL + "/",
			},
		},
		SourceClientOptions: &models.SourceClientOptions{
			TestMode: true,
		},
		OauthClient: ts.Client(),
		Logger:      logrus.New(),
	}

	var binary map[string]interface{}

	returnedURL, err := client.GetRequest("/binary", &binary)
	require.NoError(t, err)

	prettyJson, err := json.MarshalIndent(binary, "", " ")
	require.NoError(t, err)
	fmt.Println("binary", string(prettyJson))

	// assert
	require.Equal(t, ts.URL+"/binary", returnedURL)

	require.Equal(t, "Binary", binary["resourceType"])
	require.Equal(t, "application/pdf", binary["contentType"])

	meta := binary["meta"].(map[string]interface{})
	source := meta["source"].(string)
	require.Equal(t, source, returnedURL)

	dataStr, ok := binary["data"].(string)
	require.True(t, ok)

	decodedData, err := base64.StdEncoding.DecodeString(dataStr)
	require.NoError(t, err)
	require.Equal(t, rawData, decodedData)

}
