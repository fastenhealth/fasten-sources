package client_auth_method

import (
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tink-crypto/tink-go/v2/jwt"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"io"
	"net/http"
	"net/url"
	"time"
)

// https://fhir.epic.com/Documentation?docId=oauth2&section=Standalone-Oauth2-Launch-Using-Refresh-Token-JWT
func CreatePrivateKeyJWTClientAssertion(jwtPrivateKeyHandle *keyset.Handle, jwtIssuer string, jwtSubject string, jwtAudience string) (string, *time.Time, error) {

	// Retrieve the JWT Signer primitive from privateKeysetHandle.
	signer, err := jwt.NewSigner(jwtPrivateKeyHandle)
	if err != nil {
		return "", nil, fmt.Errorf("%w: failed to generate signer for client assertion - %v", pkg.ErrSMARTTokenRefreshFailure, err)
	}

	// Use the primitive to create and sign a token. In this case, the primary key of the
	// keyset will be used (which is also the only key in this example).
	timeNow := time.Now()
	expires := timeNow.Add(time.Minute * 4)
	notBefore := timeNow.Add(-20 * time.Second) // allow for some clock skew
	jwtType := "JWT"
	jwtID := uuid.New().String() // A unique identifier for the JWT.
	rawJWT, err := jwt.NewRawJWT(&jwt.RawJWTOptions{
		Audience:   &jwtAudience,
		Subject:    &jwtSubject,
		Issuer:     &jwtIssuer,
		ExpiresAt:  &expires,
		IssuedAt:   &timeNow,
		NotBefore:  &notBefore,
		TypeHeader: &jwtType,
		JWTID:      &jwtID,
	})
	if err != nil {
		return "", nil, fmt.Errorf("%w: failed to generate JWT client assertion - %v", pkg.ErrSMARTTokenRefreshFailure, err)
	}
	token, err := signer.SignAndEncode(rawJWT)
	if err != nil {
		return "", nil, fmt.Errorf("%w: failed to sign JWT client assertion - %v", pkg.ErrSMARTTokenRefreshFailure, err)
	}
	return token, &expires, nil
}

// We're generating a new Access Token using a Refresh Token using a JWT Bearer token
// See:
// - https://fhir.epic.com/Documentation?docId=oauth2&section=Standalone-Oauth2-Launch-Using-Refresh-Token-JWT
func PrivateKeyJWTBearerRefreshToken(globalLogger logrus.FieldLogger, jwtPrivateKeyHandle *keyset.Handle, sourceCredential models.SourceCredential, endpointDef definitionsModels.LighthouseSourceDefinition, testHttpClient ...*http.Client) (*models.TokenRefreshResponse, error) {
	jwtToken, _, err := CreatePrivateKeyJWTClientAssertion(
		jwtPrivateKeyHandle,
		endpointDef.ClientId,
		endpointDef.ClientId,
		endpointDef.TokenEndpoint,
	)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating client assertion before refreshing token: %w", err)
	}

	//send this signed jwt to the token endpoint to get a new access token
	// https://fhir.epic.com/Documentation?docId=oauth2&section=JWKS

	postForm := url.Values{
		"grant_type":            {"refresh_token"},
		"refresh_token":         {sourceCredential.GetRefreshToken()},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"client_assertion":      {jwtToken},
	}

	var httpClient *http.Client
	if len(testHttpClient) > 0 && testHttpClient[0] != nil {
		httpClient = testHttpClient[0]
		//} else if debugMode == true && apiMode == pkg.ApiModeTest {
		//	//enable debug logging for sandbox mode only.
		//	logger.Warnf("debug mode enabled")
		//	httpClient = &http.Client{Transport: &debugSetterTransport{}}
	} else {
		httpClient = &http.Client{}
	}

	tokenResp, err := httpClient.PostForm(endpointDef.TokenEndpoint, postForm)

	if err != nil {
		return nil, fmt.Errorf("%w: an error occurred while sending client token request, %v", pkg.ErrSMARTTokenRefreshFailure, err)
	}

	defer tokenResp.Body.Close()
	if tokenResp.StatusCode >= 300 || tokenResp.StatusCode < 200 {

		b, err := io.ReadAll(tokenResp.Body)
		if err == nil {
			//TODO: we should eventually remove this logging.
			globalLogger.Errorf("Error Response body: %s", string(b))
		}

		return nil, fmt.Errorf("%w: an error occurred while reading token response, status code was not 200: %d", pkg.ErrSMARTTokenRefreshFailure, tokenResp.StatusCode)
	}

	var tokenResponse models.TokenRefreshResponse
	err = json.NewDecoder(tokenResp.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: an error occurred while parsing client token response: %v", pkg.ErrSMARTTokenRefreshFailure, err)
	}

	return &tokenResponse, nil

}
