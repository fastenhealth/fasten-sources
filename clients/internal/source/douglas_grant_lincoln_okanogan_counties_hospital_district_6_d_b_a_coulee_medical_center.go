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

type SourceClientDouglasGrantLincolnOkanoganCountiesHospitalDistrict6DBACouleeMedicalCenter struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/f18be9cb-2758-4a4b-a2a6-5e5131bfb5bc/metadata
func GetSourceClientDouglasGrantLincolnOkanoganCountiesHospitalDistrict6DBACouleeMedicalCenter(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientDouglasGrantLincolnOkanoganCountiesHospitalDistrict6DBACouleeMedicalCenter{baseClient}, updatedSourceCred, err
}