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

// https://fhir4.healow.com/fhir/r4/EIBCAA/metadata
func GetSourceAllianceSpineAndPainCenters(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/EIBCAA"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/EIBCAA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/EIBCAA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAllianceSpineAndPainCenters]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Alliance Spine and Pain Centers"
	sourceDef.SourceType = pkg.SourceTypeAllianceSpineAndPainCenters
	sourceDef.Category = []string{"174400000X", "208VP0014X"}
	sourceDef.Aliases = []string{"ALLIANCE SPINE AND PAIN CENTERS"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083212963", "1114578515", "1174101703", "1235756792", "1447963608", "1467014936", "1518012004", "1720443252", "1861077315"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}