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

// https://fhir-myrecord.cerner.com/r4/70be2934-ad89-4799-aefd-32f31fb0f9ef/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/70be2934-ad89-4799-aefd-32f31fb0f9ef/metadata
func GetSourceRockefellerUniversityHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/70be2934-ad89-4799-aefd-32f31fb0f9ef/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/70be2934-ad89-4799-aefd-32f31fb0f9ef/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/70be2934-ad89-4799-aefd-32f31fb0f9ef"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/70be2934-ad89-4799-aefd-32f31fb0f9ef"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Rockefeller University Hospital"
	sourceDef.SourceType = pkg.SourceTypeRockefellerUniversityHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}