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

// https://fhir-myrecord.cerner.com/r4/53393c7a-0b87-4dcb-bbca-68e82f45f8dc/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/53393c7a-0b87-4dcb-bbca-68e82f45f8dc/metadata
func GetSourceRenalMedicalAssociationsLlc(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/53393c7a-0b87-4dcb-bbca-68e82f45f8dc/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/53393c7a-0b87-4dcb-bbca-68e82f45f8dc/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/53393c7a-0b87-4dcb-bbca-68e82f45f8dc"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/53393c7a-0b87-4dcb-bbca-68e82f45f8dc"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Renal Medical Associations, LLC"
	sourceDef.SourceType = pkg.SourceTypeRenalMedicalAssociationsLlc
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
