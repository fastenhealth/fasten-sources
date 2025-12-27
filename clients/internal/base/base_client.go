package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fastenhealth/fasten-sources/clients/client_auth_method"
	"github.com/fastenhealth/fasten-sources/clients/models"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"golang.org/x/oauth2"
)

type SourceClientBase struct {
	FastenEnv pkg.FastenLighthouseEnvType
	Context   context.Context
	Logger    logrus.FieldLogger

	OauthClient                *http.Client
	SourceCredential           models.SourceCredential
	SourceCredentialRepository models.SourceCredentialRepository
	EndpointDefinition         *definitionsModels.LighthouseSourceDefinition
	Headers                    map[string]string

	ResourceTypesUsCore []string
	FhirVersion         string

	SourceClientOptions *models.SourceClientOptions

	//Mutex to prevent multiple token refreshes from happening at the same time
	refreshMutex sync.Mutex
}

func (c *SourceClientBase) SyncAllBundle(db models.StorageRepository, bundleFile *os.File, bundleFhirVersion pkg.FhirVersion) (models.UpsertSummary, error) {
	panic("SyncAllBundle functionality is not available on this client")
}
func (c *SourceClientBase) ExtractPatientId(bundleFile *os.File) (string, pkg.FhirVersion, error) {
	panic("SyncAllBundle functionality is not available on this client")
}

