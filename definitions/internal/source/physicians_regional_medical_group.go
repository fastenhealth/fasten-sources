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
func GetSourcePhysiciansRegionalMedicalGroup(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePhysiciansRegionalMedicalGroup]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Physicians Regional Medical Group"
	sourceDef.SourceType = pkg.SourceTypePhysiciansRegionalMedicalGroup
	sourceDef.Category = []string{"152W00000X", "156FC0800X", "156FX1800X", "207Q00000X", "207R00000X", "207RG0100X", "207V00000X", "207X00000X", "2086S0129X", "208VP0000X", "363A00000X", "363AM0700X", "363AS0400X", "363L00000X", "363LA2200X", "363LF0000X"}
	sourceDef.Aliases = []string{"COLLIER BOULEVARD HMA PHYSICIAN MANAGEMENT LLC", "COLLIER HMA PHYSICIAN MANAGEMENT LLC", "PHYSICIANS REGIONAL MEDICAL GROUP"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1124696000", "1124696919", "1164090874", "1326610833", "1467020248", "1659949345", "1770154791", "1962713248", "1972171759"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}