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

// https://epicproxy.et0502.epichosted.com/FHIRProxy/api/FHIR/R4/.well-known/smart-configuration
// https://epicproxy.et0502.epichosted.com/FHIRProxy/api/FHIR/R4/metadata
func GetSourceUniversityOfCaliforniaSanDiego(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicproxy.et0502.epichosted.com/FhirProxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicproxy.et0502.epichosted.com/FhirProxy/oauth2/token"

	sourceDef.Audience = "https://epicproxy.et0502.epichosted.com/FHIRProxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicproxy.et0502.epichosted.com/FHIRProxy/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "University of California San Diego"
	sourceDef.SourceType = pkg.SourceTypeUniversityOfCaliforniaSanDiego
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
