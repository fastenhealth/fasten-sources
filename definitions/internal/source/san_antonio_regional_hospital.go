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

// https://fhir-myrecord.cerner.com/r4/4034ed0f-c20d-41d3-8c75-4fd876e23d84/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/4034ed0f-c20d-41d3-8c75-4fd876e23d84/metadata
func GetSourceSanAntonioRegionalHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/4034ed0f-c20d-41d3-8c75-4fd876e23d84/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/4034ed0f-c20d-41d3-8c75-4fd876e23d84/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/4034ed0f-c20d-41d3-8c75-4fd876e23d84"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/4034ed0f-c20d-41d3-8c75-4fd876e23d84"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "San Antonio Regional Hospital"
	sourceDef.SourceType = pkg.SourceTypeSanAntonioRegionalHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
