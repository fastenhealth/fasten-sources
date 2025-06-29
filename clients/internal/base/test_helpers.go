package base

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	govcr "github.com/seborama/govcr/v15"
	"github.com/seborama/govcr/v15/cassette/track"
)

type bearerTransport struct {
	accessToken         string
	underlyingTransport http.RoundTripper
}

func (bt *bearerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", "application/fhir+json")
	req.Header.Add("X-Transaction-Id", time.Now().String())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bt.accessToken))
	return bt.underlyingTransport.RoundTrip(req)
}

func OAuthVcrSetup(t *testing.T, enableRecording bool, accessToken ...string) *http.Client {

	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, //disable certificate validation because we're playing back http requests.
	}

	customTransport := bearerTransport{}
	customTransport.underlyingTransport = tr
	if enableRecording && len(accessToken) == 1 {
		customTransport.accessToken = accessToken[0]
	} else {
		customTransport.accessToken = "PLACEHOLDER"
	}

	insecureClient := http.Client{
		Transport: &customTransport,
	}

	//this line ensures that we do not attempt to create new recordings.
	//Set enableRecording if you would like to make recordings.
	recordMode := govcr.WithOfflineMode()
	if enableRecording {
		recordMode = govcr.WithLiveOnlyMode()
	}

	vcrConfig := govcr.NewVCR(
		govcr.NewCassetteLoader(path.Join("testdata", "govcr-fixtures", t.Name()+".cassette")),
		govcr.WithTrackRecordingMutators(track.ResponseDeleteTLS()),
		govcr.WithTrackReplayingMutators(track.ResponseDeleteTLS()),
		govcr.WithClient(&insecureClient),
		recordMode,
	)

	vcrConfig.SetRequestMatchers(
		govcr.DefaultMethodMatcher,
		govcr.DefaultURLMatcher,
		func(httpRequest, trackRequest *track.Request) bool {
			// we can safely mutate our inputs:
			// mutations affect other RequestMatcher's but _not_ the
			// original HTTP request or the cassette Tracks.
			httpRequest.Header.Del("X-Custom-Timestamp")
			trackRequest.Header.Del("X-Custom-Timestamp")

			// HTTP headers are case-insensitive
			httpRequest.Header.Del("User-Agent")
			trackRequest.Header.Del("User-Agent")
			httpRequest.Header.Del("user-agent")
			trackRequest.Header.Del("user-agent")

			return govcr.DefaultHeaderMatcher(httpRequest, trackRequest)
		},
	)

	return vcrConfig.HTTPClient()
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

func ToMap[K comparable, V any](sm *sync.Map) map[K]V {
	m := make(map[K]V)
	sm.Range(func(k, v any) bool {
		if key, ok := k.(K); ok {
			if value, ok := v.(V); ok {
				m[key] = value
			}
		}
		return true
	})
	return m
}
