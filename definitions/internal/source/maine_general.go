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

// https://fhirprod.mainegeneral.org/FHIR/metadata
func GetSourceMaineGeneral(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhirprod.mainegeneral.org/authorization/connect/authorize"
	sourceDef.TokenEndpoint = "https://fhirprod.mainegeneral.org/authorization/connect/token"

	sourceDef.Audience = "https://fhirprod.mainegeneral.org/FHIR"

	sourceDef.ApiEndpointBaseUrl = "https://fhirprod.mainegeneral.org/FHIR"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMaineGeneral]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Maine General"
	sourceDef.SourceType = pkg.SourceTypeMaineGeneral
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}