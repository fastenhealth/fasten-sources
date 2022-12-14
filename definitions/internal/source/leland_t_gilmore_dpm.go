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

// https://fhir-myrecord.cerner.com/r4/a2be18e0-7f28-4d4f-a48e-668c32f30fa0/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/a2be18e0-7f28-4d4f-a48e-668c32f30fa0/metadata
func GetSourceLelandTGilmoreDpm(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a2be18e0-7f28-4d4f-a48e-668c32f30fa0/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a2be18e0-7f28-4d4f-a48e-668c32f30fa0/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a2be18e0-7f28-4d4f-a48e-668c32f30fa0"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a2be18e0-7f28-4d4f-a48e-668c32f30fa0"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Leland T. Gilmore, DPM"
	sourceDef.SourceType = pkg.SourceTypeLelandTGilmoreDpm
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
