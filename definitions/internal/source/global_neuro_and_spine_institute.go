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

// https://api.platform.athenahealth.com/fhir/r4/metadata
func GetSourceGlobalNeuroAndSpineInstitute(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeGlobalNeuroAndSpineInstitute]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Global Neuro and Spine Institute"
	sourceDef.SourceType = pkg.SourceTypeGlobalNeuroAndSpineInstitute
	sourceDef.Category = []string{"208100000X", "2081P0004X", "2081P2900X", "208VP0014X", "332B00000X", "363AS0400X", "363L00000X", "363LP0808X"}
	sourceDef.Aliases = []string{"GLOBAL NEURO AND SPINE INSTITUTE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003500307", "1134813439", "1164116455", "1326732637", "1437843745", "1811200462", "1902590201"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}