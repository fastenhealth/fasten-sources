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

// https://fhir-myrecord.cerner.com/r4/9fd949d7-03ce-4a1f-a9ad-7dc3af5a0306/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/9fd949d7-03ce-4a1f-a9ad-7dc3af5a0306/metadata
func GetSourceDrSherwinParikh(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/9fd949d7-03ce-4a1f-a9ad-7dc3af5a0306/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/9fd949d7-03ce-4a1f-a9ad-7dc3af5a0306/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/9fd949d7-03ce-4a1f-a9ad-7dc3af5a0306"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/9fd949d7-03ce-4a1f-a9ad-7dc3af5a0306"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Dr. Sherwin Parikh"
	sourceDef.SourceType = pkg.SourceTypeDrSherwinParikh
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}