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

// https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46/metadata
func GetSourceTriCountyMentalHealthServices(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/b39fa050-f9bd-471d-b705-2049f60e1d46/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/b39fa050-f9bd-471d-b705-2049f60e1d46/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Tri-County Mental Health Services"
	sourceDef.SourceType = pkg.SourceTypeTriCountyMentalHealthServices
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
