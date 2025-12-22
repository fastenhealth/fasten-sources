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

func TestGetSourceClientDrChrono_SyncAll(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("7670a27f-cb14-475a-ac72-c06a0a8dfea0")
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeDrChrono)
	fakeSourceCredential.EXPECT().GetEndpointId().AnyTimes().Return("6a01ffff-d73e-4728-a315-9b23bd77a4cc")

	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	httpClient := testutils.OAuthVcrSetup(t, false)
	client, err := GetDynamicSourceClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository, models.WithTestHttpClient(httpClient))

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 85, resp.TotalResources)
	require.Equal(t, 73, len(resp.UpdatedResources))
}
