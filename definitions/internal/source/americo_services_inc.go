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

// https://fhir-myrecord.cerner.com/r4/8dc8f3b9-781b-44a1-8832-9b919d42e2ad/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/8dc8f3b9-781b-44a1-8832-9b919d42e2ad/metadata
func GetSourceAmericoServicesInc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/8dc8f3b9-781b-44a1-8832-9b919d42e2ad/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/8dc8f3b9-781b-44a1-8832-9b919d42e2ad/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/8dc8f3b9-781b-44a1-8832-9b919d42e2ad"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/8dc8f3b9-781b-44a1-8832-9b919d42e2ad"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Americo Services Inc."
	sourceDef.SourceType = pkg.SourceTypeAmericoServicesInc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}