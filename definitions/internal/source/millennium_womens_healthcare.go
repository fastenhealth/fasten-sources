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

// https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/73952/metadata
func GetSourceMillenniumWomensHealthcare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/authorize?cust=73952"
	sourceDef.TokenEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/token?cust=73952"

	sourceDef.Audience = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/73952"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/73952"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMillenniumWomensHealthcare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Millennium Womens Healthcare"
	sourceDef.SourceType = pkg.SourceTypeMillenniumWomensHealthcare
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}