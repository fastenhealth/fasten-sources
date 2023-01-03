package base

import (
	"crypto/tls"
	"fmt"
	"github.com/seborama/govcr"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"
	"time"
)

type bearerTransport struct {
	accessToken         string
	underlyingTransport http.RoundTripper
}

func (bt *bearerTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	req.Header.Add("X-Transaction-Id", time.Now().String())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bt.accessToken))
	return bt.underlyingTransport.RoundTrip(req)
}

func OAuthVcrSetup(t *testing.T, enableRecording bool) *http.Client {

	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, //disable certificate validation because we're playing back http requests.
	}
	insecureClient := http.Client{
		Transport: &bearerTransport{
			accessToken:         "PLACEHOLDER",
			underlyingTransport: tr,
		},
	}

	vcrConfig := govcr.VCRConfig{
		Logging:      true,
		CassettePath: path.Join("testdata", "govcr-fixtures"),
		Client:       &insecureClient,

		//this line ensures that we do not attempt to create new recordings.
		//Comment this out if you would like to make recordings.
		DisableRecording: !enableRecording,
	}

	// HTTP headers are case-insensitive
	vcrConfig.RequestFilters.Add(govcr.RequestDeleteHeaderKeys("User-Agent", "user-agent"))

	vcr := govcr.NewVCR(t.Name(), &vcrConfig)
	return vcr.Client
}

// helpers
func ReadTestFixture(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	return ioutil.ReadAll(jsonFile)
}
