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

type SourceClientGwyneddMercyCollege struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/7f97a135-de8b-4d0d-8cae-944f42aa5be1/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/7f97a135-de8b-4d0d-8cae-944f42aa5be1/metadata
func GetSourceClientGwyneddMercyCollege(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientGwyneddMercyCollege{baseClient}, updatedSourceCred, err
}