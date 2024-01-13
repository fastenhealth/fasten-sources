// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir.humana.com/sandbox/api/metadata
/*
https://developers.humana.com/apis/oauth
*/
func GetSourceHumana(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://fhir.humana.com/sandbox/api"
	sourceDef.Scopes = []string{"internal", "launch/patient", "offline_access", "openid", "patient/AllergyIntolerance.read", "patient/CarePlan.read", "patient/CareTeam.read", "patient/Condition.read", "patient/Coverage.read", "patient/DocumentReference.read", "patient/ExplanationOfBenefit.read", "patient/Goal.read", "patient/Immunization.read", "patient/List.read", "patient/MedicationRequest.read", "patient/Observation.read", "patient/Patient.read", "patient/Procedure.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHumana]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeHumana))
	sourceDef.Confidential = true

	return sourceDef, nil
}
