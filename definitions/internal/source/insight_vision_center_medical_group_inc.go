// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4/metadata
func GetSourceInsightVisionCenterMedicalGroupInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeInsightVisionCenterMedicalGroupInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "InSight Vision Center Medical Group Inc"
	sourceDef.SourceType = pkg.SourceTypeInsightVisionCenterMedicalGroupInc
	sourceDef.Hidden = true
	sourceDef.BrandLogo = "insight-vision-center-medical-group-inc.jpg"
	sourceDef.PatientAccessUrl = "https://worker.mturk.com/projects/3CTCX9NXCJJWWANBW47QJI84M2SJLA/tasks/3QE4DGPGC48P20SE01DWP5TX6EFG4U?assignment_id=3B4YI393VNTDGA1AB48W16OVJ6GSSV&from_queue=true"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}