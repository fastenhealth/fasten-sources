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

// https://ssproxy.osfhealthcare.org/fhir-proxy/api/FHIR/R4/.well-known/smart-configuration
// https://ssproxy.osfhealthcare.org/fhir-proxy/api/FHIR/R4/metadata
func GetSourceOsfHealthcare(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://ssproxy.osfhealthcare.org/fhir-proxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://ssproxy.osfhealthcare.org/fhir-proxy/oauth2/token"

	sourceDef.Audience = "https://ssproxy.osfhealthcare.org/fhir-proxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://ssproxy.osfhealthcare.org/fhir-proxy/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "OSF HealthCare"
	sourceDef.SourceType = pkg.SourceTypeOsfHealthcare
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
