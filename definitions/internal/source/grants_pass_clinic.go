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

// https://FHIR.grantspassclinic.com/FHIR/metadata
func GetSourceGrantsPassClinic(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://FHIR.grantspassclinic.com/authorization/connect/authorize"
	sourceDef.TokenEndpoint = "https://FHIR.grantspassclinic.com/authorization/connect/token"

	sourceDef.Audience = "https://FHIR.grantspassclinic.com/FHIR"

	sourceDef.ApiEndpointBaseUrl = "https://FHIR.grantspassclinic.com/FHIR"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeGrantsPassClinic]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Grants Pass Clinic"
	sourceDef.SourceType = pkg.SourceTypeGrantsPassClinic
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}