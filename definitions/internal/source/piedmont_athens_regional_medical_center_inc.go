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

// https://fhir-myrecord.cerner.com/r4/6fa884ca-f316-49ef-a286-ec1da30b45fd/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/6fa884ca-f316-49ef-a286-ec1da30b45fd/metadata
func GetSourcePiedmontAthensRegionalMedicalCenterInc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/6fa884ca-f316-49ef-a286-ec1da30b45fd/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/6fa884ca-f316-49ef-a286-ec1da30b45fd/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/6fa884ca-f316-49ef-a286-ec1da30b45fd"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/6fa884ca-f316-49ef-a286-ec1da30b45fd"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Piedmont Athens Regional Medical Center Inc."
	sourceDef.SourceType = pkg.SourceTypePiedmontAthensRegionalMedicalCenterInc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
