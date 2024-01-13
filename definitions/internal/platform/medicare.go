// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://sandbox.bluebutton.cms.gov/.well-known/openid-configuration-v2
/*
https://groups.google.com/g/Developer-group-for-cms-blue-button-api/c/mVNFJI4dxbs
https://groups.google.com/g/developer-group-for-cms-blue-button-api/c/77ZDwZWHloM/m/jQHZVNznBAAJ?utm_medium=email&utm_source=footer
*/
func GetSourceMedicare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://sandbox.bluebutton.cms.gov"
	sourceDef.Scopes = []string{"patient/Coverage.read", "patient/ExplanationOfBenefit.read", "patient/Patient.read", "profile"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMedicare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeMedicare))
	sourceDef.Confidential = true

	return sourceDef, nil
}
