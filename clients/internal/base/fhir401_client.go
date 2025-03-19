package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/fastenhealth/fasten-sources/clients/models"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir401"
	fhirutils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

type SourceClientFHIR401 struct {
	*SourceClientBase
}

func GetSourceClientFHIR401(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, endpointDefinition *definitionsModels.LighthouseSourceDefinition, clientOptions ...func(options *models.SourceClientOptions)) (*SourceClientFHIR401, error) {
	baseClient, err := NewBaseClient(env, ctx, globalLogger, sourceCreds, endpointDefinition, clientOptions...)
	if err != nil {
		return nil, err
	}
	baseClient.FhirVersion = "4.0.1"
	return &SourceClientFHIR401{
		SourceClientBase: baseClient,
	}, err
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Sync
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR401) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	bundle, err := c.GetPatientBundle(c.SourceCredential.GetPatientId())
	if err != nil {
		return models.UpsertSummary{
			UpdatedResources: []string{},
		}, err
	}

	return c.SyncAllByPatientEverythingBundle(db, bundle)
}

// If the Patient/$everything or Patient/$export endpoints are supported, this function will allow us to easily process and add resources to the DB.
// This funciton should mimic SyncAllByResourceName logic
func (c *SourceClientFHIR401) SyncAllByPatientEverythingBundle(db models.DatabaseRepository, bundle interface{}) (models.UpsertSummary, error) {
	summary := models.UpsertSummary{
		UpdatedResources: []string{},
	}

	rawResourceModels, internalFragmentReferenceLookup, err := c.ProcessBundle(bundle.(fhir401.Bundle))
	if err != nil {
		c.Logger.Infof("An error occurred while processing patient bundle %s", c.SourceCredential.GetPatientId())
		return summary, err
	}
	summary.TotalResources = len(rawResourceModels)
	//error map storage.
	syncErrors := map[string]error{}

	//lookup table for every resource ID found by Fasten
	lookupResourceReferences := map[string]bool{}

	for ndx, apiModel := range rawResourceModels {
		err = c.ProcessResource(db, apiModel, lookupResourceReferences, internalFragmentReferenceLookup, &summary)
		if err != nil {
			syncErrors[apiModel.SourceResourceType] = err
			continue
		}

		if ndx%100 == 0 {
			db.BackgroundJobCheckpoint(c.Context,
				map[string]interface{}{
					"stage":          "EverythingBundle",
					"stage_progress": ndx,
					"summary":        summary,
				},
				map[string]interface{}{"errors": syncErrors},
			)
		}
	}

	//process any pending resources
	lookupResourceReferences, syncErrors = c.ProcessPendingResources(db, &summary, lookupResourceReferences, syncErrors)

	if len(syncErrors) > 0 {
		//TODO: ignore errors.
		c.Logger.Errorf("%d error(s) occurred during sync. \n %v", len(syncErrors), syncErrors)
	}

	db.BackgroundJobCheckpoint(c.Context,
		map[string]interface{}{
			"stage":   "PendingResources",
			"summary": summary,
		},
		map[string]interface{}{"errors": syncErrors},
	)
	return summary, nil
}

