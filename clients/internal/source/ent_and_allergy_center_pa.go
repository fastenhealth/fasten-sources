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

type SourceClientEntAndAllergyCenterPa struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/bde9959c-b44e-4698-a1d4-e1ee4cc98a86/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/bde9959c-b44e-4698-a1d4-e1ee4cc98a86/metadata
func GetSourceClientEntAndAllergyCenterPa(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientEntAndAllergyCenterPa{baseClient}, updatedSourceCred, err
}
