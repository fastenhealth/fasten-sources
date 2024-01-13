// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://kpx-service-bus.kp.org/service/cdo/siae/healthplankpxv1rc/metadata
// https://developer.kp.org/#/apis
func GetSourceKaiser(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://kpx-service-bus.kp.org/service/cdo/siae/healthplankpxv1rc/FHIR/api"
	sourceDef.Scopes = []string{"launch/patient", "offline_access", "patient/*.read", "sandbox"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeKaiser]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeKaiser))
	sourceDef.Confidential = true

	return sourceDef, nil
}
