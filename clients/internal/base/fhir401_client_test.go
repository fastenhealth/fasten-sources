package base

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"testing"

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

func TestNewFHIR401Client(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sc := mock_models.NewMockSourceCredential(mockCtrl)
	sc.EXPECT().GetAccessToken().Return("test-access-token")
	sc.EXPECT().GetRefreshToken().Return("test-refresh-token")

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	endpointDefinition := &definitionsModels.LighthouseSourceDefinition{}

	//test
	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, endpointDefinition, models.WithTestHttpClient(&http.Client{}))

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
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cignaSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("6c0454af-1631-4c4d-905d-5710439df983"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, cignaSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)

	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/bundle/cigna_syntheticuser05-everything.json")
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
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)

	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/bundle/cerner_open_12724067_DocumentReference.json")
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
	db := mock_models.NewMockDatabaseRepository(mockCtrl)
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/document_reference/cerner_document_reference_206130480.json")
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
	referencedResourcesLookupMap := ToMap[string, bool](referencedResourcesLookup)
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
	db := mock_models.NewMockDatabaseRepository(mockCtrl)
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/document_reference/encounter_resource_broken_parse.json")
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

	referencedResourcesLookupMap := ToMap[string, bool](referencedResourcesLookup)

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
	db := mock_models.NewMockDatabaseRepository(mockCtrl)
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/document_reference/observation_broken_parse.json")
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
	referencedResourcesLookupMap := ToMap[string, bool](referencedResourcesLookup)

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
	db := mock_models.NewMockDatabaseRepository(mockCtrl)
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	medicareSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("6ae6c14e-b927-4ce0-862f-91123cb8d774"))
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, medicareSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	db.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/document_reference/medicare-eob.json")
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
	referencedResourcesLookupMap := ToMap[string, bool](referencedResourcesLookup)

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
	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)

	access_token := "token here"
	httpClient := OAuthVcrSetup(t, false, access_token)

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(
		definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"),
	)
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		testLogger,
		fakeSourceCredential,
		cernerSandboxDefinition,
		models.WithTestHttpClient(httpClient),
	)
	require.NoError(t, err)

	// Read lookup_references.json and convert to sync.Map
	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/bundle/cerner_pending_resources.json")
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
	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)

	access_token := "token here"
	httpClient := OAuthVcrSetup(t, false, access_token)

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(
		definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"),
	)
	require.NoError(t, err)

	client, err := GetSourceClientFHIR401(
		pkg.FastenLighthouseEnvSandbox,
		context.Background(),
		testLogger,
		fakeSourceCredential,
		cernerSandboxDefinition,
		models.WithTestHttpClient(httpClient),
	)
	require.NoError(t, err)

	// Read lookup_references.json and convert to sync.Map
	jsonBytes, err := ReadTestFixture("testdata/fixtures/401-R4/bundle/cerner_pending_resources.json")
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
