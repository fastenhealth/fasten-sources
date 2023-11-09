package manual

import (
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	fhir401utils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"io"
	"os"
)

func extractPatientIdNDJson(bundleFile *os.File) (string, pkg.FhirVersion, error) {
	// TODO: find a way to correctly detect bundle version
	defer bundleFile.Seek(0, io.SeekStart)

	patientIds := []string{}
	d := json.NewDecoder(bundleFile)
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
			continue
		}

		resourceObjTyped := resourceObj.(models.ResourceInterface)
		currentResourceType, currentResourceId := resourceObjTyped.ResourceRef()

		if currentResourceType == "Patient" {
			patientIds = append(patientIds, *currentResourceId)
		}
	}

	if len(patientIds) == 0 {
		return "", "", fmt.Errorf("could not determine patient id")
	} else {
		return cleanPatientIdPrefix(patientIds[0]), pkg.FhirVersion401, nil
	}
}
