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

type SourceClientMedicalClinicOfNorthvillePc struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/3bd112f2-a293-4bc2-86f3-8aca32fd3911/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/3bd112f2-a293-4bc2-86f3-8aca32fd3911/metadata
func GetSourceClientMedicalClinicOfNorthvillePc(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientMedicalClinicOfNorthvillePc{baseClient}, updatedSourceCred, err
}
