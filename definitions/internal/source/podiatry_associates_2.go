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
func GetSourcePodiatryAssociates2(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePodiatryAssociates2]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Podiatry Associates"
	sourceDef.SourceType = pkg.SourceTypePodiatryAssociates2
	sourceDef.Category = []string{"213E00000X", "213EP1101X", "213ES0103X", "332B00000X"}
	sourceDef.Aliases = []string{"PODIATRY ASSOCIATES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083925192", "1194857342", "1225075070", "1235898271", "1912944588", "1942783766"}}
	sourceDef.PatientAccessUrl = "http://www.beaufortpodiatry.com/"
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}