// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	"context"
	platform "github.com/fastenhealth/fasten-sources/clients/internal/platform"
	models "github.com/fastenhealth/fasten-sources/clients/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

type SourceClientAllenCountyRegionalHospital struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/6724d0cd-db84-40bf-a313-88dc7a6cd497/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/6724d0cd-db84-40bf-a313-88dc7a6cd497/metadata
func GetSourceClientAllenCountyRegionalHospital(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientAllenCountyRegionalHospital{baseClient}, updatedSourceCred, err
}