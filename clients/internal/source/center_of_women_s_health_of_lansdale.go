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

type SourceClientCenterOfWomenSHealthOfLansdale struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/025a0f3c-14a7-44fd-90f4-ebf1f0e1cd9e/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/025a0f3c-14a7-44fd-90f4-ebf1f0e1cd9e/metadata
func GetSourceClientCenterOfWomenSHealthOfLansdale(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientCenterOfWomenSHealthOfLansdale{baseClient}, updatedSourceCred, err
}