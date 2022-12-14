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

// https://intconfg-p.well-net.org/PRD-OAUTH2/api/FHIR/R4/.well-known/smart-configuration
// https://intconfg-p.well-net.org/PRD-OAUTH2/api/FHIR/R4/metadata
func GetSourceTuftsMedicine(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://intconfg-p.well-net.org/PRD-OAUTH2/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://intconfg-p.well-net.org/PRD-OAUTH2/oauth2/token"

	sourceDef.Audience = "https://intconfg-p.well-net.org/PRD-OAUTH2/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://intconfg-p.well-net.org/PRD-OAUTH2/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Tufts Medicine"
	sourceDef.SourceType = pkg.SourceTypeTuftsMedicine
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
