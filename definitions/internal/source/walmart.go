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

// https://epicinterconnect.walmarthealth.com/Interconnect-OAuth2-PRD/api/FHIR/R4/.well-known/smart-configuration
// https://epicinterconnect.walmarthealth.com/Interconnect-OAuth2-PRD/api/FHIR/R4/metadata
func GetSourceWalmart(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicinterconnect.walmarthealth.com/Interconnect-OAuth2-PRD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicinterconnect.walmarthealth.com/Interconnect-OAuth2-PRD/oauth2/token"

	sourceDef.Audience = "https://epicinterconnect.walmarthealth.com/Interconnect-OAuth2-PRD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicinterconnect.walmarthealth.com/Interconnect-OAuth2-PRD/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Walmart"
	sourceDef.SourceType = pkg.SourceTypeWalmart
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
