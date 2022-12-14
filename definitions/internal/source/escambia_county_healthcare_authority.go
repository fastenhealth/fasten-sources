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

// https://fhir-myrecord.cerner.com/r4/a7b7f339-eea5-408f-8f14-ad4341b10c45/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/a7b7f339-eea5-408f-8f14-ad4341b10c45/metadata
func GetSourceEscambiaCountyHealthcareAuthority(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a7b7f339-eea5-408f-8f14-ad4341b10c45/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a7b7f339-eea5-408f-8f14-ad4341b10c45/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a7b7f339-eea5-408f-8f14-ad4341b10c45"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a7b7f339-eea5-408f-8f14-ad4341b10c45"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Escambia County Healthcare Authority"
	sourceDef.SourceType = pkg.SourceTypeEscambiaCountyHealthcareAuthority
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
