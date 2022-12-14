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

// https://fhir-myrecord.cerner.com/r4/14194f57-1cf1-4d77-801e-4659b5b99d4c/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/14194f57-1cf1-4d77-801e-4659b5b99d4c/metadata
func GetSourceAndersonDermatologyAndSkinSurgery(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/14194f57-1cf1-4d77-801e-4659b5b99d4c/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/14194f57-1cf1-4d77-801e-4659b5b99d4c/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/14194f57-1cf1-4d77-801e-4659b5b99d4c"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/14194f57-1cf1-4d77-801e-4659b5b99d4c"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Anderson Dermatology and Skin Surgery"
	sourceDef.SourceType = pkg.SourceTypeAndersonDermatologyAndSkinSurgery
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
