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

// https://epicproxy.et1127.epichosted.com/FHIRproxy/api/FHIR/R4/.well-known/smart-configuration
// https://epicproxy.et1127.epichosted.com/FHIRproxy/api/FHIR/R4/metadata
func GetSourceUnitedHealthServicesNewYorkNyuhs(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicproxy.et1127.epichosted.com/FHIRproxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicproxy.et1127.epichosted.com/FHIRproxy/oauth2/token"

	sourceDef.Audience = "https://epicproxy.et1127.epichosted.com/FHIRproxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicproxy.et1127.epichosted.com/FHIRproxy/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "United Health Services New York (NYUHS)"
	sourceDef.SourceType = pkg.SourceTypeUnitedHealthServicesNewYorkNyuhs
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
