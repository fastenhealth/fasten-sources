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

// https://fhir-myrecord.cerner.com/r4/93e5ba5f-01b6-4ff8-aac2-2bc00f824e93/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/93e5ba5f-01b6-4ff8-aac2-2bc00f824e93/metadata
func GetSourceCommunityFirstHealthPlans(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/93e5ba5f-01b6-4ff8-aac2-2bc00f824e93/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/93e5ba5f-01b6-4ff8-aac2-2bc00f824e93/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/93e5ba5f-01b6-4ff8-aac2-2bc00f824e93"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/93e5ba5f-01b6-4ff8-aac2-2bc00f824e93"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Community First Health Plans"
	sourceDef.SourceType = pkg.SourceTypeCommunityFirstHealthPlans
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}