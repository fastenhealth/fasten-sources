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

type SourceClientAdvancedCardioAndVeinCare struct {
	models.SourceClient
}

// https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/10066070/metadata
func GetSourceClientAdvancedCardioAndVeinCare(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientAllscripts(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientAdvancedCardioAndVeinCare{baseClient}, updatedSourceCred, err
}