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

// https://fhir-myrecord.cerner.com/r4/04b7d73f-0769-4509-8cd9-c3061dc75979/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/04b7d73f-0769-4509-8cd9-c3061dc75979/metadata
func GetSourceStateOfNewMexico(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/04b7d73f-0769-4509-8cd9-c3061dc75979/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/04b7d73f-0769-4509-8cd9-c3061dc75979/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/04b7d73f-0769-4509-8cd9-c3061dc75979"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/04b7d73f-0769-4509-8cd9-c3061dc75979"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "State of New Mexico"
	sourceDef.SourceType = pkg.SourceTypeStateOfNewMexico
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}