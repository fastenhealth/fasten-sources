package client_auth_method

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ClientSecretBasicRefreshToken performs token refresh using client_secret_basic authentication methods (client_id:client_secret)
func ClientSecretBasicRefreshToken(
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	oauthConfig *oauth2.Config,
	tokenData *oauth2.Token,
	testHttpClient ...*http.Client,
) (*models.TokenRefreshResponse, error) {

	//client_secret_basic auth. If we need to modify significantly, this should be moved to clients/client_auth_method/client_secret_basic.go
	globalLogger.Info("using refresh token to generate access token...")

	if len(testHttpClient) > 0 {
		globalLogger.Warnf("test httpClient provided, using it...")
		ctx = context.WithValue(ctx, oauth2.HTTPClient, testHttpClient[0])
	} else if debugMode == true {
		//enable debug logging for sandbox mode only.
		globalLogger.Warnf("debug mode enabled")
		hc := &http.Client{Transport: &debugLoggingTransport{}}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, hc)
	}

	src := oauthConfig.TokenSource(ctx, tokenData)
	newToken, err := src.Token() // this actually goes and renews the tokens
	if err != nil {
		globalLogger.Infof("An error occurred during token refresh: %v", err)
		return nil, fmt.Errorf("%w: %v", pkg.ErrSMARTTokenRefreshFailure, err)
	}
	tokenRefreshResp := &models.TokenRefreshResponse{}
	globalLogger.Infof("new token expiry: %s", newToken.Expiry.Format(time.RFC3339))
	if newToken.AccessToken != tokenData.AccessToken {
		tokenData = newToken

		if scope := newToken.Extra("scope"); scope != nil {
			if scopeStr, scopeStrOk := scope.(string); scopeStrOk {
				tokenRefreshResp.Scope = scopeStr // safe casting and assignment
			}
		}
	}
	tokenRefreshResp.AccessToken = strings.TrimSpace(newToken.AccessToken)
	tokenRefreshResp.RefreshToken = strings.TrimSpace(newToken.RefreshToken)
	tokenRefreshResp.TokenType = newToken.TokenType
	tokenRefreshResp.ExpiresIn = int64(newToken.Expiry.Sub(time.Now()).Seconds())
	// Optionally get patient ID from userinfo endpoint if available

	return tokenRefreshResp, nil
}

// ClientSecretBasicIntrospectToken performs token introspection using client_secret_basic authentication method
// NOTE: some EHRs, like Epic, will not use client id & secret for introspection, and will instead require the access token only.
// REQUIRES valid oauthTokenData.AccessToken to function
func ClientSecretBasicIntrospectToken(
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	endpointDef definitionsModels.LighthouseSourceDefinition,
	oauthConfig *oauth2.Config, //only used to provide the configuration (clientid & client secret) required to make the request, not used to generate a oauth2.Client
	oauthTokenData *oauth2.Token, //only used to provide the access token required to make the request
	tokenType models.TokenIntrospectTokenType,
	token string, //token to introspect, can be access or refresh token depending on tokenType
	testHttpClient ...*http.Client,
) (*models.TokenIntrospectResponse, error) {

	if len(token) == 0 {
		return nil, fmt.Errorf("no token (%s) available to introspect", tokenType)
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

	basicAuthRequest, err := http.NewRequest("POST", introspectEndpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating client secret protected request [%s]: %w", introspectEndpoint, err)
	}
	basicAuthRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	basicAuthRequest.Header.Set("Accept", "application/json")

	return clientSecretBasicAuthRequest[models.TokenIntrospectResponse](ctx, globalLogger, oauthConfig, oauthTokenData, basicAuthRequest, testHttpClient...)
}

// ClientSecretBasicUserInfoGetPatientId retrieves the patient ID from the userinfo endpoint using client_secret_basic authentication method
// NOTE: some EHRs, like Epic, will not use client id & secret for userinfo, and will instead require the access token only.
// REQUIRES valid oauthTokenData.AccessToken to function
func ClientSecretBasicUserInfoGetPatientId(
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	oauthConfig *oauth2.Config,
	oauthTokenData *oauth2.Token,
	userInfoEndpoint string,
	testHttpClient ...*http.Client,
) (string, error) {

	basicAuthRequest, err := http.NewRequest("GET", userInfoEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("error creating client secret protected request [%s]: %w", userInfoEndpoint, err)
	}
	basicAuthRequest.Header.Set("Accept", "application/json")

	respData, err := clientSecretBasicAuthRequest[map[string]any](ctx, globalLogger, oauthConfig, oauthTokenData, basicAuthRequest, testHttpClient...)
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

// helpers

// this is a generic function to make http calls protected by basic auth
// Used by token introspection, userinfo, etc.
// This function should NOT be used for token refresh, as oauth2.Config.TokenSource() should be used instead for that purpose.
// This function will ignore the refresh token in the provided oauth2.Token, to avoid accidental usage. Instead it will create a "static" token
func clientSecretBasicAuthRequest[T any](
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	oauthConfig *oauth2.Config,
	token *oauth2.Token,
	request *http.Request,
	testHttpClient ...*http.Client,
) (*T, error) {
	if len(testHttpClient) > 0 {
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

	basicAuthResponse, err := StaticOAuthClient(ctx, oauthConfig, token).Do(request)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while sending client secret protected request [%s]: %w", request.URL, err)
	}

	defer basicAuthResponse.Body.Close()
	if basicAuthResponse.StatusCode >= 300 || basicAuthResponse.StatusCode < 200 {

		b, err := io.ReadAll(basicAuthResponse.Body)
		if err == nil {
			//TODO: we should eventually remove this logging.
			globalLogger.Errorf("Error Response body: %s", string(b))
		}

		return nil, errors.Errorf("an error occurred while reading client secret protected response, status code was not 200 [%s]: %d", request.URL, basicAuthResponse.StatusCode)
	}

	var tokenResponse T
	err = json.NewDecoder(basicAuthResponse.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while parsing client secret protected response [%s]: %w", request.URL, err)
	}

	return &tokenResponse, nil
}

// oauthConfig.Client() will automatically refresh the token if needed
// WE MUST DO THIS MANUALLY, as automatic refreshes will cause rolling refresh tokens to break (since we will not be able to store the newly issued refresh token)
// delete the refresh token to avoid accidental usage
func StaticOAuthClient(ctx context.Context, oauthConfig *oauth2.Config, oauthToken *oauth2.Token) *http.Client {
	staticToken := &oauth2.Token{
		AccessToken:  oauthToken.AccessToken,
		TokenType:    oauthToken.TokenType,
		RefreshToken: "", //delete refresh token to avoid accidental usage
		Expiry:       oauthToken.Expiry,
	}

	return oauthConfig.Client(ctx, staticToken)
}
