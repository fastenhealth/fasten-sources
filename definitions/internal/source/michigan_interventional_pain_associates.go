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

// https://fhir-myrecord.cerner.com/r4/772df939-8efc-47b6-b42b-02586e91e2af/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/772df939-8efc-47b6-b42b-02586e91e2af/metadata
func GetSourceMichiganInterventionalPainAssociates(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/772df939-8efc-47b6-b42b-02586e91e2af/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/772df939-8efc-47b6-b42b-02586e91e2af/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/772df939-8efc-47b6-b42b-02586e91e2af"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/772df939-8efc-47b6-b42b-02586e91e2af"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Michigan Interventional Pain Associates"
	sourceDef.SourceType = pkg.SourceTypeMichiganInterventionalPainAssociates
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}