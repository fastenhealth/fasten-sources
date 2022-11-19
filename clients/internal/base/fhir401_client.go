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

type FHIR401Client struct {
	*BaseClient
}

func GetSourceClientFHIR401(env pkg.FastenEnvType, sourceType pkg.SourceType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (*FHIR401Client, *models.SourceCredential, error) {
	baseClient, updatedSource, err := NewBaseClient(env, sourceType, ctx, globalLogger, sourceCreds, testHttpClient...)
	return &FHIR401Client{
		baseClient,
	}, updatedSource, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Sync
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *FHIR401Client) SyncAll(db models.DatabaseRepository) error {

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
func (c *FHIR401Client) SyncAllByResourceName(db models.DatabaseRepository, resourceNames []string) error {

	//Store the Patient
	patientResource, err := c.GetPatient(c.SourceCredential.GetPatientId())
	if err != nil {
		return err
	}
	//patientJson, err := patientResource.MarshalJSON()
	//if err != nil {
	//	return err
	//}

	//patientResourceType, patientResourceId := patientResource.ResourceRef()
	//patientResourceFhir := models.ResourceFhir{
	//	OriginBase: models.OriginBase{
	//		UserID:             c.Source.UserID,
	//		SourceID:           c.Source.ID,
	//		SourceResourceType: patientResourceType,
	//		SourceResourceID:   *patientResourceId,
	//	},
	//	Payload: datatypes.JSON(patientJson),
	//}
	db.UpsertRawResource(context.Background(), c.SourceCredential, patientResource)

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
func (c *FHIR401Client) GetResourceBundle(relativeResourcePath string) (fhir401.Bundle, error) {

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

	return bundle, err

}

func (c *FHIR401Client) GetPatientBundle(patientId string) (fhir401.Bundle, error) {
	return c.GetResourceBundle(fmt.Sprintf("Patient/%s/$everything", patientId))
}

func (c *FHIR401Client) GetPatient(patientId string) (fhir401.Patient, error) {

	patient := fhir401.Patient{}
	err := c.GetRequest(fmt.Sprintf("Patient/%s", patientId), &patient)
	return patient, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Process Bundles
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *FHIR401Client) ProcessBundle(bundle fhir401.Bundle) ([]models.ResourceInterface, error) {

	//process each entry in bundle
	wrappedResourceModels := lo.FilterMap[fhir401.BundleEntry, models.ResourceInterface](bundle.Entry, func(bundleEntry fhir401.BundleEntry, _ int) (models.ResourceInterface, bool) {
		originalResource, _ := fhirutils.MapToResource(bundleEntry.Resource, false)
		return originalResource.(models.ResourceInterface), true

		//resourceType, resourceId := originalResource.(ResourceInterface).ResourceRef()

		// TODO find a way to safely/consistently get the resource updated date (and other metadata) which shoudl be added to the model.
		//if originalResource.Meta != nil && originalResource.Meta.LastUpdated != nil {
		//	if parsed, err := time.Parse(time.RFC3339Nano, *originalResource.Meta.LastUpdated); err == nil {
		//		patientProfile.UpdatedAt = parsed
		//	}
		//}
		//if resourceId == nil {
		//	//no resourceId present for this resource, we'll ignore it.
		//	return models.ResourceFhir{}, false
		//}

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

		//return wrappedResourceModel, true
	})
	return wrappedResourceModels, nil
}
