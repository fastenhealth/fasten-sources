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

// https://fhir4.healow.com/fhir/r4/AAEGBA/metadata
func GetSourceGastroenterologyAssociates1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/AAEGBA"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/AAEGBA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/AAEGBA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeGastroenterologyAssociates1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Gastroenterology Associates"
	sourceDef.SourceType = pkg.SourceTypeGastroenterologyAssociates1
	sourceDef.Category = []string{"174400000X", "207RG0100X", "363A00000X", "363L00000X", "363LA2100X", "363LF0000X", "367500000X"}
	sourceDef.Aliases = []string{"CAPITAL DIGESTIVE CARE, LLC - CHESAPEAKE", "GASTROENTEROLOGY ASSOCIATES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1124064589", "1235330994", "1265640726", "1326073925", "1376751891", "1588688949", "1770578270", "1790963106", "1811556442"}}
	sourceDef.BrandLogo = "gastroenterology-associates.png"
	sourceDef.PatientAccessUrl = "https://www.insitedigestive.com/locations/pasadena-raymond-office/"
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}