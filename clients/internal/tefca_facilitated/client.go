package tefca_facilitated

import (
	"context"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/internal"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/definitions"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
)

type TefcaFacilitatedFHIRClient struct {
	models.SourceClient
}

// GetSourceClientTefcaFacilitated TEFCA Facilitated client is a wrapper around the FHIR R4 client
// As of now, only Epic QHIN requires TEFCA Facilitated FHIR (all other QHIN communication is direct).
// For Epic, the TEFCA/QHIN RLS will respond with a FHIR endpoint + OID for each institution the patient has been to.
// Fasten must then use a standard Smart-on-FHIR connection to the Epic FHIR server to retrieve medical records for the patient.
//
// SEE: https://rce.sequoiaproject.org/wp-content/uploads/2022/10/TEFCA-Facilitated-FHIR-Implementation-Guide-Draft-for-508.pdf
//
// The problem is that we're re-using existing Endpoint Definitions (which already provide a "platformType").
// So, this client will override the platform data (and client id), and return a modified Client
//
// TODO: this client should just validate the FHIR resources via a linter.
func GetSourceClientTefcaFacilitated(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, sourceCredsDb models.SourceCredentialRepository, clientOptions ...func(options *models.SourceClientOptions)) (models.SourceClient, error) {
	//get the endpoint definition
	endpointDefinition, err := definitions.GetSourceDefinition(
		definitions.WithEndpointId(sourceCreds.GetEndpointId()),
		definitions.WithEnv(env),

		//this override will merge the endpoint definition with the `tefca-*.yaml` platform definition
		definitions.WithPlatformTypeOverride(sourceCreds.GetPlatformType()),
	)
	if err != nil {
		return nil, err
	}
	if endpointDefinition == nil {
		return nil, fmt.Errorf("error retrieving endpoint definition (%s)", sourceCreds.GetEndpointId())
	}

	dynamicSourceClient, err := internal.GetDynamicSourceClientWithDefinition(env, ctx, globalLogger, sourceCreds, sourceCredsDb, endpointDefinition, clientOptions...)
	return &TefcaFacilitatedFHIRClient{
		dynamicSourceClient,
	}, err
}
