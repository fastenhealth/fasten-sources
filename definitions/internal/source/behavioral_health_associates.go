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
func GetSourceBehavioralHealthAssociates(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBehavioralHealthAssociates]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Behavioral Health Associates"
	sourceDef.SourceType = pkg.SourceTypeBehavioralHealthAssociates
	sourceDef.Category = []string{"101YM0800X", "103TB0200X", "103TC0700X", "103TC2200X", "103TF0000X", "103TP0814X", "104100000X", "1041C0700X", "2084B0040X", "2084P0800X", "251S00000X", "363LP0808X"}
	sourceDef.Aliases = []string{"BEHAVIORAL HEALTH ASSOCIATES", "ST. DOMINIC PSYCHIATRIC ASSOCIATES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1043521073", "1386042505", "1386708188", "1396072849", "1407385263", "1700813896"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}