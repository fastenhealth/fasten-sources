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

// https://fhir-myrecord.cerner.com/r4/6b6f7688-be2d-4aa7-b765-ca3404c86a25/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/6b6f7688-be2d-4aa7-b765-ca3404c86a25/metadata
func GetSourceChildrenSDiabetesAndEndocrinology(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/6b6f7688-be2d-4aa7-b765-ca3404c86a25/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/6b6f7688-be2d-4aa7-b765-ca3404c86a25/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/6b6f7688-be2d-4aa7-b765-ca3404c86a25"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/6b6f7688-be2d-4aa7-b765-ca3404c86a25"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Children's Diabetes and Endocrinology"
	sourceDef.SourceType = pkg.SourceTypeChildrenSDiabetesAndEndocrinology
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
