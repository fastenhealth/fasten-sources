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

// https://fhir-myrecord.cerner.com/r4/c27de2ed-48d8-4303-8e9b-4f2faf8e4ea6/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/c27de2ed-48d8-4303-8e9b-4f2faf8e4ea6/metadata
func GetSourcePediatrixMedicalGroupOfFlInc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/c27de2ed-48d8-4303-8e9b-4f2faf8e4ea6/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/c27de2ed-48d8-4303-8e9b-4f2faf8e4ea6/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/c27de2ed-48d8-4303-8e9b-4f2faf8e4ea6"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/c27de2ed-48d8-4303-8e9b-4f2faf8e4ea6"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Pediatrix Medical Group of FL, Inc"
	sourceDef.SourceType = pkg.SourceTypePediatrixMedicalGroupOfFlInc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}