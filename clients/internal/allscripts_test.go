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

func TestGetSourceClientAllscripts_SyncAll(t *testing.T) {
	// TODO: need to regenerate with _count
	t.Skipf("skipping test, need to regenerate with _count")
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockStorageRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(158).Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("6709dc13-ca3e-4969-886a-fe0889eb8256")
	fakeSourceCredential.EXPECT().GetPlatformType().AnyTimes().Return(pkg.PlatformTypeAllscripts)
	fakeSourceCredential.EXPECT().GetEndpointId().AnyTimes().Return("7682675b-8247-4fda-b2cd-048bfeafc8af")

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	httpClient := base.OAuthVcrSetup(t, false)
	client, err := GetDynamicSourceClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, mockSourceCredentialRepository, models.WithTestHttpClient(httpClient))

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 158, resp.TotalResources)
	require.Equal(t, 158, len(resp.UpdatedResources))
}
