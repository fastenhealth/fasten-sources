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

// https://mobileprod.arcmd.com/FHIR/api/FHIR/R4/.well-known/smart-configuration
// https://mobileprod.arcmd.com/FHIR/api/FHIR/R4/metadata
func GetSourceAustinRegionalClinic(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://mobileprod.arcmd.com/FHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://mobileprod.arcmd.com/FHIR/oauth2/token"

	sourceDef.Audience = "https://mobileprod.arcmd.com/FHIR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://mobileprod.arcmd.com/FHIR/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Austin Regional Clinic"
	sourceDef.SourceType = pkg.SourceTypeAustinRegionalClinic
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