func NewBaseClient(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, sourceCredsDb models.SourceCredentialRepository, endpointDefinition *definitionsModels.LighthouseSourceDefinition, options ...func(clientOpts *models.SourceClientOptions)) (*SourceClientBase, error) {

	clientOptions := &models.SourceClientOptions{
		SourceClientRefreshOptions: []func(*models.SourceClientRefreshOptions){},
		Context:                    ctx,
	}

	if endpointDefinition.ClientRateLimited {
		options = append(options, models.WithRetryableHttpClient()) //make sure we have a retryable http client if the platform_type is rate limited
	}

	for _, o := range options {
		o(clientOptions)
	}

	client := &SourceClientBase{
		FastenEnv:                  env,
		Context:                    clientOptions.Context,
		Logger:                     globalLogger,
		SourceCredential:           sourceCreds,
		SourceCredentialRepository: sourceCredsDb,
		EndpointDefinition:         endpointDefinition,
		Headers:                    map[string]string{},

		// https://build.fhir.org/ig/HL7/US-Core/
		ResourceTypesUsCore: []string{
			"AllergyIntolerance",
			//"Binary",
			"CarePlan",
			"CareTeam",
			"Condition",
			"Coverage",
			"Device",
			"DiagnosticReport",
			"DocumentReference",
			"Encounter",
			"FamilyMemberHistory",
			"Goal",
			"Immunization",
			//"Location",
			//"Medication",
			"MedicationDispense",
			"MedicationRequest",
			"Observation",
			//"Organization",
			//"Patient",
			//"Practitioner",
			//"PractitionerRole",
			"Procedure",
			"Provenance",
			"QuestionnaireResponse",
			"RelatedPerson",
			"ServiceRequest",
			"Specimen",
		},
		SourceClientOptions: clientOptions,
	}

	if client.SourceClientOptions.TestHttpClient != nil {
		//Testing mode.
		client.OauthClient = clientOptions.TestHttpClient
		client.OauthClient.Timeout = 10 * time.Second
	}

	err := client.RefreshAccessToken(client.SourceClientOptions.SourceClientRefreshOptions...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *SourceClientBase) GetResourceTypesUsCore() []string {
	return c.ResourceTypesUsCore
}

func (c *SourceClientBase) GetResourceTypesAllowList() []string {
	if c.SourceClientOptions == nil {
		return []string{}
	} else {
		return c.SourceClientOptions.ResourceTypesAllowList
	}
}

func (c *SourceClientBase) GetSourceCredential() models.SourceCredential {
	return c.SourceCredential
}

func (c *SourceClientBase) RefreshAccessToken(options ...func(*models.SourceClientRefreshOptions)) error {
	refreshOptions := &models.SourceClientRefreshOptions{}
	for _, o := range options {
		o(refreshOptions)
	}

	if !refreshOptions.Force && c.SourceClientOptions.TestMode {
		//if not Forced, and test mode is enabled, we cannot refresh the access token
		return nil
	}

	c.refreshMutex.Lock()
	defer c.refreshMutex.Unlock() //NOTE: we must call c.SourceCredentialRepository.StoreToken() before this lock completes (function exit) -- if required.

	//check if we need to refresh the access token
	//https://github.com/golang/oauth2/issues/84#issuecomment-520099526
	// https://chromium.googlesource.com/external/github.com/golang/oauth2/+/8f816d62a2652f705144857bbbcc26f2c166af9e/oauth2.go#239
	conf := c.generateOAuthConfig()
	token := c.generateOAuthTokenFromCredential()

	tokenWasRefreshed := false
	if refreshOptions.Force || token.Expiry.Before(time.Now().Add(5*time.Second)) { // expired (or will expire in 5 seconds) so let's update it
		if refreshOptions.Force {
			c.Logger.Info("force refresh access token", refreshOptions.Force)
		} else {
			c.Logger.Info("access token expired, must refresh")
		}

		//check the authentication method type
		clientAuthMethod := c.EndpointDefinition.GetClientAuthMethod()

		if clientAuthMethod == pkg.ClientAuthenticationMethodTypePrivateKeyJwt {
			// this is a private key JWT client, we need to refresh the token using the private key

			if c.SourceClientOptions.ClientJWTKeysetHandle == nil {
				c.Logger.Error("no jwt keyset handle provided for private key JWT client")
				return fmt.Errorf("%w: unable to generate client assertion, missing keyset", pkg.ErrSMARTTokenRefreshFailure)
			}

			c.Logger.Info("refreshing using JWT private key...")
			tokenRefreshResponse, err := client_auth_method.PrivateKeyJWTBearerRefreshToken(
				c.Context,
				c.Logger,
				c.SourceCredential.GetClientId(),
				c.SourceClientOptions.ClientJWTKeysetHandle,
				*c.EndpointDefinition,
				c.SourceCredential.GetRefreshToken(),
				c.SourceClientOptions.TestHttpClient,
			)
			if err != nil {
				c.Logger.Error("error refreshing JWT client: ", err)
				return fmt.Errorf("%w: %v", pkg.ErrSMARTTokenRefreshFailure, err)
			}

			c.SourceCredential.SetTokens(
				tokenRefreshResponse.AccessToken,
				tokenRefreshResponse.RefreshToken,
				time.Now().Add(time.Second*time.Duration(tokenRefreshResponse.ExpiresIn)).Unix(),
				tokenRefreshResponse.Scope,
				nil,
			)

			//update the token with newly refreshed data
			token = &oauth2.Token{
				TokenType:    "Bearer",
				RefreshToken: c.SourceCredential.GetRefreshToken(),
				AccessToken:  c.SourceCredential.GetAccessToken(),
				Expiry:       time.Unix(c.SourceCredential.GetExpiresAt(), 0),
			}
			tokenWasRefreshed = true

		} else if len(c.SourceCredential.GetRefreshToken()) > 0 {

			c.Logger.Info("refreshing using client id & secret...")
			tokenRefreshResponse, err := client_auth_method.ClientSecretBasicRefreshToken(
				c.Context,
				c.Logger,
				conf,
				token,
				c.SourceClientOptions.TestHttpClient,
			)
			if err != nil {
				c.Logger.Error("error refreshing confidential client: ", err)
				return fmt.Errorf("%w: %v", pkg.ErrSMARTTokenRefreshFailure, err)
			}

			c.SourceCredential.SetTokens(
				tokenRefreshResponse.AccessToken,
				tokenRefreshResponse.RefreshToken,
				time.Now().Add(time.Second*time.Duration(tokenRefreshResponse.ExpiresIn)).Unix(),
				tokenRefreshResponse.Scope,
				nil,
			)

			//update the token with newly refreshed data
			token = &oauth2.Token{
				TokenType:    "Bearer",
				RefreshToken: c.SourceCredential.GetRefreshToken(),
				AccessToken:  c.SourceCredential.GetAccessToken(),
				Expiry:       time.Unix(c.SourceCredential.GetExpiresAt(), 0),
			}
			tokenWasRefreshed = true
		} else {
			c.Logger.Error("no refresh token available, and does not support JWT refresh. User must re-authenticate")
			return fmt.Errorf("%w: no refresh token available, and does not support JWT refresh. User must re-authenticate", pkg.ErrSMARTTokenRefreshFailure)
		}

		c.Logger.Infof("Access token refreshed successfully, expires at %s", token.Expiry.Format(time.RFC3339))
	}

	c.OauthClient = oauth2.NewClient(c.Context, oauth2.StaticTokenSource(token))
	c.OauthClient.Timeout = 120 * time.Second

	if tokenWasRefreshed {
		//try to introspect the refresh token to get the expiration if possible.
		if introspectData, introspectErr := c.IntrospectToken(models.TokenIntrospectTokenTypeRefresh); introspectErr == nil && introspectData.ExpiresAt > 0 {
			c.Logger.Infof("introspected refresh token expiration: %d", introspectData.ExpiresAt)
			refreshExpirationUnix := time.Unix(int64(introspectData.ExpiresAt), 0).Unix()
			c.SourceCredential.SetTokens(c.SourceCredential.GetAccessToken(), c.SourceCredential.GetRefreshToken(), c.SourceCredential.GetExpiresAt(), c.SourceCredential.GetScope(), &refreshExpirationUnix)
		} else {
			c.Logger.Warnf("unable to introspect refresh token for expiration: %v", introspectErr)
		}

		//store the updated tokens
		err := c.SourceCredentialRepository.StoreTokens(c.Context, c.SourceCredential)
		if err != nil {
			c.Logger.Warnf("IGNORED: error storing tokens after refresh: ", err)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// HttpClient
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GetRequest makes a GET request to the specified resource subpath or a fully qualified url
// it will then decodes the response into the specified model (which should be a pointer to map[string]interface{})
//
// This function make the assumption that FHIR endpoint responses are always JSON
func (c *SourceClientBase) GetRequest(resourceSubpathOrNext string, decodeModelPtr interface{}) (string, error) {
	//check if we need to refresh the access token
	err := c.RefreshAccessToken(c.SourceClientOptions.SourceClientRefreshOptions...)
	if err != nil {
		return "", err
	}

	resourceUrl, err := url.Parse(resourceSubpathOrNext)
	if err != nil {
		return "", err
	}
	if !resourceUrl.IsAbs() {
		resourceUrl, err = url.Parse(fmt.Sprintf("%s/%s", strings.TrimRight(c.EndpointDefinition.Url, "/"), strings.TrimLeft(resourceSubpathOrNext, "/")))
	}
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodGet, resourceUrl.String(), nil)
	if err != nil {
		return "", err
	}

	for key, val := range c.Headers {
		//req.Header.Add("Accept", "application/json+fhir")
		req.Header.Add(key, val)
	}

	c.LoggerDebugRequest(req)
	resp, err := c.OauthClient.Do(req)
	if err != nil {
		c.LoggerDebugResponse(resp, true)
		return "", fmt.Errorf("%w: %v", pkg.ErrResourceHttpError, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		b, _ := io.ReadAll(resp.Body)
		bodyContent := string(b)
		if len(bodyContent) > 600 {
			bodyContent = bodyContent[:600]
		}
		c.LoggerDebugResponse(resp, true)
		return "", fmt.Errorf("%w: an error occurred during request %s - %d - %s [%s]", pkg.ErrResourceHttpError, resourceUrl, resp.StatusCode, resp.Status, bodyContent)
	}
	contentTypeHeader := resp.Header.Get("Content-Type")
	if !isContentTypeJsonAnalog(contentTypeHeader) {
		c.LoggerDebugResponse(resp, false)
		c.Logger.Warnf("response content type is not JSON: `%s`. This should only happen for Binary resource types", resp.Header.Get("Content-Type"))

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("an error occurred while reading non-JSON response body: %s", err)
		}
		binaryResourceJsonBytes, err := json.Marshal(map[string]interface{}{
			"id":           base64.StdEncoding.EncodeToString([]byte(resourceUrl.String())),
			"resourceType": "Binary",
			"contentType":  contentTypeHeader,
			"data":         base64.StdEncoding.EncodeToString(b),
		})
		if err != nil {
			return "", fmt.Errorf("%w: an error occurred while reading non-JSON response body: %s", pkg.ErrResourceInvalidContent, err)
		}

		err = json.Unmarshal(binaryResourceJsonBytes, decodeModelPtr)
		if err != nil {
			return "", fmt.Errorf("%w: an error occurred while creating Binary response body: %s", pkg.ErrResourceInvalidContent, err)
		}

	} else {
		//this is JSON, unmarshal the model, and store it.
		err = UnmarshalJson(resp.Body, decodeModelPtr)
	}

	return resourceUrl.String(), err
}

// Use token introspection from Token Introspectin endpoint
// https://build.fhir.org/ig/HL7/smart-app-launch/token-introspection.html
func (c *SourceClientBase) IntrospectToken(tokenType models.TokenIntrospectTokenType) (*models.TokenIntrospectResponse, error) {

	var introspectToken string
	if tokenType == models.TokenIntrospectTokenTypeAccess {
		introspectToken = c.SourceCredential.GetAccessToken()

	} else if tokenType == models.TokenIntrospectTokenTypeRefresh {
		introspectToken = c.SourceCredential.GetRefreshToken()
	}
	if len(introspectToken) == 0 {
		return nil, fmt.Errorf("no token (%s) available to introspect", tokenType)
	}

	conf := c.generateOAuthConfig()
	token := c.generateOAuthTokenFromCredential()
	clientAuthMethod := c.EndpointDefinition.GetClientAuthMethod()

	if c.EndpointDefinition.PlatformType == pkg.PlatformTypeEpic {
		//try to introspect the refresh token to get the expiration if possible.
		introspectData, introspectErr := client_auth_method.ClientBearerTokenAuthIntrospectToken(
			c.Context,
			c.Logger,
			*c.EndpointDefinition,
			token,
			tokenType,
			introspectToken,
			c.SourceClientOptions.TestHttpClient,
		)
		if introspectErr != nil {
			return nil, introspectErr
		}
		return introspectData, nil
	} else if clientAuthMethod == pkg.ClientAuthenticationMethodTypePrivateKeyJwt {
		// this is a private key JWT client, we need to refresh the token using the private key

		if c.SourceClientOptions.ClientJWTKeysetHandle == nil {
			c.Logger.Error("no jwt keyset handle provided for private key JWT client")
			return nil, fmt.Errorf("%w: unable to generate client assertion, missing keyset", pkg.ErrSMARTTokenRefreshFailure)
		}

		//try to introspect the refresh token to get the expiration if possible.
		introspectData, introspectErr := client_auth_method.PrivateKeyJWTBearerIntrospectToken(
			c.Context,
			c.Logger,
			c.SourceCredential.GetClientId(),
			c.SourceClientOptions.ClientJWTKeysetHandle,
			*c.EndpointDefinition,
			tokenType,
			introspectToken,
			c.SourceClientOptions.TestHttpClient,
		)
		if introspectErr != nil {
			return nil, introspectErr
		}
		return introspectData, nil
	} else {
		//try to introspect the refresh token to get the expiration if possible.
		introspectData, introspectErr := client_auth_method.ClientSecretBasicAuthIntrospectToken(
			c.Context,
			c.Logger,
			conf,
			*c.EndpointDefinition,
			tokenType,
			introspectToken,
			c.SourceClientOptions.TestHttpClient,
		)
		if introspectErr != nil {
			return nil, introspectErr
		}
		return introspectData, nil
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helper Functions
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientBase) generateOAuthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: c.SourceCredential.GetClientId(),
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.EndpointDefinition.AuthorizationEndpoint,
			TokenURL: c.EndpointDefinition.TokenEndpoint,
		},
		//ClientSecret: "",
		//RedirectURL:  "",
		//Scopes:       nil,
	}

	//overrides
	if len(c.SourceClientOptions.ClientID) > 0 {
		conf.ClientID = c.SourceClientOptions.ClientID
	}
	if len(c.SourceClientOptions.ClientSecret) > 0 {
		conf.ClientSecret = c.SourceClientOptions.ClientSecret
	}
	if len(c.SourceClientOptions.RedirectURL) > 0 {
		conf.RedirectURL = c.SourceClientOptions.RedirectURL
	}
	if len(c.SourceClientOptions.Scopes) > 0 {
		conf.Scopes = c.SourceClientOptions.Scopes
	}
	return conf
}

func (c *SourceClientBase) generateOAuthTokenFromCredential() *oauth2.Token {
	return &oauth2.Token{
		TokenType:    "Bearer",
		RefreshToken: c.SourceCredential.GetRefreshToken(),
		AccessToken:  c.SourceCredential.GetAccessToken(),
		Expiry:       time.Unix(c.SourceCredential.GetExpiresAt(), 0),
	}
}

func UnmarshalJson(r io.Reader, decodeModelPtr interface{}) error {
	decoder := json.NewDecoder(r)
	//decoder.DisallowUnknownFields() //make sure we throw an error if unknown fields are present.
	err := decoder.Decode(decodeModelPtr)
	if err != nil {
		return fmt.Errorf("%w: %v", pkg.ErrResourceInvalidContent, err)
	}
	return err
}

func (c *SourceClientBase) LoggerDebugRequest(req *http.Request) {
	if req == nil {
		return
	}
	if dumpReq, dumpReqErr := httputil.DumpRequest(req, true); dumpReqErr == nil {
		c.Logger.Debug("Request: ", string(dumpReq))
	}
}

func (c *SourceClientBase) LoggerDebugResponse(resp *http.Response, dumpBody bool) {
	if resp == nil {
		return
	}
	if dumpResp, dumpRespErr := httputil.DumpResponse(resp, dumpBody); dumpRespErr == nil {
		c.Logger.Debug("Response: ", string(dumpResp))
	}
}

func isContentTypeJsonAnalog(contentType string) bool {
	if contentType == "" {
		return false
	}
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	return slices.Contains([]string{"application/json", "application/fhir+json", "application/json+fhir"}, mediatype)
}
