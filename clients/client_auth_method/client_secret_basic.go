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

	if tokenData == nil || len(tokenData.RefreshToken) == 0 {
		return nil, fmt.Errorf("%w: no refresh token available to refresh access token", pkg.ErrSMARTTokenRefreshFailure)
	}

	//client_secret_basic auth. If we need to modify significantly, this should be moved to clients/client_auth_method/client_secret_basic.go
	globalLogger.Info("using refresh token to generate access token...")

	if len(testHttpClient) > 0 && testHttpClient[0] != nil {
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

// ClientSecretBasicAuthIntrospectToken performs token introspection using client id and secret.
func ClientSecretBasicAuthIntrospectToken(
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	oauthConfig *oauth2.Config,
	endpointDef definitionsModels.LighthouseSourceDefinition,
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

	return clientSecretBasicAuthRequest[models.TokenIntrospectResponse](ctx, globalLogger, oauthConfig, basicAuthRequest, testHttpClient...)
}

//// ClientSecretBasicUserInfoGetPatientId retrieves the patient ID from the userinfo endpoint using client_secret_basic authentication method
//// NOTE: some EHRs, like Epic, will not use client id & secret for userinfo, and will instead require the access token only.
//// REQUIRES valid oauthTokenData.AccessToken or ID token to be passed in to function
//func ClientSecretBasicUserInfoGetPatientId(
//	ctx context.Context,
//	globalLogger logrus.FieldLogger,
//	oauthConfig *oauth2.Config,
//	userInfoEndpoint string,
//	testHttpClient ...*http.Client,
//) (string, error) {
//
//	basicAuthRequest, err := http.NewRequestWithContext(ctx, "GET", userInfoEndpoint, nil)
//	if err != nil {
//		return "", fmt.Errorf("error creating client secret protected request [%s]: %w", userInfoEndpoint, err)
//	}
//	basicAuthRequest.Header.Set("Accept", "application/json")
//
//	respData, err := clientSecretBasicAuthRequest[map[string]any](ctx, globalLogger, oauthConfig, basicAuthRequest, testHttpClient...)
//	if err != nil {
//		return "", fmt.Errorf("error getting patient info: %w", err)
//	}
//	if respData == nil {
//		return "", errors.New("empty response from userinfo endpoint")
//	} else if patientId, hasPatientId := (*respData)["patient"]; hasPatientId {
//		return patientId.(string), nil
//	}
//	return "", errors.New("could not get patient from userinfo endpoint")
//}

// #####################################################################################################################
// helpers Basic Authentication
// #####################################################################################################################

// this is a generic function to make http calls protected by client id & secret basic auth
// Used by token introspection, userinfo, etc.
// This function CANNOT be used for token refresh, as oauth2.Config.TokenSource() should be used instead for that purpose.
func clientSecretBasicAuthRequest[T any](
	ctx context.Context,
	globalLogger logrus.FieldLogger,
	oauthConfig *oauth2.Config,
	request *http.Request,
	testHttpClient ...*http.Client,
) (*T, error) {
	var httpClient *http.Client
	if len(testHttpClient) > 0 && testHttpClient[0] != nil {
		httpClient = testHttpClient[0]
	} else if debugMode == true {
		//enable debug logging for sandbox mode only.
		globalLogger.Warnf("debug mode enabled")
		httpClient = &http.Client{Transport: &debugLoggingTransport{}}
	} else if customHttpClient := ctx.Value(oauth2.HTTPClient); customHttpClient != nil {
		if customHttpClientCasted, customHttpClientCastedOk := customHttpClient.(*http.Client); customHttpClientCastedOk {
			httpClient = customHttpClientCasted
		} else {
			globalLogger.Warnf("unable to cast custom http client from context, using default http client")
			httpClient = &http.Client{}
		}
	} else {
		httpClient = &http.Client{}
	}

	//update the request with basic auth
	request.SetBasicAuth(oauthConfig.ClientID, oauthConfig.ClientSecret)

	basicAuthResponse, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while sending basic auth protected request [%s]: %w", request.URL, err)
	}

	defer basicAuthResponse.Body.Close()
	if basicAuthResponse.StatusCode >= 300 || basicAuthResponse.StatusCode < 200 {

		b, err := io.ReadAll(basicAuthResponse.Body)
		if err == nil {
			//TODO: we should eventually remove this logging.
			globalLogger.Errorf("Error Response body: %s", string(b))
		}

		return nil, errors.Errorf("an error occurred while reading basic auth protected response, status code was not 200 [%s]: %d", request.URL, basicAuthResponse.StatusCode)
	}

	var tokenResponse T
	err = json.NewDecoder(basicAuthResponse.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while parsing basic auth protected response [%s]: %w", request.URL, err)
	}

	return &tokenResponse, nil
}
