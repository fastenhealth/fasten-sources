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

// https://fhir-myrecord.cerner.com/r4/dc96a33f-fb3c-4ad3-8d24-0941136c69c8/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/dc96a33f-fb3c-4ad3-8d24-0941136c69c8/metadata
func GetSourceAdvantageHealthcareForWomen(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/dc96a33f-fb3c-4ad3-8d24-0941136c69c8/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/dc96a33f-fb3c-4ad3-8d24-0941136c69c8/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/dc96a33f-fb3c-4ad3-8d24-0941136c69c8"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/dc96a33f-fb3c-4ad3-8d24-0941136c69c8"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Advantage Healthcare For Women"
	sourceDef.SourceType = pkg.SourceTypeAdvantageHealthcareForWomen
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}