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

type SourceClientCityWideMedical struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/482289d0-2f39-497e-b3a8-2edb8b0b53e9/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/482289d0-2f39-497e-b3a8-2edb8b0b53e9/metadata
func GetSourceClientCityWideMedical(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientCityWideMedical{baseClient}, updatedSourceCred, err
}
