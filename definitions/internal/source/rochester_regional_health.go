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

// https://epicarr.rochesterregional.org/FHIR/api/FHIR/R4/.well-known/smart-configuration
// https://epicarr.rochesterregional.org/FHIR/api/FHIR/R4/metadata
func GetSourceRochesterRegionalHealth(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicarr.rochesterregional.org/FHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicarr.rochesterregional.org/FHIR/oauth2/token"

	sourceDef.Audience = "https://epicarr.rochesterregional.org/FHIR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicarr.rochesterregional.org/FHIR/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Rochester Regional Health"
	sourceDef.SourceType = pkg.SourceTypeRochesterRegionalHealth
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
