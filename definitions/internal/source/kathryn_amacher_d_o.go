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

// https://fhir-myrecord.cerner.com/r4/10521376-f6e6-41c2-a307-0de4e8ac5c2b/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/10521376-f6e6-41c2-a307-0de4e8ac5c2b/metadata
func GetSourceKathrynAmacherDO(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/10521376-f6e6-41c2-a307-0de4e8ac5c2b/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/10521376-f6e6-41c2-a307-0de4e8ac5c2b/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/10521376-f6e6-41c2-a307-0de4e8ac5c2b"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/10521376-f6e6-41c2-a307-0de4e8ac5c2b"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Kathryn Amacher, D.O."
	sourceDef.SourceType = pkg.SourceTypeKathrynAmacherDO
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
