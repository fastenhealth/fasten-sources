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

// https://fhir-myrecord.cerner.com/r4/ed8a62d9-e900-4503-b1c6-0e092e2ef12a/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/ed8a62d9-e900-4503-b1c6-0e092e2ef12a/metadata
func GetSourceSabethaCommunityHospital(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ed8a62d9-e900-4503-b1c6-0e092e2ef12a/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ed8a62d9-e900-4503-b1c6-0e092e2ef12a/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ed8a62d9-e900-4503-b1c6-0e092e2ef12a"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ed8a62d9-e900-4503-b1c6-0e092e2ef12a"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Sabetha Community Hospital"
	sourceDef.SourceType = pkg.SourceTypeSabethaCommunityHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
