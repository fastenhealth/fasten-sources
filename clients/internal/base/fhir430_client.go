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

type FHIR430Client struct {
	*BaseClient
}

func GetSourceClientFHIR430(env pkg.FastenEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (*FHIR430Client, *models.SourceCredential, error) {
	baseClient, updatedSource, err := NewBaseClient(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	return &FHIR430Client{
		baseClient,
	}, updatedSource, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FHIR
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *FHIR430Client) GetPatientEverything(patientId string) (*fhir430.Bundle, error) {

	// https://www.hl7.org/fhir/patient-operation-everything.html
	bundle := fhir430.Bundle{}
	err := c.GetRequest(fmt.Sprintf("Patient/%s/$everything", patientId), &bundle)
	return &bundle, err
}

func (c *FHIR430Client) GetPatient(patientId string) (*fhir430.Patient, error) {

	patient := fhir430.Patient{}
	err := c.GetRequest(fmt.Sprintf("Patient/%s", patientId), &patient)
	return &patient, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Process Bundles
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *FHIR430Client) ProcessBundle(bundle fhir430.Bundle) ([]models.ResourceInterface, error) {

	//process each entry in bundle
	wrappedResourceModels := lo.FilterMap[fhir430.BundleEntry, models.ResourceInterface](bundle.Entry, func(bundleEntry fhir430.BundleEntry, _ int) (models.ResourceInterface, bool) {
		originalResource, _ := fhirutils.MapToResource(bundleEntry.Resource, false)

		_, resourceId := originalResource.(models.ResourceInterface).ResourceRef()
		//return
		// TODO find a way to safely/consistently get the resource updated date (and other metadata) which shoudl be added to the model.
		//if originalResource.Meta != nil && originalResource.Meta.LastUpdated != nil {
		//	if parsed, err := time.Parse(time.RFC3339Nano, *originalResource.Meta.LastUpdated); err == nil {
		//		patientProfile.UpdatedAt = parsed
		//	}
		//}
		if resourceId == nil {
			//no resourceId present for this resource, we'll ignore it.
			return nil, false
		}

		return originalResource.(models.ResourceInterface), true

		//wrappedResourceModel := models.ResourceFhir{
		//	OriginBase: models.OriginBase{
		//		ModelBase:          models.ModelBase{},
		//		UserID:             c.Source.UserID,
		//		SourceID:           c.Source.ID,
		//		SourceResourceID:   *resourceId,
		//		SourceResourceType: resourceType,
		//	},
		//	Payload: datatypes.JSON(bundleEntry.Resource),
		//}
		//
		//return wrappedResourceModel, true
	})
	return wrappedResourceModels, nil
}
