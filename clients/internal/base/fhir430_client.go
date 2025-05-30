package base

import (
	"context"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir430"
	fhirutils "github.com/fastenhealth/gofhir-models/fhir430/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

type SourceClientFHIR430 struct {
	*SourceClientBase
}

func GetSourceClientFHIR430(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, endpointDefinition *definitionsModels.LighthouseSourceDefinition, clientOptions ...func(options *models.SourceClientOptions)) (*SourceClientFHIR430, error) {
	baseClient, err := NewBaseClient(env, ctx, globalLogger, sourceCreds, endpointDefinition, clientOptions...)
	if err != nil {
		return nil, err
	}
	baseClient.FhirVersion = "4.3.0"
	return &SourceClientFHIR430{
		baseClient,
	}, err
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FHIR
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR430) GetPatientEverything(patientId string) (*fhir430.Bundle, error) {

	// https://www.hl7.org/fhir/patient-operation-everything.html
	bundle := fhir430.Bundle{}
	_, err := c.GetRequest(fmt.Sprintf("Patient/%s/$everything", patientId), &bundle)
	return &bundle, err
}

func (c *SourceClientFHIR430) GetPatient(patientId string) (*fhir430.Patient, error) {

	patient := fhir430.Patient{}
	_, err := c.GetRequest(fmt.Sprintf("Patient/%s", patientId), &patient)
	if err != nil {
		return &patient, fmt.Errorf("%w: %v", pkg.ErrResourcePatientFailure, err)
	}
	return &patient, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Process Bundles
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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
