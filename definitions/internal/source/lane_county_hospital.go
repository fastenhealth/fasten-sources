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

// https://fhir-myrecord.cerner.com/r4/8b00a59f-c52d-4a54-aef2-92355b29c54b/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/8b00a59f-c52d-4a54-aef2-92355b29c54b/metadata
func GetSourceLaneCountyHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/8b00a59f-c52d-4a54-aef2-92355b29c54b/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/8b00a59f-c52d-4a54-aef2-92355b29c54b/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/8b00a59f-c52d-4a54-aef2-92355b29c54b"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/8b00a59f-c52d-4a54-aef2-92355b29c54b"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Lane County Hospital"
	sourceDef.SourceType = pkg.SourceTypeLaneCountyHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}