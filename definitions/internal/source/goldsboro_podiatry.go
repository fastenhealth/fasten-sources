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

// https://fhir-myrecord.cerner.com/r4/1ee7fa8c-2604-4e16-b1f2-79febf0fab94/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/1ee7fa8c-2604-4e16-b1f2-79febf0fab94/metadata
func GetSourceGoldsboroPodiatry(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/1ee7fa8c-2604-4e16-b1f2-79febf0fab94/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/1ee7fa8c-2604-4e16-b1f2-79febf0fab94/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/1ee7fa8c-2604-4e16-b1f2-79febf0fab94"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/1ee7fa8c-2604-4e16-b1f2-79febf0fab94"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Goldsboro Podiatry"
	sourceDef.SourceType = pkg.SourceTypeGoldsboroPodiatry
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}