// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package platform

import (
	"context"
	base "github.com/fastenhealth/fasten-sources/clients/internal/base"
	models "github.com/fastenhealth/fasten-sources/clients/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

type sourceClientAetna struct {
	models.SourceClient
}

// https://vteapif1.aetna.com/fhirdemo/.well-known/smart-configuration
// https://vteapif1.aetna.com/fhirdemo/v1/patientaccess/metadata
// https://developerportal.aetna.com/Aetna_TestMember_Data_V6.xls
func GetSourceClientAetna(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, error) {
	baseClient, err := base.GetSourceClientFHIR401(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	if err != nil {
		return nil, err
	}

	return sourceClientAetna{baseClient}, err
}

// Operation-PatientEverything uses non-standard endpoint - https://build.fhir.org/operation-patient-everything.html
func (c sourceClientAetna) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	bundle, err := c.GetResourceBundle("Patient")
	if err != nil {
		return models.UpsertSummary{UpdatedResources: []string{}}, err
	}
	return c.SyncAllByPatientEverythingBundle(db, bundle)
}