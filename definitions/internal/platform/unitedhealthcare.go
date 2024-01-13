// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://sandbox.fhir.flex.optum.com/R4/.well-known/smart-configuration
// https://sandbox.fhir.flex.optum.com/R4/metadata
// https://www.uhc.com/legal/interoperability-apis
func GetSourceUnitedhealthcare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://sandbox.fhir.flex.optum.com/R4"
	sourceDef.Scopes = []string{"fhirUser", "openid", "patient/Condition.read", "patient/Coverage.read", "patient/Encounter.read", "patient/ExplanationOfBenefit.read", "patient/Immunization.read", "patient/MedicationDispense.read", "patient/MedicationRequest.read", "patient/Observation.read", "patient/Patient.read", "patient/Procedure.read", "profile"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUnitedhealthcare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeUnitedhealthcare))

	return sourceDef, nil
}
