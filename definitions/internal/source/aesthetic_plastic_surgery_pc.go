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

// https://fhir-myrecord.cerner.com/r4/26b29267-7a91-4754-a0ab-557e7e4d37f2/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/26b29267-7a91-4754-a0ab-557e7e4d37f2/metadata
func GetSourceAestheticPlasticSurgeryPc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/26b29267-7a91-4754-a0ab-557e7e4d37f2/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/26b29267-7a91-4754-a0ab-557e7e4d37f2/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/26b29267-7a91-4754-a0ab-557e7e4d37f2"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/26b29267-7a91-4754-a0ab-557e7e4d37f2"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Aesthetic Plastic Surgery, PC"
	sourceDef.SourceType = pkg.SourceTypeAestheticPlasticSurgeryPc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}