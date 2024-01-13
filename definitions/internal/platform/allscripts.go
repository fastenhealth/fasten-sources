// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

/*
https://developer.allscripts.com/content/fhir/content/PRO201_Sandbox/index.html
https://allscripts.vanillacommunities.com/search?query=sandbox&scope=site&source=community
https://open.allscripts.com/fhirendpoints

Allscripts is not actually a confidential client (no client_secret present), however the token endpoint does not support CORS,
so we need to swap the code for the access_token on the server
*/
// https://tw181unityfhir.open.allscripts.com/R4/open-veradigmtwr4/metadata
// https://developer.veradigm.com/Fhir/FHIR_Sandboxes#pehr
func GetSourceAllscripts(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://open.allscripts.com"
	sourceDef.Scopes = []string{"fhirUser", "launch/patient", "offline_access", "openid", "patient/*.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"fragment", "query"}
	sourceDef.CodeChallengeMethodsSupported = []string{"S256"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAllscripts]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))
	sourceDef.Confidential = true

	return sourceDef, nil
}
