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

// https://fhir-myrecord.cerner.com/r4/60763645-ac00-4eaa-9f95-71337d5e80c1/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/60763645-ac00-4eaa-9f95-71337d5e80c1/metadata
func GetSourceCullmanRegionalMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/60763645-ac00-4eaa-9f95-71337d5e80c1/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/60763645-ac00-4eaa-9f95-71337d5e80c1/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/60763645-ac00-4eaa-9f95-71337d5e80c1"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/60763645-ac00-4eaa-9f95-71337d5e80c1"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Cullman Regional Medical Center"
	sourceDef.SourceType = pkg.SourceTypeCullmanRegionalMedicalCenter
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
