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

// https://fhir-myrecord.cerner.com/r4/HNIOVRd9TyZjn7Uq3RvU0cvF3vA4G9IF/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/HNIOVRd9TyZjn7Uq3RvU0cvF3vA4G9IF/metadata
func GetSourceFrancisVAdamsMd(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/HNIOVRd9TyZjn7Uq3RvU0cvF3vA4G9IF/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/HNIOVRd9TyZjn7Uq3RvU0cvF3vA4G9IF/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/HNIOVRd9TyZjn7Uq3RvU0cvF3vA4G9IF"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/HNIOVRd9TyZjn7Uq3RvU0cvF3vA4G9IF"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Francis V Adams, MD"
	sourceDef.SourceType = pkg.SourceTypeFrancisVAdamsMd
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}