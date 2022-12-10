package manual

import (
	"context"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/internal/base"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir401"
	fhir401utils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"github.com/fastenhealth/gofhir-models/fhir430"
	fhir430utils "github.com/fastenhealth/gofhir-models/fhir430/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

type ManualClient struct {
	FastenEnv  pkg.FastenEnvType
	SourceType pkg.SourceType
	Context    context.Context
	Logger     logrus.FieldLogger

	SourceCredential models.SourceCredential
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

func (m ManualClient) GetRequest(resourceSubpath string, decodeModelPtr interface{}) error {
	panic("implement me")
}

func (m ManualClient) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	panic("implement me")
}

func (m ManualClient) SyncAllBundle(db models.DatabaseRepository, bundleFile *os.File, bundleType string) (models.UpsertSummary, error) {
	summary := models.UpsertSummary{
		UpdatedResources: []string{},
	}

	// we need to parse the bundle into resources (might need to try a couple of different times)
	var rawResourceList []models.RawResourceFhir
	switch bundleType {
	case "fhir430":
		bundle430Data := fhir430.Bundle{}
		err := base.ParseBundle(bundleFile, &bundle430Data)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while parsing 4.3.0 bundle: %w", err)
		}
		client, _, err := base.GetSourceClientFHIR430(m.FastenEnv, m.Context, m.Logger, m.SourceCredential, http.DefaultClient)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while creating 4.3.0 client: %w", err)
		}
		rawResourceList, err = client.ProcessBundle(bundle430Data)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while processing 4.3.0 resources: %w", err)
		}

		//for _, apiModel := range rawResourceList {
		//	apiModel.ReferencedResources = client.ExtractReferencedResources(apiModel)
		//}
	case "fhir401":
		bundle401Data := fhir401.Bundle{}
		err := base.ParseBundle(bundleFile, &bundle401Data)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while parsing 4.0.1 bundle: %w", err)
		}
		client, _, err := base.GetSourceClientFHIR401(m.FastenEnv, m.Context, m.Logger, m.SourceCredential, http.DefaultClient)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while creating 4.0.1 client: %w", err)
		}
		rawResourceList, err = client.ProcessBundle(bundle401Data)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while processing 4.0.1 resources: %w", err)
		}

		for _, apiModel := range rawResourceList {
			apiModel.ReferencedResources = client.ExtractReferencedResources(apiModel)
		}

	}
	// we need to upsert all resources (and make sure they are associated with new Source)
	for _, apiModel := range rawResourceList {
		_, err := db.UpsertRawResource(m.Context, m.SourceCredential, apiModel)
		if err != nil {
			return summary, fmt.Errorf("an error occurred while upserting resources: %w", err)
		}
	}
	return summary, nil
}

func (m ManualClient) ExtractPatientId(bundleFile *os.File) (string, string, error) {
	// try from newest format to the oldest format
	bundle430Data := fhir430.Bundle{}
	bundle401Data := fhir401.Bundle{}

	var patientIds []string

	bundleType := "fhir430"
	if err := base.ParseBundle(bundleFile, &bundle430Data); err == nil {
		patientIds = lo.FilterMap[fhir430.BundleEntry, string](bundle430Data.Entry, func(bundleEntry fhir430.BundleEntry, _ int) (string, bool) {
			parsedResource, err := fhir430utils.MapToResource(bundleEntry.Resource, false)
			if err != nil {
				return "", false
			}
			typedResource := parsedResource.(models.ResourceInterface)
			resourceType, resourceId := typedResource.ResourceRef()

			if resourceId == nil || len(*resourceId) == 0 {
				return "", false
			}
			return *resourceId, resourceType == fhir430.ResourceTypePatient.String()
		})
	}
	bundleFile.Seek(0, io.SeekStart)

	//fallback
	if patientIds == nil || len(patientIds) == 0 {
		bundleType = "fhir401"
		//try parsing the bundle as a 401 bundle
		//TODO: find a better, more generic way to do this.
		err := base.ParseBundle(bundleFile, &bundle401Data)
		if err != nil {
			return "", "", err
		}
		patientIds = lo.FilterMap[fhir401.BundleEntry, string](bundle401Data.Entry, func(bundleEntry fhir401.BundleEntry, _ int) (string, bool) {
			parsedResource, err := fhir401utils.MapToResource(bundleEntry.Resource, false)
			if err != nil {
				return "", false
			}
			typedResource := parsedResource.(models.ResourceInterface)
			resourceType, resourceId := typedResource.ResourceRef()

			if resourceId == nil || len(*resourceId) == 0 {
				return "", false
			}
			return *resourceId, resourceType == fhir430.ResourceTypePatient.String()
		})
	}
	bundleFile.Seek(0, io.SeekStart)

	if patientIds == nil || len(patientIds) == 0 {
		return "", "", fmt.Errorf("could not determine patient id")
	} else {
		//reset reader

		return strings.TrimLeft(patientIds[0], "Patient/"), bundleType, nil
	}
}

func GetSourceClientManual(env pkg.FastenEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	return &ManualClient{
		FastenEnv:        env,
		Context:          ctx,
		Logger:           globalLogger,
		SourceCredential: sourceCreds,
	}, nil, nil
}
