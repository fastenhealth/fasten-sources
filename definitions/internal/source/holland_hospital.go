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

// https://fhir-myrecord.cerner.com/r4/b7350dfb-b26c-45d0-a644-be758c7d66ad/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/b7350dfb-b26c-45d0-a644-be758c7d66ad/metadata
func GetSourceHollandHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/b7350dfb-b26c-45d0-a644-be758c7d66ad/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/b7350dfb-b26c-45d0-a644-be758c7d66ad/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/b7350dfb-b26c-45d0-a644-be758c7d66ad"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/b7350dfb-b26c-45d0-a644-be758c7d66ad"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Holland Hospital"
	sourceDef.SourceType = pkg.SourceTypeHollandHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
