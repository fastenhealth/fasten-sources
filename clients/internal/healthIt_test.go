package internal

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/testutils"
	"testing"

	"github.com/fastenhealth/fasten-sources/clients/internal/base"
	"github.com/fastenhealth/fasten-sources/clients/models"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetSourceClientHealthIt_SyncAll(t *testing.T) {
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
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("123d41e1-0f71-4e9f-8eb2-d1b1330201a6")
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeDynamicHealthIT)
	fakeSourceCredential.EXPECT().GetEndpointId().AnyTimes().Return("1c6664ef-aefa-422b-acc9-dc2494b327b8")

	fakeSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	httpClient := testutils.OAuthVcrSetup(t, false)
	client, err := GetDynamicSourceClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, fakeSourceCredentialRepository, models.WithTestHttpClient(httpClient))

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 274, resp.TotalResources)
	require.Equal(t, 274, len(resp.UpdatedResources))
}
