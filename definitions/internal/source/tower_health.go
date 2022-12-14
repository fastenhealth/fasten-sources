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

// https://epicsoap.readinghospital.org/FHIR-PRD/api/FHIR/R4/.well-known/smart-configuration
// https://epicsoap.readinghospital.org/FHIR-PRD/api/FHIR/R4/metadata
func GetSourceTowerHealth(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://epicsoap.readinghospital.org/FHIR-PRD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicsoap.readinghospital.org/FHIR-PRD/oauth2/token"

	sourceDef.Audience = "https://epicsoap.readinghospital.org/FHIR-PRD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicsoap.readinghospital.org/FHIR-PRD/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Tower Health"
	sourceDef.SourceType = pkg.SourceTypeTowerHealth
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
