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

// https://epichaiku.chs-mi.com/FHIRPROXY/api/FHIR/R4/metadata
func GetSourceCovenantHealthcare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epichaiku.chs-mi.com/FHIRPROXY/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epichaiku.chs-mi.com/FHIRPROXY/oauth2/token"

	sourceDef.Audience = "https://epichaiku.chs-mi.com/FHIRPROXY/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epichaiku.chs-mi.com/FHIRPROXY/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCovenantHealthcare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Covenant HealthCare"
	sourceDef.SourceType = pkg.SourceTypeCovenantHealthcare
	sourceDef.Category = []string{"282N00000X", "282NC0060X", "207P00000X", "367500000X"}
	sourceDef.Aliases = []string{"COVENANT HEALTHCARE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1588656946", "1033259569", "1255497764", "1972590412", "1225151897"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}