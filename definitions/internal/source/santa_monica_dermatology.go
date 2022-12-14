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

// https://fhir-myrecord.cerner.com/r4/e1A5O_wfjp6bxazsKQMUgQtMMK_UwF9x/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/e1A5O_wfjp6bxazsKQMUgQtMMK_UwF9x/metadata
func GetSourceSantaMonicaDermatology(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/e1A5O_wfjp6bxazsKQMUgQtMMK_UwF9x/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/e1A5O_wfjp6bxazsKQMUgQtMMK_UwF9x/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/e1A5O_wfjp6bxazsKQMUgQtMMK_UwF9x"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/e1A5O_wfjp6bxazsKQMUgQtMMK_UwF9x"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Santa Monica Dermatology"
	sourceDef.SourceType = pkg.SourceTypeSantaMonicaDermatology
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
