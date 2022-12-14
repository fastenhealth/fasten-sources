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

// https://fhir-myrecord.cerner.com/r4/6a4397dc-1436-41a9-b8c9-171f0d1bd37f/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/6a4397dc-1436-41a9-b8c9-171f0d1bd37f/metadata
func GetSourceUltimateAccessPrimaryCare(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/6a4397dc-1436-41a9-b8c9-171f0d1bd37f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/6a4397dc-1436-41a9-b8c9-171f0d1bd37f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/6a4397dc-1436-41a9-b8c9-171f0d1bd37f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/6a4397dc-1436-41a9-b8c9-171f0d1bd37f"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Ultimate Access Primary Care"
	sourceDef.SourceType = pkg.SourceTypeUltimateAccessPrimaryCare
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
