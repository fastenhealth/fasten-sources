// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://providerapi-stage.advancedmd.com/v1/r4/.well-known/smart-configuration
// https://providerapi-stage.advancedmd.com/v1/r4/metadata
/*
https://developers.advancedmd.com/fhir/base-urls
*/
func GetSourceAdvancedmdSandbox(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}
	sourceDef.AuthorizationEndpoint = "https://providerapi-stage.advancedmd.com/v1/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://providerapi-stage.advancedmd.com/v1/oauth2/token"
	sourceDef.IntrospectionEndpoint = "https://providerapi-stage.advancedmd.com/v1/oauth2/introspect"

	sourceDef.Issuer = "https://providerapi-stage.advancedmd.com/v1/r4"
	sourceDef.Scopes = []string{"fhirUser", "offline_access", "openid", "patient/*.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}
	sourceDef.Audience = "https://providerapi-stage.advancedmd.com/v1/r4"
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	sourceDef.ApiEndpointBaseUrl = "https://providerapi-stage.advancedmd.com/v1/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAdvancedmdSandbox]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAdvancedmdSandbox))
	sourceDef.Confidential = true

	sourceDef.Display = "AdvancedMD"
	sourceDef.PlatformType = pkg.SourceTypeAdvancedmdSandbox
	sourceDef.SourceType = pkg.SourceTypeAdvancedmdSandbox
	sourceDef.Category = []string{"Insurance"}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.advancedmd.com"

	return sourceDef, nil
}