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

type SourceClientGreenValleyHospital struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/e3c3f29b-49ce-4329-ade7-4c5593d87595/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/e3c3f29b-49ce-4329-ade7-4c5593d87595/metadata
func GetSourceClientGreenValleyHospital(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientGreenValleyHospital{baseClient}, updatedSourceCred, err
}
