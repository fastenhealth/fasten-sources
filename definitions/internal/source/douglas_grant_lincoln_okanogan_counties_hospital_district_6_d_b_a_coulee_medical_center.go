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

// https://fhir-myrecord.cerner.com/r4/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc/metadata
func GetSourceDouglasGrantLincolnOkanoganCountiesHospitalDistrict6DBACouleeMedicalCenter(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Douglas, Grant, Lincoln & Okanogan Counties Hospital District #6 d/b/a Coulee Medical Center"
	sourceDef.SourceType = pkg.SourceTypeDouglasGrantLincolnOkanoganCountiesHospitalDistrict6DBACouleeMedicalCenter
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}