// SyncAllByResourceName
// Given a list of resource names (Patient, Encounter, etc), we should query for all resources associated with the current patient
// then store these resources in the database. NOTE: we must extract links to other resources referenced as we find them, and process
// them as well
// Changes to this function may need to be applied to SyncAllByPatientEverythingBundle as well.
func (c *SourceClientFHIR401) SyncAllByResourceName(db models.DatabaseRepository, resourceNames []string) (models.UpsertSummary, error) {
	summary := models.UpsertSummary{
		UpdatedResources: []string{},
	}

	//Store the Patient
	patientResource, err := c.GetPatient(c.SourceCredential.GetPatientId())
	if err != nil {
		return summary, err
	}
	patientJson, err := patientResource.MarshalJSON()
	if err != nil {
		return summary, err
	}

	patientResourceType, patientResourceId := patientResource.ResourceRef()
	patientResourceFhir := models.RawResourceFhir{
		SourceResourceType: patientResourceType,
		SourceResourceID:   *patientResourceId,
		ResourceRaw:        patientJson,
	}
	isUpdated, err := db.UpsertRawResource(c.Context, c.SourceCredential, patientResourceFhir)
	if err != nil {
		c.Logger.Infof("An error occurred while storing raw resource (by name) %v", err)
		return summary, err
	}
	summary.TotalResources += 1
	if isUpdated {
		summary.UpdatedResources = append(summary.UpdatedResources, fmt.Sprintf("%s/%s", patientResourceFhir.SourceResourceType, patientResourceFhir.SourceResourceID))
	}

	//error map storage.
	syncErrors := map[string]error{}

	//lookup table for every resource ID found by Fasten
	lookupResourceReferences := map[string]bool{}

	//query for resources by resource name
	resourceNames = lo.Uniq(resourceNames)
	sort.Strings(resourceNames)
	for _, resourceType := range resourceNames {
		bundle, err := c.GetResourceBundle(fmt.Sprintf("%s?patient=%s&_count=50", resourceType, c.SourceCredential.GetPatientId()))
		if err != nil {
			syncErrors[resourceType] = err
			continue
		}
		rawResourceModels, internalFragmentReferenceLookup, err := c.ProcessBundle(bundle.(fhir401.Bundle))
		if err != nil {
			c.Logger.Infof("An error occurred while processing %s bundle %s", resourceType, c.SourceCredential.GetPatientId())
			syncErrors[resourceType] = err
			continue
		}
		summary.TotalResources += len(rawResourceModels)

		for ndx, apiModel := range rawResourceModels {
			err = c.ProcessResource(db, apiModel, lookupResourceReferences, internalFragmentReferenceLookup, &summary)
			if err != nil {
				syncErrors[resourceType] = err
				continue
			}

			if ndx%100 == 0 {
				stageErrorData := map[string]interface{}{}
				if len(syncErrors) > 0 {
					stageErrorData["errors"] = syncErrors
				}
				db.BackgroundJobCheckpoint(c.Context,
					map[string]interface{}{
						"stage":          resourceType,
						"stage_progress": ndx,
						"summary":        summary,
					},
					stageErrorData,
				)
			}
		}

		checkpointErrorData := map[string]interface{}{}
		if len(syncErrors) > 0 {
			checkpointErrorData["errors"] = syncErrors
		}

		db.BackgroundJobCheckpoint(c.Context,
			map[string]interface{}{
				"stage":   resourceType,
				"summary": summary,
			},
			checkpointErrorData,
		)
	}

	//process any pending resources
	//if we have a populated whitelist, skip this -- we don't care about references to other resources, just the ones we queried. This is a naiive solution to prevent requesting unnecessary records. This should be improved.
	if len(c.GetResourceTypesAllowList()) == 0 {
		lookupResourceReferences, syncErrors = c.ProcessPendingResources(db, &summary, lookupResourceReferences, syncErrors)
	}

	checkpointErrorData := map[string]interface{}{}
	if len(syncErrors) > 0 {
		//ignore errors.
		c.Logger.Errorf("%d error(s) occurred during sync. \n %v", len(syncErrors), syncErrors)
		checkpointErrorData["errors"] = syncErrors
	}
	db.BackgroundJobCheckpoint(c.Context,
		map[string]interface{}{
			"stage":   "PendingResources",
			"summary": summary,
		},
		checkpointErrorData,
	)
	return summary, nil
}

