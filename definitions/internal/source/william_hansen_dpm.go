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

// https://fhir-myrecord.cerner.com/r4/h2eGtPVeNW7cIXpao1aZHtpUU8ww6GX1/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/h2eGtPVeNW7cIXpao1aZHtpUU8ww6GX1/metadata
func GetSourceWilliamHansenDpm(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/h2eGtPVeNW7cIXpao1aZHtpUU8ww6GX1/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/h2eGtPVeNW7cIXpao1aZHtpUU8ww6GX1/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/h2eGtPVeNW7cIXpao1aZHtpUU8ww6GX1"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/h2eGtPVeNW7cIXpao1aZHtpUU8ww6GX1"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "William Hansen, DPM"
	sourceDef.SourceType = pkg.SourceTypeWilliamHansenDpm
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}