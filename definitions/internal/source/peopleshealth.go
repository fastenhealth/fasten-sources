// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://healthx.fhir.flex.optum.com/R4/.well-known/smart-configuration
// https://healthx.fhir.flex.optum.com/R4/metadata
func GetSourcePeopleshealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := GetSourceUnitedhealthcare(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://healthx.authz.flex.optum.com/oauth/authorize"
	sourceDef.TokenEndpoint = "https://healthx.authz.flex.optum.com/oauth/token"
	sourceDef.IntrospectionEndpoint = "https://healthx.authz.flex.optum.com/.well-known/jwks.json"
	sourceDef.UserInfoEndpoint = "https://healthx.authz.flex.optum.com/userinfo"

	sourceDef.Issuer = "https://healthx.fhir.flex.optum.com/R4"
	sourceDef.Audience = "https://healthx.fhir.flex.optum.com/R4"

	sourceDef.ApiEndpointBaseUrl = "https://healthx.fhir.flex.optum.com/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePeopleshealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeUnitedhealthcare))

	sourceDef.Display = "Peoples Health"
	sourceDef.SourceType = pkg.SourceTypePeopleshealth
	sourceDef.Category = []string{"Insurance"}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.peopleshealth.com/"

	return sourceDef, err
}