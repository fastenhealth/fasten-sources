package manual

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/internal/base"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir401"
	fhir401utils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

type ManualClient struct {
	FastenEnv  pkg.FastenLighthouseEnvType
	SourceType pkg.SourceType
	Context    context.Context
	Logger     logrus.FieldLogger

	SourceCredential models.SourceCredential
}

func (m ManualClient) GetSourceCredential() models.SourceCredential {
	return m.SourceCredential
}

func (m ManualClient) GetResourceBundle(relativeResourcePath string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m ManualClient) SyncAllByPatientEverythingBundle(db models.DatabaseRepository, bundleModel interface{}) (models.UpsertSummary, error) {
	//TODO implement me
	panic("implement me")
}

func (m ManualClient) GetUsCoreResources() []string {
	//TODO implement me
	panic("implement me")
}

func (m ManualClient) SyncAllByResourceName(db models.DatabaseRepository, resourceNames []string) (models.UpsertSummary, error) {
	//TODO implement me
	panic("implement me")
}

func (m ManualClient) GetRequest(resourceSubpath string, decodeModelPtr interface{}) (string, error) {
	panic("implement me")
}

func (m ManualClient) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	panic("implement me")
}

func (m ManualClient) SyncAllBundle(db models.DatabaseRepository, bundleFile *os.File, bundleType pkg.FhirVersion) (models.UpsertSummary, error) {
	//structurally similar to #SyncAllByResourceName in clients/internal/base/fhir401_client.go
	summary := models.UpsertSummary{
		UpdatedResources: []string{},
	}

	// we need to parse the bundle into resources (might need to try a couple of different times)
	var rawResourceList []models.RawResourceFhir

	//error map storage.
	syncErrors := map[string]error{}

	//lookup table for every resource ID found by Fasten
	lookupResourceReferences := map[string]bool{}

	//lookup table for bundle internal references -> relative references (used heavily by file Bundles)
	internalFragmentReferenceLookup := map[string]string{}

	//retrieve the FHIR client
	client, err := base.GetSourceClientFHIR401(m.FastenEnv, m.Context, m.Logger, m.SourceCredential, http.DefaultClient)
	if err != nil {
		return summary, fmt.Errorf("an error occurred while creating 4.0.1 client: %w", err)
	}

	//parse the document
	documentType, err := GetFileDocumentType(bundleFile)
	if err != nil {
		return summary, err
	}
	m.Logger.Infof("Begin processing document: %s", documentType)
	switch documentType {
	case pkg.DocumentTypeFhirBundle:

		bundle401Data := fhir401.Bundle{}
		err := base.UnmarshalJson(bundleFile, &bundle401Data)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while parsing 4.0.1 bundle: %w", err)
		}

		rawResourceList, internalFragmentReferenceLookup, err = client.ProcessBundle(bundle401Data)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while processing 4.0.1 resources: %w", err)
		}

		for _, apiModel := range rawResourceList {
			err = client.ProcessResource(db, apiModel, lookupResourceReferences, internalFragmentReferenceLookup, &summary)
			if err != nil {
				syncErrors[apiModel.SourceResourceType] = err
				continue
			}
		}

	case pkg.DocumentTypeFhirNDJSON:
		d := json.NewDecoder(bundleFile)
		counter := 0
		for {

			var resource json.RawMessage
			err := d.Decode(&resource)
			if err != nil {
				// io.EOF is expected at end of stream.
				if err == io.EOF {
					break //we're done
				} else {
					continue //skip this document, invalid json
				}
			}

			resourceObj, err := fhir401utils.MapToResource(resource, false)
			if err != nil {
				syncErrors[fmt.Sprintf("index: %d", counter)] = err
				continue
			}

			resourceObjTyped := resourceObj.(models.ResourceInterface)
			resourceType, resourceId := resourceObjTyped.ResourceRef()
			if resourceId == nil {
				syncErrors[fmt.Sprintf("%s (index: %d)", resourceType, counter)] = fmt.Errorf("resource ID is nil, skipping")
				continue
			}

			apiModel := models.RawResourceFhir{
				SourceResourceID:   *resourceId,
				SourceResourceType: resourceType,
				ResourceRaw:        resource,
			}
			rawResourceList = append(rawResourceList, apiModel)
			err = client.ProcessResource(db, apiModel, lookupResourceReferences, internalFragmentReferenceLookup, &summary)
			if err != nil {
				syncErrors[apiModel.SourceResourceType] = err
				continue
			}

			counter += 1
		}
	}
	summary.TotalResources = len(rawResourceList)

	m.Logger.Infof("Completed document processing: %d resources", summary.TotalResources)

	if len(syncErrors) > 0 {
		//TODO: ignore errors.
		m.Logger.Errorf("%d error(s) occurred during sync. \n %v", len(syncErrors), syncErrors)
	}
	return summary, nil
}

func (m ManualClient) ExtractPatientId(bundleFile *os.File) (string, pkg.FhirVersion, error) {
	documentType, err := GetFileDocumentType(bundleFile)
	if err != nil {
		return "", pkg.FhirVersion401, err
	}

	switch documentType {
	case pkg.DocumentTypeFhirBundle:
		return extractPatientIdBundle(bundleFile)
	case pkg.DocumentTypeFhirNDJSON:
		return extractPatientIdNDJson(bundleFile)
	default:
		return "", pkg.FhirVersion401, fmt.Errorf("unsupported document type: %s", documentType)
	}
}

func GetSourceClientManual(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, error) {
	return &ManualClient{
		FastenEnv:        env,
		Context:          ctx,
		Logger:           globalLogger,
		SourceCredential: sourceCreds,
	}, nil
}

func cleanPatientIdPrefix(patientId string) string {
	return strings.TrimLeft(patientId, "Patient/")
}
