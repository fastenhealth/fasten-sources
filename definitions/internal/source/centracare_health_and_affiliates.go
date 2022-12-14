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

// https://epicmobile.centracare.com/fhir/api/FHIR/R4/.well-known/smart-configuration
// https://epicmobile.centracare.com/fhir/api/FHIR/R4/metadata
func GetSourceCentracareHealthAndAffiliates(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicmobile.centracare.com/fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicmobile.centracare.com/fhir/oauth2/token"

	sourceDef.Audience = "https://epicmobile.centracare.com/fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicmobile.centracare.com/fhir/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "CentraCare Health and Affiliates"
	sourceDef.SourceType = pkg.SourceTypeCentracareHealthAndAffiliates
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
