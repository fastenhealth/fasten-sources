package manual

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/models"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetSourceClientManual_ImplementsInterface(t *testing.T) {
	t.Parallel()

	//assert
	require.Implements(t, (*models.SourceClient)(nil), &ManualClient{}, "should implement the models.SourceClient interface")
}

func TestGetSourceClientManual_ExtractPatientId_Bundle(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)
	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/bundle/synthea_Tania553_Harris789_545c2380-b77f-4919-ab5d-0f615f877250.json")
	require.NoError(t, err)

	//test
	resp, ver, err := client.ExtractPatientId(bundleFile)

	//assert
	require.NoError(t, err)
	require.Equal(t, pkg.FhirVersion401, ver)
	require.Equal(t, "57959813-8cd2-4e3c-8970-e4364b74980a", resp)
}

func TestGetSourceClientManual_ExtractPatientId_CCDAToFHIRConvertedBundle(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)
	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/bundle/ccda_to_fhir_converted_C-CDA_R2-1_CCD.xml.json")
	require.NoError(t, err)

	//test
	resp, ver, err := client.ExtractPatientId(bundleFile)

	//assert
	require.NoError(t, err)
	require.Equal(t, pkg.FhirVersion401, ver)
	require.Equal(t, "12345", resp)
}

func TestGetSourceClientManual_ExtractPatientId_BundleIPS(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)
	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/international-patient-summary/IPS-bundle-01.json")
	require.NoError(t, err)

	//test
	resp, ver, err := client.ExtractPatientId(bundleFile)

	//assert
	require.NoError(t, err)
	require.Equal(t, pkg.FhirVersion401, ver)
	require.Equal(t, "2b90dd2b-2dab-4c75-9bb9-a355e07401e8", resp)
}

func TestGetSourceClientManual_ExtractPatientId_NDJSON(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)
	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/phr-ndjson-jsonl/TimmySmart-FosterCareTimeline.phr")
	require.NoError(t, err)

	//test
	resp, ver, err := client.ExtractPatientId(bundleFile)

	//assert
	require.NoError(t, err)
	require.Equal(t, pkg.FhirVersion401, ver)
	require.Equal(t, "12724069", resp)
}

func TestGetSourceClientManual_SyncAllBundle(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(234).Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)

	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/bundle/synthea_Tania553_Harris789_545c2380-b77f-4919-ab5d-0f615f877250.json")
	require.NoError(t, err)

	//test
	resp, err := client.SyncAllBundle(fakeDatabase, bundleFile, pkg.FhirVersion401)
	require.NoError(t, err)

	//assert
	require.Equal(t, 195+1, resp.TotalResources)
	require.Equal(t, 234, len(resp.UpdatedResources))
	require.NoError(t, err)
}

func TestGetSourceClientManual_SyncAllBundle_IPS(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(20).Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)

	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/international-patient-summary/IPS-bundle-01.json")
	require.NoError(t, err)

	//test
	resp, err := client.SyncAllBundle(fakeDatabase, bundleFile, pkg.FhirVersion401)
	require.NoError(t, err)

	//assert
	require.Equal(t, 20, resp.TotalResources)
	require.Equal(t, 20, len(resp.UpdatedResources))
	require.NoError(t, err)
}

func TestGetSourceClientManual_SyncAllBundle_NDJSON(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)

	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/phr-ndjson-jsonl/TimmySmart-FosterCareTimeline.phr")
	require.NoError(t, err)

	//test
	resp, err := client.SyncAllBundle(fakeDatabase, bundleFile, pkg.FhirVersion401)
	require.NoError(t, err)

	//assert
	require.Equal(t, 123, resp.TotalResources)
	require.Equal(t, 123, len(resp.UpdatedResources))
	require.NoError(t, err)
}

func TestGetSourceClientManual_SyncAllBundle_NDJSON2(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeManual)

	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/phr-ndjson-jsonl/JohnDoe.phr")
	require.NoError(t, err)

	//test
	resp, err := client.SyncAllBundle(fakeDatabase, bundleFile, pkg.FhirVersion401)
	require.NoError(t, err)

	//assert
	require.Equal(t, 12, resp.TotalResources)
	require.Equal(t, 12, len(resp.UpdatedResources))
	require.NoError(t, err)
}
