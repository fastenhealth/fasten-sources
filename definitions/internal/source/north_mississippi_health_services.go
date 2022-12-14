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

// https://eiclbext.nmhs.net/interconnect-fhir-prd/api/FHIR/R4/.well-known/smart-configuration
// https://eiclbext.nmhs.net/interconnect-fhir-prd/api/FHIR/R4/metadata
func GetSourceNorthMississippiHealthServices(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env)
	sourceDef.AuthorizationEndpoint = "https://eiclbext.nmhs.net/interconnect-generaloauth2services-prd/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://eiclbext.nmhs.net/interconnect-generaloauth2services-prd/oauth2/token"

	sourceDef.Audience = "https://eiclbext.nmhs.net/interconnect-fhir-prd/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://eiclbext.nmhs.net/interconnect-fhir-prd/api/FHIR/R4"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "North Mississippi Health Services"
	sourceDef.SourceType = pkg.SourceTypeNorthMississippiHealthServices
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
