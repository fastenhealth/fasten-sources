package manual

import (
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/internal/base"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/gofhir-models/fhir401"
	fhir401utils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"github.com/fastenhealth/gofhir-models/fhir430"
	fhir430utils "github.com/fastenhealth/gofhir-models/fhir430/utils"
	"github.com/samber/lo"
	"io"
	"log"
	"os"
)

func extractPatientIdBundle(bundleFile *os.File) (string, pkg.FhirVersion, error) {
	defer bundleFile.Seek(0, io.SeekStart)

	// TODO: find a way to correctly detect bundle version

	bundleType := pkg.FhirVersion401
	patientIds, err := parse401Bundle(bundleFile)
	bundleFile.Seek(0, io.SeekStart)

	//fallback to 430 bundle
	if err != nil || patientIds == nil || len(patientIds) == 0 {
		log.Printf("failed to parse bundle or extract patientIds as 401, trying as 430: %v", err)
		bundleType = pkg.FhirVersion430

		patientIds, err = parse430Bundle(bundleFile)

	}
	if err != nil {
		//failed to parse the bundle as 401 and 430, return an error
		return "", "", fmt.Errorf("could not determine bundle version: %v", err)
	} else if patientIds == nil || len(patientIds) == 0 {
		return "", "", fmt.Errorf("could not determine patient id")
	} else {
		return cleanPatientIdPrefix(patientIds[0]), bundleType, nil
	}
}

//TODO: find a better, more generic way to do this.

func parse401Bundle(bundleFile *os.File) ([]string, error) {
	bundle401Data := fhir401.Bundle{}
	//try parsing the bundle as a 401 bundle
	if err := base.UnmarshalJson(bundleFile, &bundle401Data); err == nil {
		patientIds := lo.FilterMap[fhir401.BundleEntry, string](bundle401Data.Entry, func(bundleEntry fhir401.BundleEntry, _ int) (string, bool) {
			parsedResource, err := fhir401utils.MapToResource(bundleEntry.Resource, false)
			if err != nil {
				log.Printf("failed to parse resource: %v", err)
				return "", false
			}
			typedResource := parsedResource.(models.ResourceInterface)
			resourceType, resourceId := typedResource.ResourceRef()

			if resourceId == nil || len(*resourceId) == 0 {
				return "", false
			}
			return *resourceId, resourceType == fhir401.ResourceTypePatient.String()
		})

		return patientIds, nil
	} else {
		return nil, err
	}

}

func parse430Bundle(bundleFile *os.File) ([]string, error) {
	bundle430Data := fhir430.Bundle{}
	//try parsing the bundle as a 430 bundle
	if err := base.UnmarshalJson(bundleFile, &bundle430Data); err == nil {
		patientIds := lo.FilterMap[fhir430.BundleEntry, string](bundle430Data.Entry, func(bundleEntry fhir430.BundleEntry, _ int) (string, bool) {
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
		return patientIds, nil
	} else {
		return nil, err
	}
}
