package factory_test

import (
	"bytes"
	"context"
	"github.com/fastenhealth/fasten-sources/clients/factory"
	"github.com/fastenhealth/fasten-sources/clients/internal/fasten"
	"github.com/fastenhealth/fasten-sources/clients/internal/manual"
	"github.com/fastenhealth/fasten-sources/clients/internal/tefca_direct"
	"github.com/fastenhealth/fasten-sources/clients/internal/tefca_facilitated"
	"github.com/fastenhealth/fasten-sources/clients/models"
	mockModels "github.com/fastenhealth/fasten-sources/clients/models/mock"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"net/http"
	"testing"
	"time"
)

// test keyset from https://github.com/tink-crypto/tink-go/blob/main/jwt/jwt_test.go#L34C2-L47C4
const privateJSONKeyset = `{
		"primaryKeyId": 1742360595,
		"key": [
			{
				"keyData": {
					"typeUrl": "type.googleapis.com/google.crypto.tink.JwtEcdsaPrivateKey",
					"value": "GiBgVYdAPg3Fa2FVFymGDYrI1trHMzVjhVNEMpIxG7t0HRJGIiBeoDMF9LS5BDCh6YgqE3DjHwWwnEKEI3WpPf8izEx1rRogbjQTXrTcw/1HKiiZm2Hqv41w7Vd44M9koyY/+VsP+SAQAQ==",
					"keyMaterialType": "ASYMMETRIC_PRIVATE"
				},
				"status": "ENABLED",
				"keyId": 1742360595,
				"outputPrefixType": "TINK"
			}
		]
	}`

func TestGetSourceClientWithOptions(t *testing.T) {

	privateKeysetHandle, err := insecurecleartextkeyset.Read(
		keyset.NewJSONReader(bytes.NewBufferString(privateJSONKeyset)))
	if err != nil {
		t.Fatalf("failed to read keyset: %v", err)
	}

	//testing keyset handle

	tests := []struct {
		name             string
		platformType     pkg.PlatformType
		credentialType   pkg.SourceCredentialType
		endpointId       string
		clientId         string
		options          []func(*models.SourceClientOptions)
		expectedError    bool
		expectedClientFn func(t *testing.T, sc models.SourceClient)
	}{
		{
			name:          "Manual platform type with custom HTTP client",
			platformType:  pkg.PlatformTypeManual,
			options:       []func(*models.SourceClientOptions){},
			expectedError: false,
			expectedClientFn: func(t *testing.T, sc models.SourceClient) {
				// Assert that the client is of the correct type
				_, ok := sc.(*manual.ManualClient)
				assert.True(t, ok, "Expected SourceClient to be of type ManualClient")
			},
		},
		{
			name:          "Fasten platform type with test mode",
			platformType:  pkg.PlatformTypeFasten,
			options:       []func(*models.SourceClientOptions){},
			expectedError: false,
			expectedClientFn: func(t *testing.T, sc models.SourceClient) {
				// Assert that the client is of the correct type
				_, ok := sc.(*fasten.FastenClient)
				assert.True(t, ok, "Expected SourceClient to be of type FastenClient")
			},
		},
		{
			name:           "TEFCA Direct credential type with client ID and secret",
			platformType:   pkg.PlatformTypeTEFCA,
			credentialType: pkg.SourceCredentialTypeTefcaDirect,
			options:        []func(*models.SourceClientOptions){},
			expectedError:  false,
			expectedClientFn: func(t *testing.T, sc models.SourceClient) {
				// Assert that the client is of the correct type
				_, ok := sc.(*tefca_direct.TefcaClient)
				assert.True(t, ok, "Expected SourceClient to be of type TefcaClient")
			},
		},
		{
			name:           "TEFCA Facilitated credential type with redirect URL and scopes",
			platformType:   pkg.PlatformTypeTEFCAEpic,
			credentialType: pkg.SourceCredentialTypeTefcaFacilitated,
			endpointId:     "8e2f5de7-46ac-4067-96ba-5e3f60ad52a4",
			clientId:       "test-client-id",
			options: []func(*models.SourceClientOptions){
				models.WithTestHttpClient(&http.Client{}),
				models.WithClientJWTKeysetHandle(privateKeysetHandle),
			},
			expectedError: false,
			expectedClientFn: func(t *testing.T, sc models.SourceClient) {
				// Assert that the client is of the correct type
				_, tfcOk := sc.(*tefca_facilitated.TefcaFacilitatedFHIRClient)
				assert.True(t, tfcOk, "Expected SourceClient to be of type TefcaFacilitatedFHIRClient")

				//TODO: assert that the endpoint definition is set correctly
			},
		},
		//{
		//	name:           "Default to Smart on FHIR with resource type allow list",
		//	platformType:   pkg.PlatformTypeEpic,
		//	credentialType: pkg.SourceCredentialTypeSmartOnFhir,
		//	endpointId:     "8e2f5de7-46ac-4067-96ba-5e3f60ad52a4",
		//	clientId:       "test-client-id",
		//	options: []func(*models.SourceClientOptions){
		//		models.WithResourceTypeAllowList([]string{"Patient", "Observation"}),
		//	},
		//	expectedError: false,
		//	expectedClientFn: func(t *testing.T, sc models.SourceClient) {
		//		// Assert that the client is of the correct type
		//		_, tfcOk := sc.(*tefca_facilitated.TefcaFacilitatedFHIRClient)
		//		assert.True(t, tfcOk, "Expected SourceClient to be of type TefcaFacilitatedFHIRClient")
		//
		//	},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			logger := logrus.New()

			mockController := gomock.NewController(t)
			defer mockController.Finish()
			mockSourceCredential := mockModels.NewMockSourceCredential(mockController)
			mockSourceCredential.EXPECT().GetPlatformType().Return(tt.platformType).AnyTimes()
			mockSourceCredential.EXPECT().GetSourceCredentialType().Return(tt.credentialType).AnyTimes()
			mockSourceCredential.EXPECT().GetEndpointId().Return(tt.endpointId).AnyTimes()
			mockSourceCredential.EXPECT().GetClientId().Return(tt.clientId).AnyTimes()
			mockSourceCredential.EXPECT().GetRefreshToken().Return("testing").AnyTimes()
			mockSourceCredential.EXPECT().GetAccessToken().Return("").AnyTimes()
			mockSourceCredential.EXPECT().GetExpiresAt().Return(time.Now().Unix()).AnyTimes()

			client, err := factory.GetSourceClient(pkg.FastenLighthouseEnvSandbox, ctx, logger, mockSourceCredential, tt.options...)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
				// Add more assertions based on the expected client behavior
				if tt.expectedClientFn != nil {
					tt.expectedClientFn(t, client)
				}
			}
		})
	}
}
