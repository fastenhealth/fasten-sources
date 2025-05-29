package tefca_direct

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/internal/manual"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
)

type TefcaClient struct {
	models.SourceClient
}

// TEFCA Direct client is a wrapper around the manual upload client.
// This is because it is assumed that TEFCA direct communication will happen out of band, the CCDA will be converted to FHIR and then "stored" using this client
// TODO: this client should just validate the FHIR resources via a linter.
func GetSourceClientTefca(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, clientOptions ...func(options *models.SourceClientOptions)) (models.SourceClient, error) {
	manualClient, err := manual.GetSourceClientManual(env, ctx, globalLogger, sourceCreds, clientOptions...)
	return &TefcaClient{
		manualClient,
	}, err
}
