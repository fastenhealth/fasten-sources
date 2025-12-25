package client_auth_method_test

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/client_auth_method"
	"github.com/fastenhealth/fasten-sources/clients/testutils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"testing"
	"time"
)

func TestClientBasicSecretTokenRefreshToken_Epic(t *testing.T) {
	//setup
	oauthConfig := &oauth2.Config{
		ClientID:     "1e3ce324-5c1b-45a9-8ad1-7a0c1f51df43",
		ClientSecret: "PLACEHOLDER_CLIENT_SECRET",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/authorize",
			TokenURL: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token",
		},
		RedirectURL: "https://example.com",
	}

	testHttpClient := testutils.OAuthVcrSetup(t, false)

	tokenData := &oauth2.Token{
		TokenType:    "Bearer",
		AccessToken:  "PLACEHOLDER_ACCESS_TOKEN",
		RefreshToken: "PLACEHOLDER_REFRESH_TOKEN",
		Expiry:       time.Now().Add(-time.Minute),
	}

	response, err := client_auth_method.ClientSecretBasicRefreshToken(context.TODO(), logrus.New(), oauthConfig, tokenData, testHttpClient)
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "PLACEHOLDER_NEW_ACCESS_TOKEN", response.AccessToken)
	assert.Equal(t, "PLACEHOLDER_NEW_REFRESH_TOKEN", response.RefreshToken)
	assert.Equal(t, int64(3599), response.ExpiresIn)
}
