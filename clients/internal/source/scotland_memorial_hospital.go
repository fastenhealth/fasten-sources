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

type SourceClientScotlandMemorialHospital struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/bf2bc40a-17cb-4353-8de2-68c75c849bc8/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/bf2bc40a-17cb-4353-8de2-68c75c849bc8/metadata
func GetSourceClientScotlandMemorialHospital(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientScotlandMemorialHospital{baseClient}, updatedSourceCred, err
}