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

// https://fhir-myrecord.cerner.com/r4/f17aed19-e2f6-49d4-bba0-64508a2357d2/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/f17aed19-e2f6-49d4-bba0-64508a2357d2/metadata
func GetSourceFairfieldMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f17aed19-e2f6-49d4-bba0-64508a2357d2/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f17aed19-e2f6-49d4-bba0-64508a2357d2/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f17aed19-e2f6-49d4-bba0-64508a2357d2"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f17aed19-e2f6-49d4-bba0-64508a2357d2"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Fairfield Medical Center"
	sourceDef.SourceType = pkg.SourceTypeFairfieldMedicalCenter
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}