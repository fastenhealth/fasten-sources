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

func TestGetSourceClientKaiser_SyncAll(t *testing.T) {
	//TODO: need to regenerate with _count
	// t.Skipf("skipping test, need to regenerate with _count")
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
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("4a085566-49d8-fa6f-d934-8494253b3148")
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeKaiser)
	fakeSourceCredential.EXPECT().GetEndpointId().AnyTimes().Return("9d0fa28a-0c5b-4065-9ee6-284ec9577a57")

	httpClient := base.OAuthVcrSetup(t, false)
	client, err := GetDynamicSourceClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, models.WithTestHttpClient(httpClient))

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 7, resp.TotalResources)
	require.Equal(t, 7, len(resp.UpdatedResources))
}
