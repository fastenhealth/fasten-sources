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

// https://fhir4.healow.com/fhir/r4/AIDHAA/metadata
func GetSourceElPasoPainCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/AIDHAA"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/AIDHAA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/AIDHAA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeElPasoPainCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "El Paso Pain Center"
	sourceDef.SourceType = pkg.SourceTypeElPasoPainCenter
	sourceDef.Category = []string{"207LP2900X", "208VP0014X", "332B00000X"}
	sourceDef.Aliases = []string{"EL PASO PAIN CENTER", "LAS CRUCES PAIN CENTER", "NEW MEXICO PAIN CENTER OF ALAMOGORDO", "NEW MEXICO PAIN CENTER OF ALBUQUERQUE", "NEW MEXICO PAIN CENTER OF ROSWELL", "NEW MEXICO SPINE & JOINT CENTER", "NORTH TEXAS PAIN CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1285861773", "1710638622", "1881041887"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}