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

// https://fhir4.healow.com/fhir/r4/HBIHCA/metadata
func GetSourceRothmanInstitute(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/HBIHCA"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/HBIHCA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/HBIHCA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRothmanInstitute]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Rothman Institute"
	sourceDef.SourceType = pkg.SourceTypeRothmanInstitute
	sourceDef.Category = []string{"207QS0010X", "207X00000X", "208100000X", "335E00000X"}
	sourceDef.Aliases = []string{"ROTHMAN INSTITUTE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1386174399", "1861433278", "1962877225", "1992211601"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}