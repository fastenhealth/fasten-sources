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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10085725/metadata
func GetSourceFranklinCountyMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/0dbea437-4b07-496e-a588-a72201195722/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/0dbea437-4b07-496e-a588-a72201195722/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10085725"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10085725"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeFranklinCountyMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Franklin County Medical Center"
	sourceDef.SourceType = pkg.SourceTypeFranklinCountyMedicalCenter
	sourceDef.Category = []string{"207Q00000X", "207V00000X", "251E00000X", "251G00000X", "261QR1300X", "282NC0060X", "314000000X", "3336H0001X", "3336I0012X", "3336L0003X"}
	sourceDef.Aliases = []string{"WILLOW VALLEY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1093056475", "1144663543", "1376579003", "1548355811"}}
	sourceDef.PatientAccessUrl = "https://www.fcmc.org/"
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}