package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"golang.org/x/oauth2"
	"io"
	"log"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type SourceClientBase struct {
	FastenEnv pkg.FastenLighthouseEnvType
	Context   context.Context
	Logger    logrus.FieldLogger

	OauthClient      *http.Client
	SourceCredential models.SourceCredential
	Headers          map[string]string

	UsCoreResources []string
	FhirVersion     string

	//When test mode is enabled, tokens will not be refreshed, and Http client provided will be used (usually go-vcr for playback)
	testMode bool

	//Mutex to prevent multiple token refreshes from happening at the same time
	refreshMutex sync.Mutex
}

func (c *SourceClientBase) SyncAllBundle(db models.DatabaseRepository, bundleFile *os.File, bundleFhirVersion pkg.FhirVersion) (models.UpsertSummary, error) {
	panic("SyncAllBundle functionality is not available on this client")
}
func (c *SourceClientBase) ExtractPatientId(bundleFile *os.File) (string, pkg.FhirVersion, error) {
	panic("SyncAllBundle functionality is not available on this client")
}

func NewBaseClient(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (*SourceClientBase, error) {

	client := &SourceClientBase{
		FastenEnv:        env,
		Context:          ctx,
		Logger:           globalLogger,
		SourceCredential: sourceCreds,
		Headers:          map[string]string{},

		// https://build.fhir.org/ig/HL7/US-Core/
		UsCoreResources: []string{
			"AllergyIntolerance",
			//"Binary",
			"CarePlan",
			"CareTeam",
			"Condition",
			//"Coverage",
			"Device",
			"DiagnosticReport",
			"DocumentReference",
			"Encounter",
			"Goal",
			"Immunization",
			//"Location",
			//"Medication",
			"MedicationRequest",
			"Observation",
			//"Organization",
			//"Patient",
			//"Practitioner",
			//"PractitionerRole",
			"Procedure",
			//"Provenance",
			//"RelatedPerson",
			// "ServiceRequest",
			// "Specimen",
		},
	}

	if len(testHttpClient) > 0 {
		//Testing mode.
		client.testMode = true
		client.OauthClient = testHttpClient[0]
		client.OauthClient.Timeout = 10 * time.Second
	}

	err := client.RefreshAccessToken()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *SourceClientBase) GetUsCoreResources() []string {
	return c.UsCoreResources
}

func (c *SourceClientBase) GetSourceCredential() models.SourceCredential {
	return c.SourceCredential
}

func (c *SourceClientBase) RefreshAccessToken() error {
	if c.testMode {
		//if test mode is enabled, we cannot refresh the access token
		return nil
	}

	c.refreshMutex.Lock()
	defer c.refreshMutex.Unlock()

	//check if we need to refresh the access token
	//https://github.com/golang/oauth2/issues/84#issuecomment-520099526
	// https://chromium.googlesource.com/external/github.com/golang/oauth2/+/8f816d62a2652f705144857bbbcc26f2c166af9e/oauth2.go#239
	conf := &oauth2.Config{
		ClientID:     c.SourceCredential.GetClientId(),
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.SourceCredential.GetOauthAuthorizationEndpoint(),
			TokenURL: c.SourceCredential.GetOauthTokenEndpoint(),
		},
		//RedirectURL:  "",
		//Scopes:       nil,
	}

	token := &oauth2.Token{
		TokenType:    "Bearer",
		RefreshToken: c.SourceCredential.GetRefreshToken(),
		AccessToken:  c.SourceCredential.GetAccessToken(),
		Expiry:       time.Unix(c.SourceCredential.GetExpiresAt(), 0),
	}

	if token.Expiry.Before(time.Now().Add(5 * time.Second)) { // expired (or will expire in 5 seconds) so let's update it
		c.Logger.Info("access token expired, must refresh")

		if c.SourceCredential.IsDynamicClient() {
			c.Logger.Info("refreshing dynamic client...")
			err := c.SourceCredential.RefreshDynamicClientAccessToken()
			if err != nil {
				c.Logger.Error("error refreshing dynamic client: ", err)
				return err
			}

			//update the token with newly refreshed data
			token = &oauth2.Token{
				TokenType:    "Bearer",
				RefreshToken: c.SourceCredential.GetRefreshToken(),
				AccessToken:  c.SourceCredential.GetAccessToken(),
				Expiry:       time.Unix(c.SourceCredential.GetExpiresAt(), 0),
			}
		} else if len(c.SourceCredential.GetRefreshToken()) > 0 {
			c.Logger.Info("using refresh token to generate access token...")

			src := conf.TokenSource(c.Context, token)
			newToken, err := src.Token() // this actually goes and renews the tokens
			if err != nil {
				return err
			}
			log.Printf("new token expiry: %s", newToken.Expiry.Format(time.RFC3339))
			if newToken.AccessToken != token.AccessToken {
				token = newToken

				// update the "source" credential with new data (which will need to be sent
				c.SourceCredential.SetTokens(newToken.AccessToken, newToken.RefreshToken, newToken.Expiry.Unix())
				//updatedSource.AccessToken = newToken.AccessToken
				//updatedSource.ExpiresAt = newToken.Expiry.Unix()
				//// Don't overwrite `RefreshToken` with an empty value
				//// if this was a token refreshing request.
				//if newToken.RefreshToken != "" {
				//	updatedSource.RefreshToken = newToken.RefreshToken
				//}

			}
		} else {
			c.Logger.Error("no refresh token available, and not dynamic client. User must re-authenticate")
			return errors.New("no refresh token available, and not dynamic client. User must re-authenticate")
		}
	}

	c.OauthClient = oauth2.NewClient(c.Context, oauth2.StaticTokenSource(token))
	c.OauthClient.Timeout = 30 * time.Second

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// HttpClient
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//GetRequest makes a GET request to the specified resource subpath or a fully qualified url
//it will then decodes the response into the specified model (which should be a pointer to map[string]interface{})
//
//This function make the assumption that FHIR endpoint responses are always JSON
func (c *SourceClientBase) GetRequest(resourceSubpathOrNext string, decodeModelPtr interface{}) (string, error) {
	//check if we need to refresh the access token
	err := c.RefreshAccessToken()
	if err != nil {
		return "", err
	}

	resourceUrl, err := url.Parse(resourceSubpathOrNext)
	if err != nil {
		return "", err
	}
	if !resourceUrl.IsAbs() {
		resourceUrl, err = url.Parse(fmt.Sprintf("%s/%s", strings.TrimRight(c.SourceCredential.GetApiEndpointBaseUrl(), "/"), strings.TrimLeft(resourceSubpathOrNext, "/")))
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
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		b, _ := io.ReadAll(resp.Body)
		bodyContent := string(b)
		if len(bodyContent) > 300 {
			bodyContent = bodyContent[:300]
		}
		c.LoggerDebugResponse(resp, true)
		return "", fmt.Errorf("An error occurred during request %s - %d - %s [%s]", resourceUrl, resp.StatusCode, resp.Status, bodyContent)
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
			"resourceType": "Binary",
			"contentType":  contentTypeHeader,
			"data":         base64.StdEncoding.EncodeToString(b),
		})
		if err != nil {
			return "", fmt.Errorf("an error occurred while reading non-JSON response body: %s", err)
		}

		err = json.Unmarshal(binaryResourceJsonBytes, decodeModelPtr)
		if err != nil {
			return "", fmt.Errorf("an error occurred while creating Binary response body: %s", err)
		}

	} else {
		//this is JSON, unmarshal the model, and store it.
		err = UnmarshalJson(resp.Body, decodeModelPtr)
	}

	return resourceUrl.String(), err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func UnmarshalJson(r io.Reader, decodeModelPtr interface{}) error {
	decoder := json.NewDecoder(r)
	//decoder.DisallowUnknownFields() //make sure we throw an error if unknown fields are present.
	err := decoder.Decode(decodeModelPtr)
	if err != nil {
		return err
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
