package internal

import (
	"context"
	"testing"

	"github.com/fastenhealth/fasten-sources/clients/internal/base"
	"github.com/fastenhealth/fasten-sources/clients/models"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetSourceClientMedhost_SyncAll(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	fakeDatabase.EXPECT().BackgroundJobCheckpoint(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return()

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("06da68a7d680692b380ce451dbcab682bf3ba63fe634e6db1892bb6c3b060623")
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeMedhost)
	fakeSourceCredential.EXPECT().GetEndpointId().AnyTimes().Return("9e5d5b7a-880f-481b-8ae4-77a3d24cfa49")

	httpClient := base.OAuthVcrSetup(t, false)
	client, err := GetDynamicSourceClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, models.WithTestHttpClient(httpClient))

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 139, resp.TotalResources)
	require.Equal(t, 141, len(resp.UpdatedResources))
}
