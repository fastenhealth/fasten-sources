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

// https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/10052317/metadata
func GetSourcePomereneHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/authorize?cust=10052317"
	sourceDef.TokenEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/token?cust=10052317"

	sourceDef.Audience = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/10052317"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/10052317"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePomereneHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Pomerene Hospital"
	sourceDef.SourceType = pkg.SourceTypePomereneHospital
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}