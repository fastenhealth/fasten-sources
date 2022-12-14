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

// https://fhirprd.ceenta.com/proxy/api/FHIR/R4/.well-known/smart-configuration
// https://fhirprd.ceenta.com/proxy/api/FHIR/R4/metadata
func GetSourceCharlotteEyeEarNoseAndThroatAssociates(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://fhirprd.ceenta.com/proxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://fhirprd.ceenta.com/proxy/oauth2/token"

	sourceDef.Audience = "https://fhirprd.ceenta.com/proxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://fhirprd.ceenta.com/proxy/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Charlotte Eye Ear Nose & Throat Associates"
	sourceDef.SourceType = pkg.SourceTypeCharlotteEyeEarNoseAndThroatAssociates
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
