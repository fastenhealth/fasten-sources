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

// https://fhir-myrecord.cerner.com/r4/3585fb85-29d2-4cc4-ba39-28265125f97f/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/3585fb85-29d2-4cc4-ba39-28265125f97f/metadata
func GetSourceIronCountyMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/3585fb85-29d2-4cc4-ba39-28265125f97f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/3585fb85-29d2-4cc4-ba39-28265125f97f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/3585fb85-29d2-4cc4-ba39-28265125f97f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/3585fb85-29d2-4cc4-ba39-28265125f97f"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Iron County Medical Center"
	sourceDef.SourceType = pkg.SourceTypeIronCountyMedicalCenter
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}