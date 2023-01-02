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

func TestGetSourceClientEpic_SyncAll(t *testing.T) {
	t.Parallel()
	//setup
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
	//fakeDatabase.EXPECT().UpsertRawResource("web.database.location").AnyTimes().Return(testDatabase.Name())

	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("erXuFYUfucBZaryVksYEcMg3")
	fakeSourceCredential.EXPECT().GetSourceType().AnyTimes().Return(pkg.SourceTypeEpic)
	fakeSourceCredential.EXPECT().GetApiEndpointBaseUrl().AnyTimes().Return("https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4")

	httpClient := base.OAuthVcrSetup(t, true)
	client, _, err := GetSourceClientEpic(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, httpClient)

	//test
	_, err = client.SyncAll(fakeDatabase)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
}
