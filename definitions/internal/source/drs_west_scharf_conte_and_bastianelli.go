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

// https://fhir-myrecord.cerner.com/r4/20f47704-6d52-45dd-af98-57014a72c180/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/20f47704-6d52-45dd-af98-57014a72c180/metadata
func GetSourceDrsWestScharfConteAndBastianelli(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/20f47704-6d52-45dd-af98-57014a72c180/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/20f47704-6d52-45dd-af98-57014a72c180/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/20f47704-6d52-45dd-af98-57014a72c180"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/20f47704-6d52-45dd-af98-57014a72c180"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Drs. West, Scharf, Conte and Bastianelli"
	sourceDef.SourceType = pkg.SourceTypeDrsWestScharfConteAndBastianelli
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}