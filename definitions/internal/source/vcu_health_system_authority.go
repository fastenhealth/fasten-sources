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

// https://fhir-myrecord.cerner.com/r4/072ee25e-9ba2-4626-bd4b-17ff1ccdc763/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/072ee25e-9ba2-4626-bd4b-17ff1ccdc763/metadata
func GetSourceVcuHealthSystemAuthority(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/072ee25e-9ba2-4626-bd4b-17ff1ccdc763/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/072ee25e-9ba2-4626-bd4b-17ff1ccdc763/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/072ee25e-9ba2-4626-bd4b-17ff1ccdc763"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/072ee25e-9ba2-4626-bd4b-17ff1ccdc763"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "VCU Health System Authority"
	sourceDef.SourceType = pkg.SourceTypeVcuHealthSystemAuthority
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}