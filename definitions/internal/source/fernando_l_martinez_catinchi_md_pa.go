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

// https://fhir-myrecord.cerner.com/r4/4e9afe93-70c0-4e43-adb4-49ec1d261955/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/4e9afe93-70c0-4e43-adb4-49ec1d261955/metadata
func GetSourceFernandoLMartinezCatinchiMdPa(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/4e9afe93-70c0-4e43-adb4-49ec1d261955/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/4e9afe93-70c0-4e43-adb4-49ec1d261955/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/4e9afe93-70c0-4e43-adb4-49ec1d261955"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/4e9afe93-70c0-4e43-adb4-49ec1d261955"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Fernando L Martinez-Catinchi MD PA"
	sourceDef.SourceType = pkg.SourceTypeFernandoLMartinezCatinchiMdPa
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
