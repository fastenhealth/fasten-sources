// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhirapi.bcbsal.org/edifecs/fhir/R4/.well-known/smart-configuration
// https://fhirapi.bcbsal.org/edifecs/fhir/R4/metadata
// https://www.bcbsal.org/web/documents/1511503/9929524/FHIR+Documentation.pdf
func GetSourceBcbsal(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://fhirapi.bcbsal.org/edifecs/fhir/R4/"
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBcbsal]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeBcbsal))
	sourceDef.Confidential = true

	return sourceDef, nil
}
