// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

/*
https://groups.google.com/g/cerner-fhir-developers
http://fhir.cerner.com/millennium/r4/#authorization
*/
// https://fhir-ehr.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/metadata
// https://docs.google.com/document/d/10RnVyF1etl_17pyCyK96tyhUWRbrTyEcqpwzW-Z-Ybs/edit
func GetSourceCerner(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}
	sourceDef.Issuer = "https://authorization.cerner.com"
	sourceDef.Scopes = []string{"fhirUser", "offline_access", "openid", "patient/Account.read", "patient/AllergyIntolerance.read", "patient/Appointment.read", "patient/Binary.read", "patient/CarePlan.read", "patient/CareTeam.read", "patient/ChargeItem.read", "patient/Communication.read", "patient/Condition.read", "patient/Consent.read", "patient/Coverage.read", "patient/Device.read", "patient/DiagnosticReport.read", "patient/DocumentReference.read", "patient/Encounter.read", "patient/FamilyMemberHistory.read", "patient/Goal.read", "patient/Immunization.read", "patient/InsurancePlan.read", "patient/MedicationAdministration.read", "patient/MedicationRequest.read", "patient/NutritionOrder.read", "patient/Observation.read", "patient/Patient.read", "patient/Person.read", "patient/Procedure.read", "patient/Provenance.read", "patient/Questionnaire.read", "patient/QuestionnaireResponse.read", "patient/RelatedPerson.read", "patient/Schedule.read", "patient/ServiceRequest.read", "patient/Slot.read", "user/Location.read", "user/Organization.read", "user/Practitioner.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCerner]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	return sourceDef, nil
}
