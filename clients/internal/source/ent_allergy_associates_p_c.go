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

type SourceClientEntAllergyAssociatesPC struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/15dd2bda-9d00-4c3d-967c-f374acf8d6ab/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/15dd2bda-9d00-4c3d-967c-f374acf8d6ab/metadata
func GetSourceClientEntAllergyAssociatesPC(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientEntAllergyAssociatesPC{baseClient}, updatedSourceCred, err
}
