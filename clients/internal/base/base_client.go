package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type SourceClientBase struct {
	FastenEnv pkg.FastenEnvType
	Context   context.Context
	Logger    logrus.FieldLogger

	OauthClient      *http.Client
	SourceCredential models.SourceCredential
	Headers          map[string]string

	UsCoreResources []string
	FhirVersion     string
}

func (c *SourceClientBase) SyncAllBundle(db models.DatabaseRepository, bundleFile *os.File) error {
	panic("SyncAllBundle functionality is not available on this client")
}

func NewBaseClient(env pkg.FastenEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (*SourceClientBase, *models.SourceCredential, error) {
	var httpClient *http.Client
	var updatedSource *models.SourceCredential
	if len(testHttpClient) == 0 {
		//check if we need to refresh the access token
		//https://github.com/golang/oauth2/issues/84#issuecomment-520099526
		// https://chromium.googlesource.com/external/github.com/golang/oauth2/+/8f816d62a2652f705144857bbbcc26f2c166af9e/oauth2.go#239
		conf := &oauth2.Config{
			ClientID:     sourceCreds.GetClientId(),
			ClientSecret: "",
			Endpoint: oauth2.Endpoint{
				AuthURL:  sourceCreds.GetOauthAuthorizationEndpoint(),
				TokenURL: sourceCreds.GetOauthTokenEndpoint(),
			},
			//RedirectURL:  "",
			//Scopes:       nil,
		}
		token := &oauth2.Token{
			TokenType:    "Bearer",
			RefreshToken: sourceCreds.GetRefreshToken(),
			AccessToken:  sourceCreds.GetAccessToken(),
			Expiry:       time.Unix(sourceCreds.GetExpiresAt(), 0),
		}
		if token.Expiry.Before(time.Now()) { // expired so let's update it
			log.Println("access token expired, refreshing...")
			src := conf.TokenSource(ctx, token)
			newToken, err := src.Token() // this actually goes and renews the tokens
			if err != nil {
				return nil, nil, err
			}
			if newToken.AccessToken != token.AccessToken {
				token = newToken

				// update the "source" credential with new data (which will need to be sent
				sourceCreds.RefreshTokens(newToken.AccessToken, newToken.RefreshToken, newToken.Expiry.Unix())
				updatedSource = &sourceCreds
				//updatedSource.AccessToken = newToken.AccessToken
				//updatedSource.ExpiresAt = newToken.Expiry.Unix()
				//// Don't overwrite `RefreshToken` with an empty value
				//// if this was a token refreshing request.
				//if newToken.RefreshToken != "" {
				//	updatedSource.RefreshToken = newToken.RefreshToken
				//}

			}
		}

		// OLD CODE
		httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

	} else {
		//Testing mode.
		httpClient = testHttpClient[0]
	}

	httpClient.Timeout = 10 * time.Second

	return &SourceClientBase{
		FastenEnv:        env,
		Context:          ctx,
		Logger:           globalLogger,
		OauthClient:      httpClient,
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
			//"MedicationRequest",
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
	}, updatedSource, nil
}

func (c *SourceClientBase) GetUsCoreResources() []string {
	return c.UsCoreResources
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// HttpClient
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientBase) GetRequest(resourceSubpathOrNext string, decodeModelPtr interface{}) error {
	resourceUrl, err := url.Parse(resourceSubpathOrNext)
	if err != nil {
		return err
	}
	if !resourceUrl.IsAbs() {
		resourceUrl, err = url.Parse(fmt.Sprintf("%s/%s", strings.TrimRight(c.SourceCredential.GetApiEndpointBaseUrl(), "/"), strings.TrimLeft(resourceSubpathOrNext, "/")))
	}
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, resourceUrl.String(), nil)
	if err != nil {
		return err
	}

	for key, val := range c.Headers {
		//req.Header.Add("Accept", "application/json+fhir")
		req.Header.Add(key, val)
	}

	//resp, err := c.OauthClient.Get(url)
	resp, err := c.OauthClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("An error occurred during request %s - %d - %s [%s]", resourceUrl, resp.StatusCode, resp.Status, string(b))
	}

	err = ParseBundle(resp.Body, decodeModelPtr)
	return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func ParseBundle(r io.Reader, decodeModelPtr interface{}) error {
	decoder := json.NewDecoder(r)
	//decoder.DisallowUnknownFields() //make sure we throw an error if unknown fields are present.
	err := decoder.Decode(decodeModelPtr)
	if err != nil {
		return err
	}
	return err
}
