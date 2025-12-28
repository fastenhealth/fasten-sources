package client_auth_method

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// NOTE: these methods are non-standard variations of the typical client authentication methods.
// These are specific to certain EHRs, like Epic, which require an access token to be used for authentication.
// SEE: client_secret_basic.go instead for the more common client authentication methods.

// ClientBearerTokenAuthIntrospectToken performs token introspection using bearer token. This is an uncommon variation specifically for Epic
// which requires an Access Token to be used for authentication, rather than a client id & secret which is more common.
// REQUIRES valid oauthTokenData.AccessToken to function
func ClientBearerTokenAuthIntrospectToken(
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	endpointDef definitionsModels.LighthouseSourceDefinition,
	oauthTokenData *oauth2.Token, //only used to provide the access token required to make the request
	tokenType models.TokenIntrospectTokenType,
	token string, //token to introspect, can be access or refresh token depending on tokenType
	testHttpClient ...*http.Client,
) (*models.TokenIntrospectResponse, error) {

	if len(token) == 0 {
		return nil, fmt.Errorf("no token (%s) available to introspect", tokenType)
	}
	if len(oauthTokenData.AccessToken) == 0 {
		return nil, fmt.Errorf("no access token available for bearer token authentication")
	}

	introspectEndpoint := endpointDef.IntrospectionEndpoint
	if len(introspectEndpoint) == 0 {
		//replace the token endpoint with the introspection endpoint
		introspectEndpoint = strings.TrimSuffix(strings.TrimSuffix(endpointDef.TokenEndpoint, "/"), "/token") + "/introspect"
	}

	formData := url.Values{}
	if tokenType == models.TokenIntrospectTokenTypeAccess {
		//formData.Set("token_type_hint", "access_token") //don't attempt to hint for access token
		formData.Set("token", token)

	} else if tokenType == models.TokenIntrospectTokenTypeRefresh {
		formData.Set("token_type_hint", "refresh_token")
		formData.Set("token", token)
	} else {
		return nil, errors.Errorf("no token (%s) available to introspect", tokenType)
	}

	bearerTokenAuthRequest, err := http.NewRequest("POST", introspectEndpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating bearer token protected request [%s]: %w", introspectEndpoint, err)
	}
	bearerTokenAuthRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	bearerTokenAuthRequest.Header.Set("Accept", "application/json")

	return clientBearerTokenAuthRequest[models.TokenIntrospectResponse](ctx, globalLogger, oauthTokenData, bearerTokenAuthRequest, testHttpClient...)
}

// ClientBearerTokenAuthUserInfoGetPatientId retrieves the patient ID from the userinfo endpoint using bearer token authentication method
// NOTE: some EHRs, like Epic, will not use client id & secret for userinfo, and will instead require the access token only.
// REQUIRES valid oauthTokenData.AccessToken to function
func ClientBearerTokenAuthUserInfoGetPatientId(
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	oauthTokenData *oauth2.Token,
	userInfoEndpoint string,
	testHttpClient ...*http.Client,
) (string, error) {

	bearerTokenAuthRequest, err := http.NewRequest("GET", userInfoEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("error creating bearer token protected request [%s]: %w", userInfoEndpoint, err)
	}
	bearerTokenAuthRequest.Header.Set("Accept", "application/json")

	respData, err := clientBearerTokenAuthRequest[map[string]any](ctx, globalLogger, oauthTokenData, bearerTokenAuthRequest, testHttpClient...)
	if err != nil {
		return "", fmt.Errorf("error getting patient info: %w", err)
	}
	if respData == nil {
		return "", errors.New("empty response from userinfo endpoint")
	} else if patientId, hasPatientId := (*respData)["patient"]; hasPatientId {
		return patientId.(string), nil
	}
	return "", errors.New("could not get patient from userinfo endpoint")
}

// #####################################################################################################################
// helpers Bearer
// #####################################################################################################################

// this is a generic function to make http calls protected by bearer token auth
// Used by token introspection, userinfo, etc.
// This function CANNOT be used for token refresh, as oauth2.Config.TokenSource() should be used instead for that purpose.
// It's only used for server-to-server communication. see Backend OAuth 2.0
func clientBearerTokenAuthRequest[T any](
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	token *oauth2.Token,
	request *http.Request,
	testHttpClient ...*http.Client,
) (*T, error) {
	if len(testHttpClient) > 0 && testHttpClient[0] != nil {
		globalLogger.Warnf("test httpClient provided, using it...")
		ctx = context.WithValue(ctx, oauth2.HTTPClient, testHttpClient[0])
	} else if debugMode == true {
		//} else if debugMode == true {
		//enable debug logging for sandbox mode only.
		globalLogger.Warnf("debug mode enabled")
		hc := &http.Client{Transport: &debugLoggingTransport{}}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, hc)
		//} else if !source.Confidential {
		//	//some sources (like Athena) seem to timeout if a Authorization header is provided, when a client_id:client_secret is not present.
		//	logger.Infof("clean authorization transport enabled for non-confidential source")
		//	hc := &http.Client{Transport: &cleanAuthorizationTransport{}}
		//	ctx = context.WithValue(ctx, oauth2.HTTPClient, hc)
	}

	bearerTokenAuthResponse, err := StaticOAuthClient(ctx, token).Do(request)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while sending bearer token protected request [%s]: %w", request.URL, err)
	}

	defer bearerTokenAuthResponse.Body.Close()
	if bearerTokenAuthResponse.StatusCode >= 300 || bearerTokenAuthResponse.StatusCode < 200 {

		b, err := io.ReadAll(bearerTokenAuthResponse.Body)
		if err == nil {
			//TODO: we should eventually remove this logging.
			globalLogger.Errorf("Error Response body: %s", string(b))
		}

		return nil, errors.Errorf("an error occurred while reading bearer token protected response, status code was not 200 [%s]: %d - %s", request.URL, bearerTokenAuthResponse.StatusCode, bearerTokenAuthResponse.Status)
	}

	var tokenResponse T
	err = json.NewDecoder(bearerTokenAuthResponse.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while parsing bearer token protected response [%s]: %w", request.URL, err)
	}

	return &tokenResponse, nil
}

// oauthConfig.Client() will automatically refresh the token if needed, so we need to protect ourselves from that possibility
// WE MUST DO THIS MANUALLY, as automatic refreshes will cause rolling refresh tokens to break (since we will not be able to store the newly issued refresh token)
// delete the refresh token to avoid accidental usage
func StaticOAuthClient(ctx context.Context, oauthToken *oauth2.Token) *http.Client {
	staticToken := &oauth2.Token{
		AccessToken:  oauthToken.AccessToken,
		TokenType:    oauthToken.TokenType,
		RefreshToken: "", //delete refresh token to avoid accidental usage
		Expiry:       oauthToken.Expiry,
	}

	oauthConfig := &oauth2.Config{} //empty config, we only need the Client() method.

	return oauthConfig.Client(ctx, staticToken)
}
