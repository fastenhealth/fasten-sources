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

// https://fhir-myrecord.cerner.com/r4/e1cd2646-a978-48a1-b1dd-27086f7bcfa4/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/e1cd2646-a978-48a1-b1dd-27086f7bcfa4/metadata
func GetSourceYvelPMoreauMD(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/e1cd2646-a978-48a1-b1dd-27086f7bcfa4/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/e1cd2646-a978-48a1-b1dd-27086f7bcfa4/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/e1cd2646-a978-48a1-b1dd-27086f7bcfa4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/e1cd2646-a978-48a1-b1dd-27086f7bcfa4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Yvel P. Moreau M.D"
	sourceDef.SourceType = pkg.SourceTypeYvelPMoreauMD
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}