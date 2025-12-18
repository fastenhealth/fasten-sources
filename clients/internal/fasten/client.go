package fasten

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/internal/manual"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
)

type FastenClient struct {
	models.SourceClient
}

func GetSourceClientFasten(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, sourceCredsDb models.SourceCredentialRepository, clientOptions ...func(options *models.SourceClientOptions)) (models.SourceClient, error) {
	manualClient, err := manual.GetSourceClientManual(env, ctx, globalLogger, sourceCreds, sourceCredsDb, clientOptions...)
	return &FastenClient{
		manualClient,
	}, err
}
