// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package platform

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir.careevolution.com/Master.Adapter1.WebClient/api/fhir-r4/.well-known/smart-configuration
// https://fhir.careevolution.com/Master.Adapter1.WebClient/api/fhir-r4/metadata
/*
https://fhir.careevolution.com/TestPatientAccounts.html
https://fhir.docs.careevolution.com/overview/tutorials/
*/
func GetSourceCareevolution(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef := models.LighthouseSourceDefinition{}

	sourceDef.Issuer = "https://fhir.careevolution.com"
	sourceDef.Scopes = []string{"fhirUser", "launch/patient", "offline_access", "openid", "patient/*.read"}
	sourceDef.GrantTypesSupported = []string{"authorization_code"}
	sourceDef.ResponseType = []string{"code"}
	sourceDef.ResponseModesSupported = []string{"query"}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCareevolution]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCareevolution))
	sourceDef.Confidential = true

	return sourceDef, nil
}
