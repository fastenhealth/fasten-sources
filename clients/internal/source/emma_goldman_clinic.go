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

type SourceClientEmmaGoldmanClinic struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/52689f1e-9825-4b5d-8853-96de57ae62a9/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/52689f1e-9825-4b5d-8853-96de57ae62a9/metadata
func GetSourceClientEmmaGoldmanClinic(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientEmmaGoldmanClinic{baseClient}, updatedSourceCred, err
}
