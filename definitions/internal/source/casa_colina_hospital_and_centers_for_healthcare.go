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

// https://fhir-myrecord.cerner.com/r4/a477462c-6289-4395-a870-ea27830a6de9/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/a477462c-6289-4395-a870-ea27830a6de9/metadata
func GetSourceCasaColinaHospitalAndCentersForHealthcare(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a477462c-6289-4395-a870-ea27830a6de9/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a477462c-6289-4395-a870-ea27830a6de9/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a477462c-6289-4395-a870-ea27830a6de9"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a477462c-6289-4395-a870-ea27830a6de9"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Casa Colina Hospital and Centers for Healthcare"
	sourceDef.SourceType = pkg.SourceTypeCasaColinaHospitalAndCentersForHealthcare
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}