package internal

import (
	"context"
	"fmt"
	base "github.com/fastenhealth/fasten-sources/clients/internal/base"
	models "github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/definitions"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	logrus "github.com/sirupsen/logrus"
)

type dynamicSourceClient struct {
	models.SourceClient
	EndpointDefinition *definitionsModels.LighthouseSourceDefinition
}

func GetDynamicSourceClient(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, clientOptions ...func(options *models.SourceClientOptions)) (models.SourceClient, error) {

	//get the endpoint definition
	endpointDefinition, err := definitions.GetSourceDefinition(definitions.GetSourceConfigOptions{
		EndpointId: sourceCreds.GetEndpointId(),
		Env:        env,
	})
	if err != nil {
		return nil, err
	}
	if endpointDefinition == nil {
		return nil, fmt.Errorf("error retrieving endpoint definition (%s)", sourceCreds.GetEndpointId())
	}

	baseClient, err := base.GetSourceClientFHIR401(env, ctx, globalLogger, sourceCreds, endpointDefinition, clientOptions...)
	if err != nil {
		return nil, err
	}

	// API requires the following headers for every request
	if len(endpointDefinition.ClientHeaders) > 0 {
		for k, v := range endpointDefinition.ClientHeaders {
			baseClient.Headers[k] = v
		}
	}

	return dynamicSourceClient{SourceClient: baseClient, EndpointDefinition: endpointDefinition}, err
}

func GetDynamicSourceClientWithDefinition(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, endpointDefinition *definitionsModels.LighthouseSourceDefinition, clientOptions ...func(options *models.SourceClientOptions)) (models.SourceClient, error) {

	baseClient, err := base.GetSourceClientFHIR401(env, ctx, globalLogger, sourceCreds, endpointDefinition, clientOptions...)
	if err != nil {
		return nil, err
	}

	// API requires the following headers for every request
	if len(endpointDefinition.ClientHeaders) > 0 {
		for k, v := range endpointDefinition.ClientHeaders {
			baseClient.Headers[k] = v
		}
	}

	return dynamicSourceClient{SourceClient: baseClient, EndpointDefinition: endpointDefinition}, err
}

func (c dynamicSourceClient) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {

	if c.EndpointDefinition.MissingOpPatientEverything {
		//Operation-PatientEverything is not supported - https://build.fhir.org/operation-patient-everything.html
		//Manually processing individual resources

		var supportedResources []string
		if len(c.GetResourceTypesAllowList()) > 0 {
			supportedResources = c.GetResourceTypesAllowList()
			supportedResources = append(supportedResources, "Patient") //always ensure the patient is included
		} else {
			//no override provided, attempt to sync all resources
			supportedResources = c.GetResourceTypesUsCore()
			if c.EndpointDefinition.ClientSupportedResources != nil {
				supportedResources = append(supportedResources, c.EndpointDefinition.ClientSupportedResources...)
			}
		}
		return c.SyncAllByResourceName(db, supportedResources)
	} else if len(c.EndpointDefinition.CustomOpPatientEverything) > 0 {
		//Operation-PatientEverything uses non-standard endpoint - https://build.fhir.org/operation-patient-everything.html

		bundle, err := c.GetResourceBundle(c.EndpointDefinition.CustomOpPatientEverything)
		if err != nil {
			return models.UpsertSummary{UpdatedResources: []string{}}, err
		}
		return c.SyncAllByPatientEverythingBundle(db, bundle)
	} else {
		//Standard Operation-PatientEverything

		return c.SourceClient.SyncAll(db)
	}

}
