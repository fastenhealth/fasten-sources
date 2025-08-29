package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fastenhealth/fasten-sources/clients/factory"
	clientModels "github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/definitions"
	"github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/fasten-sources/tools/test-smart-client/utils"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/proxy"
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
		RefreshToken    string `json:"refreshToken"`
		PatientId       string `json:"patientId"`
		ExpiresIn       int64  `json:"expiresIn"`
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

	proxyAddr := flag.String("proxy", "", "SOCKS5 proxy address to use. eg: socks5://socksproxy:socksproxy@localhost:1047")
	flag.Parse()
	if *proxyAddr != "" {
		log.Printf("proxyAddress: %s", *proxyAddr)
	}

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
	serverUrl := "http://localhost:9999"

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

		sourceClient, err := GenerateAuthClient(&requestData, proxyAddr, logger)
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

	http.HandleFunc("/api/source/syncall", func(res http.ResponseWriter, req *http.Request) {

		log.Printf("Source SyncAll Request: %s %v", req.Method, req.URL.Path)
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

		sourceClient, err := GenerateAuthClient(&requestData, proxyAddr, logger)
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while initializing hub client using source credential: %v", err), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{}

		storageRepo, err := utils.NewStorageRepository(logger)
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while initializing storage repository: %w", err), http.StatusInternalServerError)
			return
		}

		//Start Benchmarking
		start := time.Now()
		summary, err := sourceClient.SyncAll(storageRepo)
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while fetching data from source: %v", err), http.StatusInternalServerError)
			return
		}
		elapsed := time.Since(start)
		log.Printf("SyncAll took %s", elapsed)

		// wrtie test file
		err = storageRepo.Close()
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while closing storage repository: %w", err), http.StatusInternalServerError)
			return
		}

		response["elapsed"] = elapsed.String()
		response["total_records"] = summary.TotalResources
		response["errors"] = storageRepo.ErrorData

		//read the file and include the content in the reponse.
		data, err := os.ReadFile(storageRepo.LocalFilepath)
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while reading storage file: %v", err), http.StatusInternalServerError)
			return
		}
		response["data"] = string(data)

		responeJson, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			JSONError(res, fmt.Errorf("an error occurred while marshalling response: %v", err), http.StatusInternalServerError)
			return
		}

		res.Write(responeJson)
		return
	})

	http.HandleFunc("/api/source/cors/", CORSProxyHandler)

	go func() {
		log.Println("You will now be taken to your browser for authentication")
		time.Sleep(1 * time.Second)
		err := open.Run(serverUrl)
		if err != nil {
			log.Printf("an error occurred opening browser: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Printf("Authentication URL: %s\n", serverUrl)
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
	Scope                      string
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

func (s *fakeSourceCredential) SetTokens(accessToken string, refreshToken string, expiresAt int64, scope string) {
	s.AccessToken = accessToken
	s.RefreshToken = refreshToken
	s.ExpiresAt = expiresAt
	s.Scope = scope
}

func (s *fakeSourceCredential) GetSourceCredentialType() pkg.SourceCredentialType {
	return pkg.SourceCredentialTypeSmartOnFhir
}

// there are security implications to this, but we're only using this permissive proxy locally.
func CORSProxyHandler(proxyRes http.ResponseWriter, proxyReq *http.Request) {

	originalReqUrl := strings.TrimPrefix(proxyReq.URL.Path, "/api/source/cors/")

	//SECURITY: the proxy URL must start with the same URL as the endpoint.TokenUri
	corsUrl := fmt.Sprintf("https://%s", strings.TrimPrefix(originalReqUrl, "/"))

	remote, _ := url.Parse(corsUrl)
	remote.RawQuery = proxyReq.URL.Query().Encode()

	proxy := httputil.ReverseProxy{}
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = proxyReq.Header
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", remote.Host)
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		log.Printf("corsURL: %s\n remote.Path: %s\n Header: %v", corsUrl, remote.Path, req.Header)
		req.URL.Path = remote.Path
		req.Body = proxyReq.Body

		//TODO: throw an error if the remote.Host is not allowed

		reqDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			return
		}
		log.Printf("proxy req: %q", reqDump)
	}

	proxy.ModifyResponse = func(r *http.Response) error {
		//b, _ := ioutil.ReadAll(r.Body)
		//buf := bytes.NewBufferString("Monkey")
		//buf.Write(b)
		//r.Body = ioutil.NopCloser(buf)
		r.Header.Set("Access-Control-Allow-Methods", "GET,HEAD")
		r.Header.Set("Access-Control-Allow-Credentials", "true")
		r.Header.Set("Access-Control-Allow-Origin", "*")

		//dump, err := httputil.DumpResponse(r, true)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("DUMP RESPONSE: %q", dump)

		return nil
	}

	newProxyReq, _ := http.NewRequest(proxyReq.Method, corsUrl, proxyReq.Body)
	proxy.ServeHTTP(proxyRes, newProxyReq)
}

func GenerateAuthClient(requestData *ResourceRequest, proxyAddr *string, logger logrus.FieldLogger) (clientModels.SourceClient, error) {
	//populate a fake source credential
	sc := fakeSourceCredential{
		ClientId:   requestData.SourceDefinition.ClientId,
		PatientId:  requestData.RequestData.PatientId,
		EndpointId: requestData.SourceDefinition.Id,

		OauthAuthorizationEndpoint: requestData.SourceDefinition.AuthorizationEndpoint,
		OauthTokenEndpoint:         requestData.SourceDefinition.TokenEndpoint,
		ApiEndpointBaseUrl:         requestData.SourceDefinition.Url,
		RefreshToken:               requestData.RequestData.RefreshToken,
		AccessToken:                requestData.RequestData.AccessToken,
		ExpiresAt:                  time.Now().Add(time.Duration(requestData.RequestData.ExpiresIn) * time.Second).Unix(),
	}

	bgContext := context.WithValue(context.Background(), "AUTH_USERNAME", "temp")

	//check if a definition exists for this endpoint/platform
	if existingDef, err := definitions.GetSourceDefinition(definitions.WithEndpointId(requestData.SourceDefinition.Id)); err == nil && existingDef != nil {
		logger.Infof("found existing definition for %s: %v", requestData.SourceDefinition.Id, existingDef)
		//an exsting definition was found for this platform type, lets add some of the options to this custom definition
		requestData.SourceDefinition.ClientHeaders = existingDef.ClientHeaders
		requestData.SourceDefinition.MissingOpPatientEverything = existingDef.MissingOpPatientEverything
		requestData.SourceDefinition.CustomOpPatientEverything = existingDef.CustomOpPatientEverything
		requestData.SourceDefinition.ClientSupportedResources = existingDef.ClientSupportedResources
	} else {
		//TODO: SourceDefinition is always misisng ClientHeaders, lets set them ourselves
		requestData.SourceDefinition.ClientHeaders = map[string]string{
			"Accept": "application/json+fhir",
		}
	}

	clientOptions := []func(options *clientModels.SourceClientOptions){}

	//TODO: create a SOCKS Proxy Http Client for testing Quest
	// https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-3-socks-proxies/
	if *proxyAddr != "" {

		parsedProxyAddr, err := url.Parse(*proxyAddr)
		if err != nil {
			log.Fatalf("an error occurred while parsing proxy address: %v", err)
		}
		proxyPassword, hasProxyPassword := parsedProxyAddr.User.Password()
		if !hasProxyPassword {
			log.Fatalf("proxy address must have a password")
		}

		auth := proxy.Auth{
			User:     parsedProxyAddr.User.Username(),
			Password: proxyPassword,
		}
		dialer, err := proxy.SOCKS5("tcp", parsedProxyAddr.Host, &auth, nil)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{
			Transport: &http.Transport{
				Dial: dialer.Dial,
			},
		}
		clientOptions = append(clientOptions, clientModels.WithHttpClient(client))
	}

	sourceClient, err := factory.GetSourceClientWithDefinition(
		pkg.FastenLighthouseEnvProduction,
		bgContext,
		logger,
		&sc,
		&requestData.SourceDefinition,
		clientOptions...,
	)

	return sourceClient, err
}