func (c *SourceClientFHIR401) ProcessPendingResources(db models.DatabaseRepository, summary *models.UpsertSummary, lookupResourceReferences map[string]bool, syncErrors map[string]error) (map[string]bool, map[string]error) {
	// now that we've processed all resources by resource type, lets see if there's any extracted resources that we haven't processed.
	// NOTE: this is effectively a recursive operation since an extracted resource id may reference other resources.
	extractionLoopCount := 0
	for {
		//loop "forever" until we've processed all pending resources

		pendingLookupResourceReferences := lo.PickBy[string, bool](lookupResourceReferences, func(resourceId string, isCompleted bool) bool { return !isCompleted })
		pendingResourceReferences := lo.Keys(pendingLookupResourceReferences)

		//convert absolute urls to relative urls if possible
		for i, _ := range pendingResourceReferences {
			pendingResourceId := pendingResourceReferences[i]
			//convert absolute urls to relative url if possible
			if strings.HasPrefix(pendingResourceId, c.EndpointDefinition.Url) {
				pendingResourceReferences[i] = strings.TrimPrefix(strings.TrimPrefix(pendingResourceId, c.EndpointDefinition.Url), "/")
			}
		}

		c.Logger.Infof("Start processing Extracted Resource Identifiers: %d (%d loops completed)", len(pendingResourceReferences), extractionLoopCount)
		if len(pendingResourceReferences) == 0 {
			break
		}

		if extractionLoopCount > 10 {
			//bail out
			c.Logger.Warnf("we've attempted to extract resources more than 10 times. This should not happen.")
			break
		}

		//process pending resources
		summary.TotalResources += len(pendingResourceReferences)
		for ndx, pendingResourceIdOrUri := range pendingResourceReferences {
			var resourceRaw map[string]interface{}

			resourceSourceUri, err := c.GetRequest(pendingResourceIdOrUri, &resourceRaw)
			if err != nil {
				lookupResourceReferences[pendingResourceIdOrUri] = true //skip this failing resource
				c.Logger.Warnf("skipping resource (%s), request failed: %v", pendingResourceIdOrUri, err)
				continue
			}
			resourceRawJson, err := json.Marshal(resourceRaw)
			if err != nil {
				c.Logger.Warnf("skipping resource (%s), could not decode json: %v", pendingResourceIdOrUri, err)
				lookupResourceReferences[pendingResourceIdOrUri] = true //skip this failing resource
				continue
			}

			pendingResourceIdParts := []string{}
			if strings.HasPrefix(pendingResourceIdOrUri, "http://") || strings.HasPrefix(pendingResourceIdOrUri, "https://") {
				//this resource is an absolute url, and is most likely a Binary resource. Lets confirm, and throw an error if not.
				if pendingResourceType, ok := resourceRaw["resourceType"].(string); !ok || pendingResourceType != "Binary" {
					c.Logger.Errorf("[ERROR THIS SHOULD NOT HAPPEN] skipping absolute resource (%s), resource is not a Binary resource: %v", pendingResourceIdOrUri, err)
					lookupResourceReferences[pendingResourceIdOrUri] = true //skip this failing resource
					continue
				} else {
					//this is a binary resource, but it does not have an "id", so we need to generate one. (lookup will happen using sourceUri anyways)
					//TODO: if we use content addressable storage, we can use the hash of the resource as the id, however we can only store one SourceURI at a time, which could cause issues
					// so we'll generate a base64 encoded version of the sourceUri as the id
					pendingResourceIdParts = []string{pendingResourceType, base64.StdEncoding.EncodeToString([]byte(resourceSourceUri))}
				}
			} else {
				pendingResourceIdParts = strings.SplitN(pendingResourceIdOrUri, "/", 2)
			}

			pendingRawResource := models.RawResourceFhir{
				SourceResourceType: pendingResourceIdParts[0],
				SourceResourceID:   pendingResourceIdParts[1],
				ResourceRaw:        resourceRawJson,
				SourceUri:          resourceSourceUri,
			}

			//process resource will store the resource in the database, and potentially extract new resources we need to process.
			err = c.ProcessResource(db, pendingRawResource, lookupResourceReferences, map[string]string{}, summary)
			if err != nil {
				syncErrors[pendingResourceIdOrUri] = err
				continue
			}

			if ndx%100 == 0 {
				stageErrorData := map[string]interface{}{}
				if len(syncErrors) > 0 {
					stageErrorData["errors"] = syncErrors
				}
				db.BackgroundJobCheckpoint(c.Context,
					map[string]interface{}{
						"stage":          "PendingResources",
						"stage_progress": ndx,
						"summary":        summary,
					},
					stageErrorData,
				)
			}
		}
		extractionLoopCount += 1
	}
	return lookupResourceReferences, syncErrors
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FHIR
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR401) GetResourceBundle(relativeResourcePath string) (interface{}, error) {

	// https://www.hl7.org/fhir/patient-operation-everything.html
	bundle := fhir401.Bundle{}
	_, err := c.GetRequest(relativeResourcePath, &bundle)
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
		_, err := c.GetRequest(next, &nextBundle)
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

	//c.Logger.Debugf("BUNDLE - %v", bundle)
	return bundle, err

}

