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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10078847/metadata
func GetSourceWmcPhysicianPracticesLlc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/6574a6f8-ecc8-4c77-9996-a44a00aabae4/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/6574a6f8-ecc8-4c77-9996-a44a00aabae4/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10078847"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10078847"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWmcPhysicianPracticesLlc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "WMC Physician Practices LLC"
	sourceDef.SourceType = pkg.SourceTypeWmcPhysicianPracticesLlc
	sourceDef.Category = []string{"174400000X", "207R00000X", "207RC0000X", "207T00000X", "207VG0400X", "207X00000X", "208000000X", "208600000X", "208800000X"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003299652", "1083096580", "1093162463", "1104280213", "1386026888", "1518303288", "1710360367", "1841672334", "1932563046", "1982086476"}}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}