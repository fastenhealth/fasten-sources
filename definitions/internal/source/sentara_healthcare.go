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

// https://epicfhir.sentara.com/ARR-FHIR-PRD/api/FHIR/R4/.well-known/smart-configuration
// https://epicfhir.sentara.com/ARR-FHIR-PRD/api/FHIR/R4/metadata
func GetSourceSentaraHealthcare(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicfhir.sentara.com/ARR-FHIR-PRD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicfhir.sentara.com/ARR-FHIR-PRD/oauth2/token"

	sourceDef.Audience = "https://epicfhir.sentara.com/ARR-FHIR-PRD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicfhir.sentara.com/ARR-FHIR-PRD/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Sentara Healthcare"
	sourceDef.SourceType = pkg.SourceTypeSentaraHealthcare
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
