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

// https://fhir-myrecord.cerner.com/r4/f8a4e006-d37e-48eb-96c9-9fb24047ef7d/metadata
func GetSourceRichardDavenportMdAndAssociatesSc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f8a4e006-d37e-48eb-96c9-9fb24047ef7d/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f8a4e006-d37e-48eb-96c9-9fb24047ef7d/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f8a4e006-d37e-48eb-96c9-9fb24047ef7d"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f8a4e006-d37e-48eb-96c9-9fb24047ef7d"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRichardDavenportMdAndAssociatesSc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Richard Davenport, MD & Associates, SC"
	sourceDef.SourceType = pkg.SourceTypeRichardDavenportMdAndAssociatesSc
	sourceDef.PatientAccessUrl = "https://www.md.com/doctor/richard-davenport-md"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}