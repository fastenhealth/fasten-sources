package platform

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/internal/base"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSourceClientCerner_SyncAll(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(694).Return(true, nil)

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("12742397")
	fakeSourceCredential.EXPECT().GetSourceType().AnyTimes().Return(pkg.SourceTypeCerner)
	fakeSourceCredential.EXPECT().GetApiEndpointBaseUrl().AnyTimes().Return("https://fhir-myrecord.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d")

	httpClient := base.OAuthVcrSetup(t, false)
	client, _, err := GetSourceClientCerner(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, httpClient)

	//test
	resp, err := client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, 931, resp.TotalResources)
	require.Equal(t, 694, len(resp.UpdatedResources))
}
