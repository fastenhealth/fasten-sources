package internal

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/testutils"
	"testing"

	"github.com/fastenhealth/fasten-sources/clients/models"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetSourceClientCareevolution_SyncAll(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(168).Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("6709dc13-ca3e-4969-886a-fe0889eb8256")
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeCareevolution)
	fakeSourceCredential.EXPECT().GetEndpointId().AnyTimes().Return("8b47cf7b-330e-4ede-9967-4caa7be623aa")
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	httpClient := testutils.OAuthVcrSetup(t, false)
	client, err := GetDynamicSourceClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, mockSourceCredentialRepository, models.WithTestHttpClient(httpClient))

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 168, resp.TotalResources)
	require.Equal(t, 168, len(resp.UpdatedResources))
}
