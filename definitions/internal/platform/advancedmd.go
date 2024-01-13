// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://providerapi-stage.advancedmd.com/v1/r4/.well-known/smart-configuration
// https://providerapi-stage.advancedmd.com/v1/r4/metadata
/*
https://developers.advancedmd.com/fhir/base-urls
*/
func GetSourceAdvancedmd(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://providerapi-stage.advancedmd.com/v1/r4"
	sourceDef.Scopes = []string{"fhirUser", "offline_access", "openid", "patient/*.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAdvancedmd]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAdvancedmd))
	sourceDef.Confidential = true
	return sourceDef, nil
}
