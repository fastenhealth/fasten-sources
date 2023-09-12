package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/factory"
	defFactory "github.com/fastenhealth/fasten-sources/definitions/factory"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

type ResourceRequest struct {
	ResourceType    string `json:"resourceType"`
	ResourceRequest string `json:"resourceRequest"`
	SourceMode      string `json:"sourceMode"`
	SourceType      string `json:"sourceType"`
	AccessToken     string `json:"accessToken"`
	ClientId        string `json:"clientId"`
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%v", req.URL.Path)
		if strings.HasPrefix(req.URL.Path, "/callback") {
			http.ServeFile(res, req, "html/callback.html")
		} else {
			http.ServeFile(res, req, "html/index.html")
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
			log.Printf("Error decoding request body: %v", err)
			httputil.DumpRequest(req, true)
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("%v", requestData)
		logger := logrus.WithField("callback", requestData.SourceType)

		//get the source defintiion
		sourceConfig, err := defFactory.GetSourceConfig(pkg.FastenLighthouseEnvType(requestData.SourceMode), pkg.SourceType(requestData.SourceType), map[pkg.SourceType]string{})
		if err != nil {
			logger.Errorln("An error occurred while initializing source config", err)
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		//populate a fake source credential
		sc := fakeSourceCredential{
			SourceType:                 pkg.SourceType(requestData.SourceType),
			ClientId:                   requestData.ClientId,
			PatientId:                  "",
			OauthAuthorizationEndpoint: sourceConfig.AuthorizationEndpoint,
			OauthTokenEndpoint:         sourceConfig.TokenEndpoint,
			ApiEndpointBaseUrl:         sourceConfig.ApiEndpointBaseUrl,
			RefreshToken:               "",
			AccessToken:                requestData.AccessToken,
			ExpiresAt:                  time.Now().Add(1 * time.Hour).Unix(),
		}

		bgContext := context.WithValue(context.Background(), "AUTH_USERNAME", "temp")

		sourceClient, err := factory.GetSourceClient(pkg.FastenLighthouseEnvType(requestData.SourceMode), pkg.SourceType(requestData.SourceType), bgContext, logger, &sc)
		if err != nil {
			logger.Errorln("An error occurred while initializing hub client using source credential", err)
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		var response map[string]interface{}

		_, err = sourceClient.GetRequest(fmt.Sprintf("%s/%s", requestData.ResourceType, strings.TrimLeft(requestData.ResourceRequest, "/")), &response)
		if err != nil {
			logger.Errorln("An error occurred while fetching data from source", err)
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		responeJson, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			logger.Errorln("An error occurred while marshalling response", err)
			http.Error(res, err.Error(), http.StatusBadRequest)
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
			log.Fatalf("an error occurred opening browser: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Printf("Authentication URL: %s\n", url)
	}()

	log.Fatal(http.ListenAndServe(":9999", nil))

}

// implements model.fakeSourceCredential
type fakeSourceCredential struct {
	SourceType                 pkg.SourceType
	ClientId                   string
	PatientId                  string
	OauthAuthorizationEndpoint string
	OauthTokenEndpoint         string
	ApiEndpointBaseUrl         string
	RefreshToken               string
	AccessToken                string
	ExpiresAt                  int64
}

func (s *fakeSourceCredential) GetSourceType() pkg.SourceType {
	return s.SourceType
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
