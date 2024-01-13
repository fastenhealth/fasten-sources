// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir.collablynk.com/edifecs/fhir/R4/.well-known/smart-configuration
// https://fhir.collablynk.com/edifecs/fhir/R4/metadata
// https://confluence.hl7.org/display/FHIR/Public+Test+Servers
func GetSourceEdifecs(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://fhir.collablynk.com/edifecs/fhir/R4/"
	sourceDef.Scopes = []string{"fhirUser", "openid", "profile"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeEdifecs]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEdifecs))
	sourceDef.Confidential = true

	//sourceDef.SourceType = pkg.SourceTypeEdifecs
	//sourceDef.Category = []string{"Sandbox"}
	//sourceDef.Aliases = []string{}
	//sourceDef.PatientAccessUrl = "https://www.edifecs.com/"

	return sourceDef, nil
}
