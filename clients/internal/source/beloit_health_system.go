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

type SourceClientBeloitHealthSystem struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/f4477d7c-dbb4-4046-92c9-1cc56c248ecb/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/f4477d7c-dbb4-4046-92c9-1cc56c248ecb/metadata
func GetSourceClientBeloitHealthSystem(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientBeloitHealthSystem{baseClient}, updatedSourceCred, err
}