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

// https://fhir-myrecord.cerner.com/r4/3fe0dd43-a33d-4fd4-a49d-ab2fa460e893/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/3fe0dd43-a33d-4fd4-a49d-ab2fa460e893/metadata
func GetSourceLuisFSotoLlc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/3fe0dd43-a33d-4fd4-a49d-ab2fa460e893/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/3fe0dd43-a33d-4fd4-a49d-ab2fa460e893/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/3fe0dd43-a33d-4fd4-a49d-ab2fa460e893"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/3fe0dd43-a33d-4fd4-a49d-ab2fa460e893"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Luis F. Soto,LLC."
	sourceDef.SourceType = pkg.SourceTypeLuisFSotoLlc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
