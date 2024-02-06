package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/factory"
	"github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

/*
"sourceDefinition": this.sourceDefinition,
                        "requestData": {
                            "resourceType": resourceType,
                            "resourceRequest": resourceRequest,
                            "accessToken": this.accessToken,
                        },

*/

type ResourceRequest struct {
	SourceDefinition models.LighthouseSourceDefinition `json:"sourceDefinition"`
	RequestData      struct {
		ResourceType    string `json:"resourceType"`
		ResourceRequest string `json:"resourceRequest"`
		AccessToken     string `json:"accessToken"`
	} `json:"requestData"`
}

//go:embed html
var staticHtml embed.FS

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	log.Printf("%v", err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   err,
	})
	return
}
func main() {
	log.Printf("Starting oauth cli")
	defer log.Printf("Finished oauth cli")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%v", req.URL.Path)
		if strings.HasPrefix(req.URL.Path, "/callback") {
			callbackHtml, err := staticHtml.ReadFile("html/callback.html")
			if err != nil {
				log.Fatalf("error reading static files: %v", err)
			}
			res.Write(callbackHtml)
			//http.ServeFile(res, req, "html/callback.html")
		} else {
			indexHtml, err := staticHtml.ReadFile("html/index.html")
			if err != nil {
				log.Fatalf("error reading static files: %v", err)
			}
			res.Write(indexHtml)

			//http.ServeFile(res, req, "html/index.html")
		}
	})
	url := "http://localhost:9999"

	http.HandleFunc("/api/source/request", func(res http.ResponseWriter, req *http.Request) {

		log.Printf("Source Request: %s %v", req.Method, req.URL.Path)
		//write simple json response
		res.Header().Set("Content-Type", "application/json")

		//read post body (in json) and unmarshal into a struct
		var requestData ResourceRequest
		err := json.NewDecoder(req.Body).Decode(&requestData)
		if err != nil {

			JSONError(res, fmt.Errorf("error decoding request body: %v", err), http.StatusBadRequest)
			httputil.DumpRequest(req, true)
			return
		}
		log.Printf("%v", requestData)
		logger := logrus.WithField("callback", requestData.SourceDefinition.PlatformType)

		//populate a fake source credential
		sc := fakeSourceCredential{
			ClientId:   requestData.SourceDefinition.ClientId,
			PatientId:  "",
			EndpointId: requestData.SourceDefinition.Id,

			OauthAuthorizationEndpoint: requestData.SourceDefinition.AuthorizationEndpoint,
			OauthTokenEndpoint:         requestData.SourceDefinition.TokenEndpoint,
			ApiEndpointBaseUrl:         requestData.SourceDefinition.Url,
			RefreshToken:               "",
			AccessToken:                requestData.RequestData.AccessToken,
			ExpiresAt:                  time.Now().Add(1 * time.Hour).Unix(),
		}

		bgContext := context.WithValue(context.Background(), "AUTH_USERNAME", "temp")

		//TODO: SourceDefinition is always misisng ClientHeaders, lets set them ourselves
		requestData.SourceDefinition.ClientHeaders = map[string]string{
			"Accept": "application/json+fhir",
		}

		sourceClient, err := factory.GetSourceClientWithDefinition(
			pkg.FastenLighthouseEnvProduction,
			bgContext,
			logger,
			&sc,
			&requestData.SourceDefinition,
		)
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while initializing hub client using source credential: %v", err), http.StatusInternalServerError)
			return
		}

		var response map[string]interface{}

		_, err = sourceClient.GetRequest(fmt.Sprintf("%s/%s", requestData.RequestData.ResourceType, strings.TrimLeft(requestData.RequestData.ResourceRequest, "/")), &response)
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while fetching data from source: %v", err), http.StatusInternalServerError)
			return
		}

		responeJson, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while marshalling response: %v", err), http.StatusInternalServerError)
			return
		}

		res.Write(responeJson)
		return
	})

	go func() {
		log.Println("You will now be taken to your browser for authentication")
		time.Sleep(1 * time.Second)
		err := open.Run(url)
		if err != nil {
			log.Printf("an error occurred opening browser: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Printf("Authentication URL: %s\n", url)
	}()

	log.Fatal(http.ListenAndServe(":9999", nil))

}

// implements model.fakeSourceCredential
type fakeSourceCredential struct {
	EndpointId string
	PortalId   string
	BrandId    string

	PlatformType               pkg.PlatformType
	ClientId                   string
	PatientId                  string
	OauthAuthorizationEndpoint string
	OauthTokenEndpoint         string
	ApiEndpointBaseUrl         string
	RefreshToken               string
	AccessToken                string
	ExpiresAt                  int64
}

func (s *fakeSourceCredential) GetSourceId() string {
	return "fake-source-id"
}

func (s *fakeSourceCredential) GetEndpointId() string {
	return s.EndpointId
}

func (s *fakeSourceCredential) GetPortalId() string {
	return s.PortalId
}

func (s *fakeSourceCredential) GetBrandId() string {
	return s.BrandId
}

func (s *fakeSourceCredential) GetPlatformType() pkg.PlatformType {
	return s.PlatformType
}

func (s *fakeSourceCredential) GetClientId() string {
	return s.ClientId
}

func (s *fakeSourceCredential) GetPatientId() string {
	return s.PatientId
}

func (s *fakeSourceCredential) GetOauthAuthorizationEndpoint() string {
	return s.OauthAuthorizationEndpoint
}

func (s *fakeSourceCredential) GetOauthTokenEndpoint() string {
	return s.OauthTokenEndpoint
}

func (s *fakeSourceCredential) GetApiEndpointBaseUrl() string {
	return s.ApiEndpointBaseUrl
}

func (s *fakeSourceCredential) GetRefreshToken() string {
	return s.RefreshToken
}

func (s *fakeSourceCredential) GetAccessToken() string {
	return s.AccessToken
}

func (s *fakeSourceCredential) GetExpiresAt() int64 {
	return s.ExpiresAt
}

func (s *fakeSourceCredential) SetTokens(accessToken string, refreshToken string, expiresAt int64) {
	s.AccessToken = accessToken
	s.RefreshToken = refreshToken
	s.ExpiresAt = expiresAt
}

func (s *fakeSourceCredential) IsDynamicClient() bool {
	return false
}

func (s *fakeSourceCredential) RefreshDynamicClientAccessToken() error {
	return nil
}
