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

// https://fhir4.eclinicalworks.com/fhir/r4/DAHFAD/metadata
func GetSourceMenaRegionalHealthSystem(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Audience = "https://fhir4.eclinicalworks.com/fhir/r4/DAHFAD"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.eclinicalworks.com/fhir/r4/DAHFAD"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMenaRegionalHealthSystem]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Mena Regional Health System"
	sourceDef.SourceType = pkg.SourceTypeMenaRegionalHealthSystem
	sourceDef.Category = []string{"207Q00000X", "208000000X", "208D00000X", "261QR1300X"}
	sourceDef.Aliases = []string{"MENA REGIONAL HEALTH SYSTEM"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003473463", "1598322810"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}