func (c *SourceClientFHIR401) GetPatientBundle(patientId string) (fhir401.Bundle, error) {
	/*
		Alternative bundle Urls:
		For those who have implemented FHIR based Argonaut or Bulk Data specifications, we expect common API patterns to include:

		GET {fhirBase}/Patient/{patientId}/$everything
		GET {fhirBase}/Patient/{patientId}/$export
		GET {fhirBase}/Patient/$export?patientId={patientId}
		GET {fhirBase}/Patient/{patientId}/$everything?_outputFormat=ndjson
		GET {fhirBase}/Patient/{patientId}/$export?_outputFormat=ndjson
		GET {fhirBase}/Patient/$export?patientId={patientId}&_outputFormat=ndjson
	*/

	patientBundle, err := c.GetResourceBundle(fmt.Sprintf("Patient/%s/$everything", patientId))
	if err != nil {
		return fhir401.Bundle{}, err
	}
	return patientBundle.(fhir401.Bundle), err
}

func (c *SourceClientFHIR401) GetPatient(patientId string) (fhir401.Patient, error) {

	patient := fhir401.Patient{}
	_, err := c.GetRequest(fmt.Sprintf("Patient/%s", patientId), &patient)
	return patient, fmt.Errorf("%w: %v", pkg.ErrResourcePatientFailure, err)
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Process Bundles
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *SourceClientFHIR401) ProcessBundle(bundle fhir401.Bundle) ([]models.RawResourceFhir, map[string]string, error) {

	//bundles may contain references to resources in one of 3 formats: - https://www.hl7.org/fhir/references.html#literal
	//
	// - an absolute URL
	// - a relative URL, which is relative to the Service Base URL, or, if processing a resource from a bundle, which is relative to the base URL implied by the Bundle.entry.fullUrl (see Resolving References in Bundles)
	// - an internal fragment reference (see "Contained Resources" below) -- eg. urn:uuid:c088b7af-fc41-43cc-ab80-4a9ab8d47cd9 or #c088b7af-fc41-43cc-ab80-4a9ab8d47cd9
	//
	// this last case is complicated, so we'll create a internal -> relative map that we can use when we find `urn:uuid:` references.
	internalFragmentReferenceLookup := map[string]string{}

	//process each entry in bundle
	wrappedResourceModels := lo.FilterMap[fhir401.BundleEntry, models.RawResourceFhir](bundle.Entry, func(bundleEntry fhir401.BundleEntry, _ int) (models.RawResourceFhir, bool) {
		originalResource, _ := fhirutils.MapToResource(bundleEntry.Resource, false)

		resourceType, resourceId := originalResource.(models.ResourceInterface).ResourceRef()
		if resourceId == nil {
			//no resourceId present for this resource, we'll ignore it.
			return models.RawResourceFhir{}, false
		}

		if bundleEntry.FullUrl != nil && (strings.HasPrefix(*bundleEntry.FullUrl, "urn:uuid:") || strings.HasPrefix(*bundleEntry.FullUrl, "#")) {
			internalFragmentReferenceLookup[*bundleEntry.FullUrl] = fmt.Sprintf("%s/%s", resourceType, *resourceId)
		}

		wrappedResourceModel := models.RawResourceFhir{
			SourceResourceID:   *resourceId,
			SourceResourceType: resourceType,
			ResourceRaw:        bundleEntry.Resource,
		}

		return wrappedResourceModel, true
	})
	return wrappedResourceModels, internalFragmentReferenceLookup, nil
}

