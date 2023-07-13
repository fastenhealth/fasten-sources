package manual

import (
	"context"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetSourceClientManual_SyncAllBundle(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(195+1).Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetSourceType().AnyTimes().Return(pkg.SourceTypeManual)
	client, err := GetSourceClientManual(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential)

	bundleFile, err := os.Open("testdata/fixtures/401-R4/bundle/synthea_Tania553_Harris789_545c2380-b77f-4919-ab5d-0f615f877250.json")
	require.NoError(t, err)

	//test
	resp, err := client.SyncAllBundle(fakeDatabase, bundleFile, pkg.FhirVersion401)
	require.NoError(t, err)

	//assert
	require.Equal(t, 195+1, resp.TotalResources)
	require.Equal(t, 195+1, len(resp.UpdatedResources))
	require.NoError(t, err)
}
