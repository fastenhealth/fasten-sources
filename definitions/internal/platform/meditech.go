// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://greenfield-apis.meditech.com/v1/uscore/R4/.well-known/smart-configuration
// https://greenfield-apis.meditech.com/v1/uscore/R4/metadata
// https://fhir.meditech.com/explorer/authorization
func GetSourceMeditech(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}
	sourceDef.AuthorizationEndpoint = "https://greenfield-apis.meditech.com/oauth/authorize"
	sourceDef.TokenEndpoint = "https://greenfield-apis.meditech.com/oauth/token"

	sourceDef.Issuer = "https://greenfield-apis.meditech.com"
	sourceDef.Scopes = []string{"fhirUser", "openid", "patient/*.read", "profile"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.Audience = "https://greenfield-apis.meditech.com/v1/uscore/R4"
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	sourceDef.ApiEndpointBaseUrl = "https://greenfield-apis.meditech.com/v1/uscore/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMeditech]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeMeditech))

	sourceDef.Display = "Meditech (Sandbox)"
	sourceDef.PlatformType = pkg.SourceTypeMeditech
	sourceDef.SourceType = pkg.SourceTypeMeditech
	sourceDef.Category = []string{"Sandbox"}

	return sourceDef, nil
}