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

// https://fhir-myrecord.cerner.com/r4/812e0a2c-a2ee-43b3-baef-01bf949ae18f/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/812e0a2c-a2ee-43b3-baef-01bf949ae18f/metadata
func GetSourceLataMundkurMd(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/812e0a2c-a2ee-43b3-baef-01bf949ae18f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/812e0a2c-a2ee-43b3-baef-01bf949ae18f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/812e0a2c-a2ee-43b3-baef-01bf949ae18f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/812e0a2c-a2ee-43b3-baef-01bf949ae18f"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Lata Mundkur, MD"
	sourceDef.SourceType = pkg.SourceTypeLataMundkurMd
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
