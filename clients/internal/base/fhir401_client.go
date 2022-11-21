package base

import (
	"context"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir401"
	fhirutils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SourceClientFHIR401 struct {
	*SourceClientBase
}

func GetSourceClientFHIR401(env pkg.FastenEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (*SourceClientFHIR401, *models.SourceCredential, error) {
	baseClient, updatedSource, err := NewBaseClient(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	baseClient.FhirVersion = "4.0.1"
	return &SourceClientFHIR401{
		SourceClientBase: baseClient,
	}, updatedSource, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Sync
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR401) SyncAll(db models.DatabaseRepository) error {

	bundle, err := c.GetPatientBundle(c.SourceCredential.GetPatientId())
	if err != nil {
		return err
	}

	rawResourceModels, err := c.ProcessBundle(bundle)
	if err != nil {
		c.Logger.Infof("An error occurred while processing patient bundle %s", c.SourceCredential.GetPatientId())
		return err
	}
	//todo, create the resources in dependency order

	for _, rawResource := range rawResourceModels {
		err = db.UpsertRawResource(context.Background(), c.SourceCredential, rawResource)
		if err != nil {
			return err
		}
	}
	return nil
}

//TODO, find a way to sync references that cannot be searched by patient ID.
func (c *SourceClientFHIR401) SyncAllByResourceName(db models.DatabaseRepository, resourceNames []string) error {

	//Store the Patient
	patientResource, err := c.GetPatient(c.SourceCredential.GetPatientId())
	if err != nil {
		return err
	}
	patientJson, err := patientResource.MarshalJSON()
	if err != nil {
		return err
	}

	patientResourceType, patientResourceId := patientResource.ResourceRef()
	patientResourceFhir := models.RawResourceFhir{
		SourceResourceType: patientResourceType,
		SourceResourceID:   *patientResourceId,
		RawResource:        patientJson,
	}
	err = db.UpsertRawResource(context.Background(), c.SourceCredential, patientResourceFhir)
	if err != nil {
		c.Logger.Infof("An error occurred while storing raw resource (by name) %v", err)
		return err
	}
	//error map storage.
	syncErrors := map[string]error{}

	//Store all other resources.
	for _, resourceType := range resourceNames {
		bundle, err := c.GetResourceBundle(fmt.Sprintf("%s?patient=%s", resourceType, c.SourceCredential.GetPatientId()))
		if err != nil {
			syncErrors[resourceType] = err
			continue
		}
		rawResourceModels, err := c.ProcessBundle(bundle)
		if err != nil {
			c.Logger.Infof("An error occurred while processing %s bundle %s", resourceType, c.SourceCredential.GetPatientId())
			syncErrors[resourceType] = err
			continue
		}
		for _, apiModel := range rawResourceModels {
			err = db.UpsertRawResource(context.Background(), c.SourceCredential, apiModel)
			if err != nil {
				syncErrors[resourceType] = err
				continue
			}
		}
	}

	if len(syncErrors) > 0 {
		return fmt.Errorf("%d error(s) occurred during sync: %v", len(syncErrors), syncErrors)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FHIR
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR401) GetResourceBundle(relativeResourcePath string) (fhir401.Bundle, error) {

	// https://www.hl7.org/fhir/patient-operation-everything.html
	bundle := fhir401.Bundle{}
	err := c.GetRequest(relativeResourcePath, &bundle)
	if err != nil {
		return bundle, err
	}
	var next string
	var prev string
	var self string
	for _, link := range bundle.Link {
		if link.Relation == "next" {
			next = link.Url
		} else if link.Relation == "self" {
			self = link.Url
		} else if link.Relation == "previous" {
			prev = link.Url
		}
	}

	for len(next) > 0 && next != self && next != prev {
		c.Logger.Debugf("Paginated request => %s", next)
		nextBundle := fhir401.Bundle{}
		err := c.GetRequest(next, &nextBundle)
		if err != nil {
			return bundle, nil //ignore failures when paginating?
		}
		bundle.Entry = append(bundle.Entry, nextBundle.Entry...)

		next = "" //reset the pointers
		self = ""
		prev = ""
		for _, link := range nextBundle.Link {
			if link.Relation == "next" {
				next = link.Url
			} else if link.Relation == "self" {
				self = link.Url
			} else if link.Relation == "previous" {
				prev = link.Url
			}
		}
	}

	c.Logger.Infof("BUNDLE - %v", bundle)
	return bundle, err

}

func (c *SourceClientFHIR401) GetPatientBundle(patientId string) (fhir401.Bundle, error) {
	return c.GetResourceBundle(fmt.Sprintf("Patient/%s/$everything", patientId))
}

func (c *SourceClientFHIR401) GetPatient(patientId string) (fhir401.Patient, error) {

	patient := fhir401.Patient{}
	err := c.GetRequest(fmt.Sprintf("Patient/%s", patientId), &patient)
	return patient, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Process Bundles
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR401) ProcessBundle(bundle fhir401.Bundle) ([]models.RawResourceFhir, error) {

	//process each entry in bundle
	wrappedResourceModels := lo.FilterMap[fhir401.BundleEntry, models.RawResourceFhir](bundle.Entry, func(bundleEntry fhir401.BundleEntry, _ int) (models.RawResourceFhir, bool) {
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
			RawResource:        bundleEntry.Resource,
		}

		return wrappedResourceModel, true
	})
	return wrappedResourceModels, nil
}
