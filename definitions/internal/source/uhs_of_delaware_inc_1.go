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

// https://fhir-myrecord.cerner.com/r4/ccc9af17-4ec5-4dee-8735-97dd2f249def/metadata
func GetSourceUhsOfDelawareInc1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ccc9af17-4ec5-4dee-8735-97dd2f249def/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ccc9af17-4ec5-4dee-8735-97dd2f249def/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ccc9af17-4ec5-4dee-8735-97dd2f249def"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ccc9af17-4ec5-4dee-8735-97dd2f249def"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUhsOfDelawareInc1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "UHS of Delaware, Inc"
	sourceDef.SourceType = pkg.SourceTypeUhsOfDelawareInc1
	sourceDef.BrandLogo = "uhs-of-delaware-inc.png"
	sourceDef.PatientAccessUrl = "https://uhs.com/about-uhs/corporate-information/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}