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

// https://fhir-myrecord.cerner.com/r4/77c2b089-5ae8-473f-98a4-de9173e168aa/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/77c2b089-5ae8-473f-98a4-de9173e168aa/metadata
func GetSourceAssociationForMentalHealthAndWellness(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/77c2b089-5ae8-473f-98a4-de9173e168aa/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/77c2b089-5ae8-473f-98a4-de9173e168aa/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/77c2b089-5ae8-473f-98a4-de9173e168aa"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/77c2b089-5ae8-473f-98a4-de9173e168aa"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Association for Mental Health and Wellness"
	sourceDef.SourceType = pkg.SourceTypeAssociationForMentalHealthAndWellness
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
