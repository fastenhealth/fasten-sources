package client_auth_method_test

import (
	"bytes"
	"fmt"
	clientAuth "github.com/fastenhealth/fasten-sources/clients/client_auth_method"
	mockModels "github.com/fastenhealth/fasten-sources/clients/models/mock"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/jwt"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// test keyset for testing.
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

// The corresponding public keyset created with
// "tinkey create-public-keyset --in private_keyset.cfg"
const publicJSONKeyset = `{
		"primaryKeyId": 1742360595,
		"key": [
			{
				"keyData": {
					"typeUrl": "type.googleapis.com/google.crypto.tink.JwtEcdsaPublicKey",
					"value": "EAEaIG40E1603MP9RyoomZth6r+NcO1XeODPZKMmP/lbD/kgIiBeoDMF9LS5BDCh6YgqE3DjHwWwnEKEI3WpPf8izEx1rQ==",
					"keyMaterialType": "ASYMMETRIC_PUBLIC"
				},
				"status": "ENABLED",
				"keyId": 1742360595,
				"outputPrefixType": "TINK"
			}
		]
	}`

func TestCreatePrivateKeyJWTClientAssertion(t *testing.T) {
	// Mock keyset.Handle
	mockKeysetHandle, err := insecurecleartextkeyset.Read(
		keyset.NewJSONReader(bytes.NewBufferString(privateJSONKeyset)))
	require.NoError(t, err)

	// Mock parameters
	jwtIssuer := "test-issuer"
	jwtSubject := "test-subject"
	jwtAudience := "test-audience"

	// Call the function
	token, expires, err := clientAuth.CreatePrivateKeyJWTClientAssertion(mockKeysetHandle, jwtIssuer, jwtSubject, jwtAudience)

	log.Printf(token)
	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotNil(t, expires)
	assert.WithinDuration(t, time.Now().Add(4*time.Minute), *expires, time.Second)

	// Retrieve the Verifier primitive from publicKeysetHandle.
	publicKeysetHandle, err := keyset.ReadWithNoSecrets(
		keyset.NewJSONReader(bytes.NewBufferString(publicJSONKeyset)))
	require.NoError(t, err)
	verifier, err := jwt.NewVerifier(publicKeysetHandle)
	require.NoError(t, err)

	// Verify the signed token.
	jwtType := "JWT"
	validator, err := jwt.NewValidator(&jwt.ValidatorOpts{ExpectedTypeHeader: &jwtType, ExpectedAudience: &jwtAudience, ExpectedIssuer: &jwtIssuer})
	require.NoError(t, err)
	verifiedJWT, err := verifier.VerifyAndDecode(token, validator)
	require.NoError(t, err)
	extractedSubject, err := verifiedJWT.Subject()
	require.NoError(t, err)
	assert.Equal(t, jwtSubject, extractedSubject)
	extractedIssuer, err := verifiedJWT.Issuer()
	require.NoError(t, err)
	require.Equal(t, jwtIssuer, extractedIssuer)
	header, err := verifiedJWT.TypeHeader()
	require.NoError(t, err)
	assert.Equal(t, "JWT", header)

	fmt.Println(extractedSubject)
}

func TestPrivateKeyJWTBearerRefreshToken(t *testing.T) {
	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/token-endpoint", r.URL.Path)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), "grant_type=refresh_token")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "new-access-token", "expires_in": 3600}`))
	}))
	defer mockServer.Close()

	// Mock keyset.Handle
	mockKeysetHandle, err := insecurecleartextkeyset.Read(
		keyset.NewJSONReader(bytes.NewBufferString(privateJSONKeyset)))
	require.NoError(t, err)

	mockCtrl := gomock.NewController(t)
	sourceCredential := mockModels.NewMockSourceCredential(mockCtrl)
	sourceCredential.EXPECT().GetRefreshToken().Return("test-refresh-token").AnyTimes()

	endpointDef := definitionsModels.LighthouseSourceDefinition{
		ClientId: "test-client-id",
		PatientAccessEndpoint: &catalog.PatientAccessEndpoint{
			TokenEndpoint: mockServer.URL + "/token-endpoint",
		},
	}

	// Call the function
	globalLogger := logrus.New()
	response, err := clientAuth.PrivateKeyJWTBearerRefreshToken(globalLogger, mockKeysetHandle, sourceCredential, endpointDef)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "new-access-token", response.AccessToken)
	assert.Equal(t, int64(3600), response.ExpiresIn)
}

func TestPrivateKeyJWTBearerRefreshToken_ErrorResponse(t *testing.T) {
	// Mock HTTP server with error response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "invalid_client"}`))
	}))
	defer mockServer.Close()

	// Mock keyset.Handle
	mockKeysetHandle, err := insecurecleartextkeyset.Read(
		keyset.NewJSONReader(bytes.NewBufferString(privateJSONKeyset)))
	require.NoError(t, err)

	mockCtrl := gomock.NewController(t)
	sourceCredential := mockModels.NewMockSourceCredential(mockCtrl)
	sourceCredential.EXPECT().GetRefreshToken().Return("test-refresh-token").AnyTimes()

	endpointDef := definitionsModels.LighthouseSourceDefinition{
		ClientId: "test-client-id",
		PatientAccessEndpoint: &catalog.PatientAccessEndpoint{
			TokenEndpoint: mockServer.URL + "/token-endpoint",
		}}

	// Call the function
	globalLogger := logrus.New()
	response, err := clientAuth.PrivateKeyJWTBearerRefreshToken(globalLogger, mockKeysetHandle, sourceCredential, endpointDef)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "status code was not 200")
}
