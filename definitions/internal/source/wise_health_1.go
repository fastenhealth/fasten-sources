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

// https://fhir.we0.hos.allscriptscloud.com/Open/metadata
func GetSourceWiseHealth1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://global.open.allscripts.com/fhirroute/fmhpatientauth/ff5ad1df-7c5d-45ef-a350-32b61970bdda/connect/authorize"
	sourceDef.TokenEndpoint = "https://global.open.allscripts.com/fhirroute/fmhpatientauth/ff5ad1df-7c5d-45ef-a350-32b61970bdda/connect/token"

	sourceDef.Audience = "https://fhir.we0.hos.allscriptscloud.com/Open"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.we0.hos.allscriptscloud.com/Open"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWiseHealth1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Wise Health"
	sourceDef.SourceType = pkg.SourceTypeWiseHealth1
	sourceDef.Category = []string{"174H00000X", "251E00000X", "251K00000X", "251S00000X", "253Z00000X", "302R00000X", "374U00000X", "405300000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1194236752"}}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}