// process a resource by:
// - inserting into the database
// - increment the updatedResources list if the resource has been updated
// - extract all external references from the resource payload (adding the the lookup table)
func (c *SourceClientFHIR401) ProcessResource(db models.DatabaseRepository, resource models.RawResourceFhir, referencedResourcesLookup map[string]bool, internalFragmentReferenceLookup map[string]string, summary *models.UpsertSummary) error {
	referencedResourcesLookup[fmt.Sprintf("%s/%s", resource.SourceResourceType, resource.SourceResourceID)] = true
	if len(resource.SourceUri) > 0 {
		//this is a reference to an external resource, we should mark that url as seen
		referencedResourcesLookup[resource.SourceUri] = true
	}

	resourceObj, err := fhirutils.MapToResource(resource.ResourceRaw, false)
	if err != nil {
		return err
	}

	resourceObjTyped := resourceObj.(models.ResourceInterface)
	currentResourceType, currentResourceId := resourceObjTyped.ResourceRef()

	//before processing this resource, we should check if it has any contained resources that we need to process first (recursively)
	containedResources := resourceObj.(models.ResourceInterface).ContainedResources()
	if containedResources != nil && len(containedResources) > 0 {
		for cndx, containedResource := range containedResources {
			containedResourceObj, err := fhirutils.MapToResource(containedResource, false)
			if err != nil {
				c.Logger.Warnf("Skipping contained resource (index %d) in %s/%s: %v", cndx, currentResourceType, *currentResourceId, err.Error())
				continue
			}

			containedResourceTyped := containedResourceObj.(models.ResourceInterface)
			containedResourceType, containedResourceId := containedResourceTyped.ResourceRef()
			if containedResourceId == nil {
				//no id present for this contained resource, we'll ignore it. (since theres no way to reference it anyways)
				c.Logger.Warnf("Skipping contained resource missing id: (%s/%s#%s index: %d)", currentResourceType, *currentResourceId, containedResourceType, cndx)
				continue
			}
			if len(c.GetResourceTypesAllowList()) > 0 && !lo.Contains(c.GetResourceTypesAllowList(), containedResourceType) {
				c.Logger.Warnf("Skipping contained resource not in allow list: (%s/%s#%s index: %d)", currentResourceType, *currentResourceId, containedResourceType, cndx)
				continue
			}

			normalizedContainedResourceId := normalizeContainedResourceId(currentResourceType, *currentResourceId, *containedResourceId)

			//generate a unique id for this contained resource by base64 url encoding this string
			base64ContainedResourceId := base64.URLEncoding.EncodeToString([]byte(normalizedContainedResourceId))

			//add this mapping to the internalFragmentReferenceLookup
			internalFragmentReferenceLookup[normalizedContainedResourceId] = fmt.Sprintf("%s/%s", containedResourceType, base64ContainedResourceId)

			containedResourceWrapped := models.RawResourceFhir{
				SourceResourceID:   base64ContainedResourceId,
				SourceResourceType: containedResourceType,
				ResourceRaw:        containedResource,
			}
			err = c.ProcessResource(db, containedResourceWrapped, referencedResourcesLookup, internalFragmentReferenceLookup, summary)
			if err != nil {
				return err
			}
		}
	}

	SourceClientFHIR401ExtractResourceMetadata(resourceObj, &resource, internalFragmentReferenceLookup)

	isUpdated, err := db.UpsertRawResource(c.Context, c.SourceCredential, resource)
	if err != nil {
		return err
	}
	if isUpdated {
		summary.UpdatedResources = append(summary.UpdatedResources, fmt.Sprintf("%s/%s", resource.SourceResourceType, resource.SourceResourceID))
	}

	for _, ref := range resource.ReferencedResources {
		if _, lookupOk := referencedResourcesLookup[ref]; !lookupOk {
			referencedResourcesLookup[ref] = false //this reference has not been seen before, set to false (pending)
		}
	}
	return nil
}
