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

type SourceClientJamesAndDyerMdPa struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/9233Br-NDL0drysKqP8RuVcAw2qdJGLE/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/9233Br-NDL0drysKqP8RuVcAw2qdJGLE/metadata
func GetSourceClientJamesAndDyerMdPa(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientJamesAndDyerMdPa{baseClient}, updatedSourceCred, err
}
