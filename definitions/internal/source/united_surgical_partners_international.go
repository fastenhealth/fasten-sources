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

// https://fhir-myrecord.cerner.com/r4/ed29dd2e-7f6b-4b48-93ba-6ba78c79cd2d/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/ed29dd2e-7f6b-4b48-93ba-6ba78c79cd2d/metadata
func GetSourceUnitedSurgicalPartnersInternational(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ed29dd2e-7f6b-4b48-93ba-6ba78c79cd2d/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ed29dd2e-7f6b-4b48-93ba-6ba78c79cd2d/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ed29dd2e-7f6b-4b48-93ba-6ba78c79cd2d"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ed29dd2e-7f6b-4b48-93ba-6ba78c79cd2d"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "United Surgical Partners International"
	sourceDef.SourceType = pkg.SourceTypeUnitedSurgicalPartnersInternational
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
