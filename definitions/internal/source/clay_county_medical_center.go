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

// https://fhir-myrecord.cerner.com/r4/85e4e207-a7a6-47a1-8e35-8e04fe8019c4/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/85e4e207-a7a6-47a1-8e35-8e04fe8019c4/metadata
func GetSourceClayCountyMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/85e4e207-a7a6-47a1-8e35-8e04fe8019c4/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/85e4e207-a7a6-47a1-8e35-8e04fe8019c4/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/85e4e207-a7a6-47a1-8e35-8e04fe8019c4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/85e4e207-a7a6-47a1-8e35-8e04fe8019c4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Clay County Medical Center"
	sourceDef.SourceType = pkg.SourceTypeClayCountyMedicalCenter
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}