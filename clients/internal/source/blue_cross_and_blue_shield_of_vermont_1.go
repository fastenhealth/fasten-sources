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

type SourceClientBlueCrossAndBlueShieldOfVermont1 struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/d03e2bb3-d17a-409b-bb16-d2cf8d81bb7f/metadata
func GetSourceClientBlueCrossAndBlueShieldOfVermont1(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientBlueCrossAndBlueShieldOfVermont1{baseClient}, updatedSourceCred, err
}