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

// https://fhir-myrecord.cerner.com/r4/d2bc3d5a-3c39-4553-86b1-b3aa0a5429a7/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/d2bc3d5a-3c39-4553-86b1-b3aa0a5429a7/metadata
func GetSourceSanJuanHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/d2bc3d5a-3c39-4553-86b1-b3aa0a5429a7/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/d2bc3d5a-3c39-4553-86b1-b3aa0a5429a7/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/d2bc3d5a-3c39-4553-86b1-b3aa0a5429a7"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/d2bc3d5a-3c39-4553-86b1-b3aa0a5429a7"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "San Juan Hospital"
	sourceDef.SourceType = pkg.SourceTypeSanJuanHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}