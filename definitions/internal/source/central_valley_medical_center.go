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

// https://fhir-myrecord.cerner.com/r4/26081706-bf97-4a1b-96e4-2d82c3343f85/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/26081706-bf97-4a1b-96e4-2d82c3343f85/metadata
func GetSourceCentralValleyMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/26081706-bf97-4a1b-96e4-2d82c3343f85/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/26081706-bf97-4a1b-96e4-2d82c3343f85/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/26081706-bf97-4a1b-96e4-2d82c3343f85"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/26081706-bf97-4a1b-96e4-2d82c3343f85"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Central Valley Medical Center"
	sourceDef.SourceType = pkg.SourceTypeCentralValleyMedicalCenter
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
