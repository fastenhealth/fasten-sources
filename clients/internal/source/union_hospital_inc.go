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

type SourceClientUnionHospitalInc struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/6c5673bf-1576-4b76-b281-a10d81e39c32/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/6c5673bf-1576-4b76-b281-a10d81e39c32/metadata
func GetSourceClientUnionHospitalInc(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientUnionHospitalInc{baseClient}, updatedSourceCred, err
}
