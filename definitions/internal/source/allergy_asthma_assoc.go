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

// https://fhir-myrecord.cerner.com/r4/c3d6a188-86bd-4e90-a26d-5bfd89cb9cd7/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/c3d6a188-86bd-4e90-a26d-5bfd89cb9cd7/metadata
func GetSourceAllergyAsthmaAssoc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/c3d6a188-86bd-4e90-a26d-5bfd89cb9cd7/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/c3d6a188-86bd-4e90-a26d-5bfd89cb9cd7/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/c3d6a188-86bd-4e90-a26d-5bfd89cb9cd7"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/c3d6a188-86bd-4e90-a26d-5bfd89cb9cd7"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Allergy & Asthma Assoc"
	sourceDef.SourceType = pkg.SourceTypeAllergyAsthmaAssoc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}