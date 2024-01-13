// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhirsandbox.healthit.gov/secure/r4/.well-known/smart-configuration
/*
CFor demo users use Username: demouser Password: Demouser1!
Mix of clinical and claim data - some synthetic, some de-identified.
User associated with multiple patients, so the system prompts to chose one when using launch/patient
*/
func GetSourceHealthit(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://fhirsandbox.healthit.gov"
	sourceDef.Scopes = []string{"fhirUser", "launch/patient", "openid", "patient/AllergyIntolerance.read", "patient/CarePlan.read", "patient/CareTeam.read", "patient/Condition.read", "patient/Device.read", "patient/DiagnosticReport.read", "patient/DocumentReference.read", "patient/Encounter.read", "patient/Goal.read", "patient/Immunization.read", "patient/Location.read", "patient/Medication.read", "patient/MedicationRequest.read", "patient/Observation.read", "patient/Organization.read", "patient/Patient.read", "patient/Practitioner.read", "patient/PractitionerRole.read", "patient/Procedure.read", "patient/Provenance.read", "patient/RelatedPerson.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHealthit]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeHealthit))

	return sourceDef, nil
}
