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

// https://fhir-myrecord.cerner.com/r4/b3b1f767-47ed-4b48-a4ad-d6b16e6b1d96/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/b3b1f767-47ed-4b48-a4ad-d6b16e6b1d96/metadata
func GetSourceTheStevenACohenMilitaryFamilyClinicAtRedRock(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/b3b1f767-47ed-4b48-a4ad-d6b16e6b1d96/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/b3b1f767-47ed-4b48-a4ad-d6b16e6b1d96/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/b3b1f767-47ed-4b48-a4ad-d6b16e6b1d96"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/b3b1f767-47ed-4b48-a4ad-d6b16e6b1d96"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "The Steven A. Cohen Military Family Clinic At Red Rock"
	sourceDef.SourceType = pkg.SourceTypeTheStevenACohenMilitaryFamilyClinicAtRedRock
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}