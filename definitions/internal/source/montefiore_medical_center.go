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

// https://soapepic.montefiore.org/FhirProxyPrd/api/FHIR/R4/.well-known/smart-configuration
// https://soapepic.montefiore.org/FhirProxyPrd/api/FHIR/R4/metadata
func GetSourceMontefioreMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://soapepic.montefiore.org/FhirProxyPrd/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://soapepic.montefiore.org/FhirProxyPrd/oauth2/token"

	sourceDef.Audience = "https://soapepic.montefiore.org/FhirProxyPrd/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://soapepic.montefiore.org/FhirProxyPrd/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Montefiore Medical Center"
	sourceDef.SourceType = pkg.SourceTypeMontefioreMedicalCenter
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
