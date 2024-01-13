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
func GetSourceAnthem(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := GetSourceCareevolution(env, clientIdLookup)
	if err != nil {
		return sourceDef, err
	}

	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAnthem]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAnthem))
	sourceDef.SecretKeyPrefix = "anthem"

	return sourceDef, nil
}
