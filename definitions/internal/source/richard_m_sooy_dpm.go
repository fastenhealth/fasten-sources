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

// https://fhir-myrecord.cerner.com/r4/51439fee-4f26-4c5b-b9eb-6e943125896d/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/51439fee-4f26-4c5b-b9eb-6e943125896d/metadata
func GetSourceRichardMSooyDpm(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/51439fee-4f26-4c5b-b9eb-6e943125896d/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/51439fee-4f26-4c5b-b9eb-6e943125896d/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/51439fee-4f26-4c5b-b9eb-6e943125896d"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/51439fee-4f26-4c5b-b9eb-6e943125896d"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Richard M. Sooy, DPM"
	sourceDef.SourceType = pkg.SourceTypeRichardMSooyDpm
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
