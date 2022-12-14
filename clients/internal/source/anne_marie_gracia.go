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

type SourceClientAnneMarieGracia struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/7bgA5p-Jyk4Mxwneej09P02G7MB2Ivog/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/7bgA5p-Jyk4Mxwneej09P02G7MB2Ivog/metadata
func GetSourceClientAnneMarieGracia(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientAnneMarieGracia{baseClient}, updatedSourceCred, err
}
