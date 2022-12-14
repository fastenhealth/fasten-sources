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

// https://fhir-myrecord.cerner.com/r4/9c3d5f6c-4d59-4a7f-9170-c96c3d9e561c/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/9c3d5f6c-4d59-4a7f-9170-c96c3d9e561c/metadata
func GetSourceUrologyAssociatesOfTheCentralCoast(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/9c3d5f6c-4d59-4a7f-9170-c96c3d9e561c/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/9c3d5f6c-4d59-4a7f-9170-c96c3d9e561c/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/9c3d5f6c-4d59-4a7f-9170-c96c3d9e561c"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/9c3d5f6c-4d59-4a7f-9170-c96c3d9e561c"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Urology Associates of the Central Coast"
	sourceDef.SourceType = pkg.SourceTypeUrologyAssociatesOfTheCentralCoast
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
