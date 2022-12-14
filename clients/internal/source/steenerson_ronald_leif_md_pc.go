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

type SourceClientSteenersonRonaldLeifMdPc struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/1008d04d-98a2-4346-b611-469bbbcd59ef/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/1008d04d-98a2-4346-b611-469bbbcd59ef/metadata
func GetSourceClientSteenersonRonaldLeifMdPc(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientSteenersonRonaldLeifMdPc{baseClient}, updatedSourceCred, err
}
