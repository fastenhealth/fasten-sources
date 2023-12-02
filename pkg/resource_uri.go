package pkg

import (
	"fmt"
	"strings"
)

const FASTENHEALTH_URN_PREFIX = "urn:fastenhealth-fhir:"

func ParseReferenceUri(referenceUri *string) (string, string, string, error) {
	if referenceUri == nil || *referenceUri == "" {
		return "", "", "", fmt.Errorf("reference cannot be empty nil")
	}
	if strings.HasPrefix(*referenceUri, FASTENHEALTH_URN_PREFIX) {
		//parse referenceUri into sourceId, resourceType, resourceId
		originalReference := strings.TrimPrefix(*referenceUri, FASTENHEALTH_URN_PREFIX)
		urnParts := strings.Split(originalReference, ":")
		if len(urnParts) != 2 {
			return "", "", "", fmt.Errorf("invalid reference (%s), must have 2 parts", *referenceUri)
		}
		sourceId := urnParts[0]
		resourceParts := strings.Split(urnParts[1], "/")
		if len(resourceParts) != 2 {
			return "", "", "", fmt.Errorf("invalid resource id (%s), must have 2 parts", *referenceUri)
		}

		return sourceId, resourceParts[0], resourceParts[1], nil

	} else {
		return "", "", "", fmt.Errorf("invalid reference (%s), must start with `%s`", *referenceUri, FASTENHEALTH_URN_PREFIX)
	}
}
