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

// https://fhir-myrecord.cerner.com/r4/ed13fa1e-5e6c-45da-822a-ab9435740c2e/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/ed13fa1e-5e6c-45da-822a-ab9435740c2e/metadata
func GetSourceSeasonsCenterForCommunityMentalHealth(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ed13fa1e-5e6c-45da-822a-ab9435740c2e/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ed13fa1e-5e6c-45da-822a-ab9435740c2e/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ed13fa1e-5e6c-45da-822a-ab9435740c2e"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ed13fa1e-5e6c-45da-822a-ab9435740c2e"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Seasons Center for Community Mental Health"
	sourceDef.SourceType = pkg.SourceTypeSeasonsCenterForCommunityMentalHealth
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}