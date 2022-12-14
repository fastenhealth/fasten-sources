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

// https://fhir-myrecord.cerner.com/r4/80f02a57-2bb0-4823-b4e1-a80c1fe53bb1/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/80f02a57-2bb0-4823-b4e1-a80c1fe53bb1/metadata
func GetSourceCompleteMedicalHealthcare(env pkg.FastenLighthouseEnvType) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/80f02a57-2bb0-4823-b4e1-a80c1fe53bb1/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/80f02a57-2bb0-4823-b4e1-a80c1fe53bb1/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/80f02a57-2bb0-4823-b4e1-a80c1fe53bb1"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/80f02a57-2bb0-4823-b4e1-a80c1fe53bb1"
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Complete Medical Healthcare"
	sourceDef.SourceType = pkg.SourceTypeCompleteMedicalHealthcare
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
