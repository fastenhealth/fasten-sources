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

// https://fhir-myrecord.cerner.com/r4/95143858-4077-4cf1-b5d2-c2a999f62bb3/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/95143858-4077-4cf1-b5d2-c2a999f62bb3/metadata
func GetSourceMobridgeRegionalHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/95143858-4077-4cf1-b5d2-c2a999f62bb3/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/95143858-4077-4cf1-b5d2-c2a999f62bb3/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/95143858-4077-4cf1-b5d2-c2a999f62bb3"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/95143858-4077-4cf1-b5d2-c2a999f62bb3"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Mobridge Regional Hospital"
	sourceDef.SourceType = pkg.SourceTypeMobridgeRegionalHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
