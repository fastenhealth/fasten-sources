package base

import (
	"bytes"
	"context"
	"github.com/fastenhealth/fasten-sources/clients/models"
	mock_models "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/definitions"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestSourceClientBase_RefreshAccessToken_WithoutValidAccessToken_ShouldFailDuringRefresh(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sc := mock_models.NewMockSourceCredential(mockCtrl)
	sc.EXPECT().GetAccessToken().Return("test-access-token").AnyTimes()
	sc.EXPECT().GetRefreshToken().Return("test-refresh-token").AnyTimes()
	sc.EXPECT().GetClientId().Return("test-client-id").AnyTimes()
	sc.EXPECT().SetTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	sc.EXPECT().GetExpiresAt().Return(time.Now().Add(-60 * time.Minute).Unix()).AnyTimes()

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"test": "TestSourceClientBase_RefreshAccessToken",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := NewBaseClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{}))
	require.NoError(t, err)
	require.NotNil(t, client)

	err = client.RefreshAccessToken(models.WithForce(true))
	require.ErrorIs(t, err, pkg.ErrSMARTTokenRefreshFailure)
}

func TestSourceClientBase_GetRequest_WithInvalidCredentials_ShouldReturnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sc := mock_models.NewMockSourceCredential(mockCtrl)
	sc.EXPECT().GetAccessToken().Return("test-access-token").AnyTimes()
	sc.EXPECT().GetRefreshToken().Return("test-refresh-token").AnyTimes()
	sc.EXPECT().GetClientId().Return("test-client-id").AnyTimes()
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"test": "TestSourceClientBase_GetRequest",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := NewBaseClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithTestHttpClient(&http.Client{
		Timeout: 2 * time.Second,
	}))
	require.NoError(t, err)
	require.NotNil(t, client)

	var decodeModel map[string]interface{}
	_, err = client.GetRequest("/test-resource", &decodeModel)
	require.ErrorIs(t, err, pkg.ErrResourceHttpError)
}

func TestUnmarshalJson_WithXMLPayload_ShouldReturnError(t *testing.T) {
	// Define a sample JSON input
	xmlInput := `This XML file does not appear to have any style information associated with it. The document tree is shown below.
<note>
<to>Tove</to>
<from>Jani</from>
<heading>Reminder</heading>
<body>Don't forget me this weekend!</body>
</note>`

	// Define a map to hold the unmarshaled data
	var result map[string]interface{}

	// Call the UnmarshalJson function
	err := UnmarshalJson(bytes.NewBufferString(xmlInput), &result)

	// Assert error occurred
	require.ErrorIs(t, err, pkg.ErrResourceInvalidContent)

}

func TestSourceClientBase_WithRetryableHttpClient_With429HTTPCodeResponse_ShouldRetry(t *testing.T) {
	t.Skipf("Skipping test as it requires a mock server to return 429 HTTP code. ")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sc := mock_models.NewMockSourceCredential(mockCtrl)
	sc.EXPECT().GetAccessToken().Return("test-access-token").AnyTimes()
	sc.EXPECT().GetRefreshToken().Return("test-refresh-token").AnyTimes()
	sc.EXPECT().GetClientId().Return("test-client-id").AnyTimes()
	sc.EXPECT().GetExpiresAt().Return(time.Now().Add(10 * time.Minute).Unix()).AnyTimes()
	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"test": "TestSourceClientBase_GetRequest",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := NewBaseClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithRetryableHttpClient())
	require.NoError(t, err)
	require.NotNil(t, client)

	var decodeModel map[string]interface{}
	_, err = client.GetRequest("https://mock.httpstatus.io/429", &decodeModel)
	require.ErrorIs(t, err, pkg.ErrResourceHttpError)
}

func TestSourceClientBase_WithRetryableHttpClient_With429HTTPCodeResponse_XRateLimitHeader_ShouldRetry(t *testing.T) {
	//t.Skipf("Skipping test as it requires a mock server to return 429 HTTP code with X-RateLimit headers.")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sc := mock_models.NewMockSourceCredential(mockCtrl)
	sc.EXPECT().GetAccessToken().Return("test-access-token").AnyTimes()
	sc.EXPECT().GetRefreshToken().Return("test-refresh-token").AnyTimes()
	sc.EXPECT().GetClientId().Return("test-client-id").AnyTimes()
	sc.EXPECT().GetExpiresAt().Return(time.Now().Add(10 * time.Minute).Unix()).AnyTimes()

	mockSourceCredentialRepository := mock_models.NewMockSourceCredentialRepository(mockCtrl)

	testLogger := logrus.WithFields(logrus.Fields{
		"test": "TestSourceClientBase_GetRequest",
	})

	cernerSandboxDefinition, err := definitions.GetSourceDefinition(definitions.WithEndpointId("3290e5d7-978e-42ad-b661-1cf8a01a989c"))
	require.NoError(t, err)

	client, err := NewBaseClient(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, sc, mockSourceCredentialRepository, cernerSandboxDefinition, models.WithRetryableHttpClient())
	require.NoError(t, err)
	require.NotNil(t, client)

	var decodeModel map[string]interface{}
	_, err = client.GetRequest("https://mock.httpstatus.io/429", &decodeModel)
	require.ErrorIs(t, err, pkg.ErrResourceHttpError)
}
