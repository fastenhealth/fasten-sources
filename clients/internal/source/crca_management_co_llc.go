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

type SourceClientCrcaManagementCoLlc struct {
	models.SourceClient
}

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/fhir/10116683/metadata
func GetSourceClientCrcaManagementCoLlc(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientAllscripts(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientCrcaManagementCoLlc{baseClient}, updatedSourceCred, err
}