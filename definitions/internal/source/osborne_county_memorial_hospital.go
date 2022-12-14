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

// https://fhir-myrecord.cerner.com/r4/bda9c6b4-e0b0-49ed-9423-c0c3e2ff982a/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/bda9c6b4-e0b0-49ed-9423-c0c3e2ff982a/metadata
func GetSourceOsborneCountyMemorialHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/bda9c6b4-e0b0-49ed-9423-c0c3e2ff982a/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/bda9c6b4-e0b0-49ed-9423-c0c3e2ff982a/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/bda9c6b4-e0b0-49ed-9423-c0c3e2ff982a"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/bda9c6b4-e0b0-49ed-9423-c0c3e2ff982a"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Osborne County Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeOsborneCountyMemorialHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
