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

// https://fhir-myrecord.cerner.com/r4/8d8e2b72-238a-4f00-88a3-8ba1a0a2d561/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/8d8e2b72-238a-4f00-88a3-8ba1a0a2d561/metadata
func GetSourceCookCountyHealthHospitalsSystem(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/8d8e2b72-238a-4f00-88a3-8ba1a0a2d561/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/8d8e2b72-238a-4f00-88a3-8ba1a0a2d561/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/8d8e2b72-238a-4f00-88a3-8ba1a0a2d561"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/8d8e2b72-238a-4f00-88a3-8ba1a0a2d561"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Cook County Health & Hospitals System"
	sourceDef.SourceType = pkg.SourceTypeCookCountyHealthHospitalsSystem
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}