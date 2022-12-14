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

// https://fhir-myrecord.cerner.com/r4/18c2ddf7-cb55-4622-a2b3-eb9aad8f5a43/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/18c2ddf7-cb55-4622-a2b3-eb9aad8f5a43/metadata
func GetSourceTheChildrenSHospitalCorporation(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/18c2ddf7-cb55-4622-a2b3-eb9aad8f5a43/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/18c2ddf7-cb55-4622-a2b3-eb9aad8f5a43/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/18c2ddf7-cb55-4622-a2b3-eb9aad8f5a43"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/18c2ddf7-cb55-4622-a2b3-eb9aad8f5a43"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "The Children's Hospital Corporation"
	sourceDef.SourceType = pkg.SourceTypeTheChildrenSHospitalCorporation
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
