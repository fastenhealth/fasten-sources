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

// https://fhir-myrecord.cerner.com/r4/dc5473f2-b01f-4963-a3d8-5474b2d33b20/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/dc5473f2-b01f-4963-a3d8-5474b2d33b20/metadata
func GetSourceCommunityMemorialHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/dc5473f2-b01f-4963-a3d8-5474b2d33b20/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/dc5473f2-b01f-4963-a3d8-5474b2d33b20/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/dc5473f2-b01f-4963-a3d8-5474b2d33b20"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/dc5473f2-b01f-4963-a3d8-5474b2d33b20"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Community Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeCommunityMemorialHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}