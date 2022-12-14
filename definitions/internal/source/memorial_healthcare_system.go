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

// https://mhssp.mhs.net/fhir/api/FHIR/R4/.well-known/smart-configuration
// https://mhssp.mhs.net/fhir/api/FHIR/R4/metadata
func GetSourceMemorialHealthcareSystem(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://mhssp.mhs.net/fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://mhssp.mhs.net/fhir/oauth2/token"

	sourceDef.Audience = "https://mhssp.mhs.net/fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://mhssp.mhs.net/fhir/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Memorial Healthcare System"
	sourceDef.SourceType = pkg.SourceTypeMemorialHealthcareSystem
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
