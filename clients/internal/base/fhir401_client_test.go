package base

import (
	"context"
	"encoding/json"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/gofhir-models/fhir401"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

// helpers
func readTestFixture(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	return ioutil.ReadAll(jsonFile)
}

func TestNewFHIR401Client(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})

	//test
	client, _, err := GetSourceClientFHIR401(context.Background(), fakeConfig, testLogger, models.SourceCredential{
		RefreshToken: "test-refresh-token",
		AccessToken:  "test-access-token",
	}, &http.Client{})

	//assert
	require.NoError(t, err)
	require.Equal(t, client.Source.AccessToken, "test-access-token")
	require.Equal(t, client.Source.RefreshToken, "test-refresh-token")
}

func TestFHIR401Client_ProcessBundle(t *testing.T) {
	t.Parallel()
	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeConfig := mock_config.NewMockInterface(mockCtrl)
	testLogger := logrus.WithFields(logrus.Fields{
		"type": "test",
	})
	client, _, err := GetSourceClientFHIR401(context.Background(), fakeConfig, testLogger, models.Source{
		RefreshToken: "test-refresh-token",
		AccessToken:  "test-access-token",
	}, &http.Client{})
	require.NoError(t, err)

	jsonBytes, err := readTestFixture("testdata/fixtures/401-R4/bundle/cigna_syntheticuser05-everything.json")
	require.NoError(t, err)
	var bundle fhir401.Bundle
	err = json.Unmarshal(jsonBytes, &bundle)
	require.NoError(t, err)

	// test
	wrappedResourceModels, err := client.ProcessBundle(bundle)
	//log.Printf("%v", wrappedResourceModels)

	//assert
	require.NoError(t, err)
	require.Equal(t, 11, len(wrappedResourceModels))
	//require.Equal(t, "A00000000000005", profile.SourceResourceID)
}
