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

// https://fhir-myrecord.cerner.com/r4/3b364ba7-1e24-4e24-b187-fb3fc86431d8/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/3b364ba7-1e24-4e24-b187-fb3fc86431d8/metadata
func GetSourceBelkysBravoMd(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/3b364ba7-1e24-4e24-b187-fb3fc86431d8/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/3b364ba7-1e24-4e24-b187-fb3fc86431d8/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/3b364ba7-1e24-4e24-b187-fb3fc86431d8"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/3b364ba7-1e24-4e24-b187-fb3fc86431d8"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Belkys Bravo, MD"
	sourceDef.SourceType = pkg.SourceTypeBelkysBravoMd
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
