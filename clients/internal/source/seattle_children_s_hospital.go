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

type SourceClientSeattleChildrenSHospital struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/449052b2-b4e6-4960-bed6-39119385345b/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/449052b2-b4e6-4960-bed6-39119385345b/metadata
func GetSourceClientSeattleChildrenSHospital(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientSeattleChildrenSHospital{baseClient}, updatedSourceCred, err
}
