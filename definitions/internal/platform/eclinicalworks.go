// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

/*
https://connect4.healow.com/apps/jsp/dev/signIn.jsp
https://fhir.eclinicalworks.com/ecwopendev/
*/
// https://fhir4.eclinicalworks.com/fhir/r4/JAFJCD/.well-known/smart-configuration
// https://fhir4.eclinicalworks.com/fhir/r4/JAFJCD/metadata
// https://fhir.eclinicalworks.com/ecwopendev/documentation
func GetSourceEclinicalworks(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://fhir4.eclinicalworks.com/fhir/r4/JAFJCD"
	sourceDef.Scopes = []string{"fhirUser", "offline_access", "openid", "patient/AllergyIntolerance.read", "patient/AllergyIntolerance.search", "patient/Binary.read", "patient/CarePlan.read", "patient/CarePlan.search", "patient/CareTeam.read", "patient/CareTeam.search", "patient/Condition.read", "patient/Condition.search", "patient/Device.read", "patient/Device.search", "patient/DiagnosticReport.read", "patient/DiagnosticReport.search", "patient/DocumentReference.read", "patient/DocumentReference.search", "patient/Encounter.read", "patient/Encounter.search", "patient/Goal.read", "patient/Goal.search", "patient/Immunization.read", "patient/Immunization.search", "patient/Location.read", "patient/Medication.read", "patient/MedicationAdministration.read", "patient/MedicationAdministration.search", "patient/MedicationRequest.read", "patient/MedicationRequest.search", "patient/Observation.read", "patient/Observation.search", "patient/Organization.read", "patient/Organization.search", "patient/Patient.read", "patient/Patient.read", "patient/Patient.search", "patient/Practitioner.read", "patient/Practitioner.search", "patient/PractitionerRole.read", "patient/PractitionerRole.search", "patient/Procedure.read", "patient/Procedure.search", "patient/Provenance.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeEclinicalworks]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))
	sourceDef.Confidential = true

	return sourceDef, nil
}
