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

// https://fhir-myrecord.cerner.com/r4/f4c4f66a-4f9a-401a-9995-69047ac19451/metadata
func GetSourceDelawareFootAndAnkleGroup(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f4c4f66a-4f9a-401a-9995-69047ac19451/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f4c4f66a-4f9a-401a-9995-69047ac19451/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f4c4f66a-4f9a-401a-9995-69047ac19451"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f4c4f66a-4f9a-401a-9995-69047ac19451"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDelawareFootAndAnkleGroup]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Delaware Foot & Ankle Group"
	sourceDef.SourceType = pkg.SourceTypeDelawareFootAndAnkleGroup
	sourceDef.PatientAccessUrl = "https://www.defootandanklegroup.net/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}