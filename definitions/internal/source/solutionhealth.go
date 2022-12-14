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

// https://epicproxyprd.solutionhealth.org/FHIR_PROD/api/FHIR/R4/.well-known/smart-configuration
// https://epicproxyprd.solutionhealth.org/FHIR_PROD/api/FHIR/R4/metadata
func GetSourceSolutionhealth(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicproxyprd.solutionhealth.org/FHIR_PROD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicproxyprd.solutionhealth.org/FHIR_PROD/oauth2/token"

	sourceDef.Audience = "https://epicproxyprd.solutionhealth.org/FHIR_PROD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicproxyprd.solutionhealth.org/FHIR_PROD/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "SolutionHealth"
	sourceDef.SourceType = pkg.SourceTypeSolutionhealth
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
