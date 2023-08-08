package base

import (
	"context"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir430"
	fhirutils "github.com/fastenhealth/gofhir-models/fhir430/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SourceClientFHIR430 struct {
	*SourceClientBase
}

func GetSourceClientFHIR430(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (*SourceClientFHIR430, error) {
	baseClient, err := NewBaseClient(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	if err != nil {
		return nil, err
	}
	baseClient.FhirVersion = "4.3.0"
	return &SourceClientFHIR430{
		baseClient,
	}, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FHIR
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR430) GetPatientEverything(patientId string) (*fhir430.Bundle, error) {

	// https://www.hl7.org/fhir/patient-operation-everything.html
	bundle := fhir430.Bundle{}
	err := c.GetRequest(fmt.Sprintf("Patient/%s/$everything", patientId), &bundle)
	return &bundle, err
}

func (c *SourceClientFHIR430) GetPatient(patientId string) (*fhir430.Patient, error) {

	patient := fhir430.Patient{}
	err := c.GetRequest(fmt.Sprintf("Patient/%s", patientId), &patient)
	return &patient, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Process Bundles
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR430) ProcessBundle(bundle fhir430.Bundle) ([]models.RawResourceFhir, error) {

	//process each entry in bundle
	wrappedResourceModels := lo.FilterMap[fhir430.BundleEntry, models.RawResourceFhir](bundle.Entry, func(bundleEntry fhir430.BundleEntry, _ int) (models.RawResourceFhir, bool) {
		originalResource, _ := fhirutils.MapToResource(bundleEntry.Resource, false)

		resourceType, resourceId := originalResource.(models.ResourceInterface).ResourceRef()
		if resourceId == nil {
			//no resourceId present for this resource, we'll ignore it.
			return models.RawResourceFhir{}, false
		}
		// TODO find a way to safely/consistently get the resource updated date (and other metadata) which shoudl be added to the model.
		//if originalResource.Meta != nil && originalResource.Meta.LastUpdated != nil {
		//	if parsed, err := time.Parse(time.RFC3339Nano, *originalResource.Meta.LastUpdated); err == nil {
		//		patientProfile.UpdatedAt = parsed
		//	}
		//}

		wrappedResourceModel := models.RawResourceFhir{
			SourceResourceID:   *resourceId,
			SourceResourceType: resourceType,
			ResourceRaw:        bundleEntry.Resource,
		}

		return wrappedResourceModel, true
	})
	return wrappedResourceModels, nil
}
