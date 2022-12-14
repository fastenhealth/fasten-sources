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

type SourceClientMountSinaiMedicalCenter struct {
	models.SourceClient
}

// https://epicfhir.msmc.com/proxysite-prd/api/FHIR/R4/.well-known/smart-configuration
// https://epicfhir.msmc.com/proxysite-prd/api/FHIR/R4/metadata
func GetSourceClientMountSinaiMedicalCenter(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientEpic(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientMountSinaiMedicalCenter{baseClient}, updatedSourceCred, err
}
