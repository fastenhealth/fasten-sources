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

type SourceClientWarrenJDegaturJrMdFaafp struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/2cb3d6db-a361-49b3-a3c3-68fe3216159f/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/2cb3d6db-a361-49b3-a3c3-68fe3216159f/metadata
func GetSourceClientWarrenJDegaturJrMdFaafp(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientWarrenJDegaturJrMdFaafp{baseClient}, updatedSourceCred, err
}