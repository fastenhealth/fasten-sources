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

// https://mountainstarhealthfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4/.well-known/smart-configuration
// https://mountainstarhealthfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4/metadata
func GetSourceHcaMountain(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://mountainstarhealthfhirprd.app.medcity.net/fhir-proxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://mountainstarhealthfhirprd.app.medcity.net/fhir-proxy/oauth2/token"

	sourceDef.Audience = "https://mountainstarhealthfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://mountainstarhealthfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "HCA Mountain"
	sourceDef.SourceType = pkg.SourceTypeHcaMountain
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
