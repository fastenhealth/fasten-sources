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

// https://fhir-myrecord.cerner.com/r4/1d0193b5-c6df-4894-ba58-5a49f2d64f8f/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/1d0193b5-c6df-4894-ba58-5a49f2d64f8f/metadata
func GetSourceTexasTechUniversityHealthSciencesCenterAtElPaso(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/1d0193b5-c6df-4894-ba58-5a49f2d64f8f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/1d0193b5-c6df-4894-ba58-5a49f2d64f8f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/1d0193b5-c6df-4894-ba58-5a49f2d64f8f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/1d0193b5-c6df-4894-ba58-5a49f2d64f8f"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Texas Tech University Health Sciences Center at El Paso"
	sourceDef.SourceType = pkg.SourceTypeTexasTechUniversityHealthSciencesCenterAtElPaso
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}