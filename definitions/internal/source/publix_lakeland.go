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

// https://fhir-myrecord.cerner.com/r4/d9a6224a-6280-41a9-8e61-5e638aebc69b/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/d9a6224a-6280-41a9-8e61-5e638aebc69b/metadata
func GetSourcePublixLakeland(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/d9a6224a-6280-41a9-8e61-5e638aebc69b/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/d9a6224a-6280-41a9-8e61-5e638aebc69b/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/d9a6224a-6280-41a9-8e61-5e638aebc69b"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/d9a6224a-6280-41a9-8e61-5e638aebc69b"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Publix - Lakeland"
	sourceDef.SourceType = pkg.SourceTypePublixLakeland
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
