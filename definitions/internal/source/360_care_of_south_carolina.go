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

// https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4/metadata
func GetSource360CareOfSouthCarolina(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceType360CareOfSouthCarolina]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "360care Of South Carolina"
	sourceDef.SourceType = pkg.SourceType360CareOfSouthCarolina
	sourceDef.Hidden = true
	sourceDef.BrandLogo = "360care-of-south-carolina.jpg"
	sourceDef.PatientAccessUrl = "https://worker.mturk.com/projects/3CTCX9NXCJJWWANBW47QJI84M2SJLA/tasks/3INZSNUD9DNNNMOZF68F0S1GFMDD9T?assignment_id=3GNCZX450WKC3MZO2NIBL1PE4WJAP8&from_queue=true"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}