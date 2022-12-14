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

// https://haikuwa.providence.org/fhirproxy/api/FHIR/R4/.well-known/smart-configuration
// https://haikuwa.providence.org/fhirproxy/api/FHIR/R4/metadata
func GetSourceProvidenceHealthAndServicesWashingtonMontana(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://haikuwa.providence.org/fhirproxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://haikuwa.providence.org/fhirproxy/oauth2/token"

	sourceDef.Audience = "https://haikuwa.providence.org/fhirproxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://haikuwa.providence.org/fhirproxy/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Providence Health & Services - Washington/Montana"
	sourceDef.SourceType = pkg.SourceTypeProvidenceHealthAndServicesWashingtonMontana
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
