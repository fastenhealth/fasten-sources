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

// https://fhir-myrecord.cerner.com/r4/1ad2a81d-d102-4f59-a87a-db4fa01056bd/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/1ad2a81d-d102-4f59-a87a-db4fa01056bd/metadata
func GetSourceRenalHypertensionClinicPc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/1ad2a81d-d102-4f59-a87a-db4fa01056bd/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/1ad2a81d-d102-4f59-a87a-db4fa01056bd/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/1ad2a81d-d102-4f59-a87a-db4fa01056bd"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/1ad2a81d-d102-4f59-a87a-db4fa01056bd"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Renal-Hypertension Clinic, PC"
	sourceDef.SourceType = pkg.SourceTypeRenalHypertensionClinicPc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}