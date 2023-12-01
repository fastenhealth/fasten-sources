package fasten

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/internal/manual"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FastenClient struct {
	models.SourceClient
}

func GetSourceClientFasten(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, error) {
	manualClient, err := manual.GetSourceClientManual(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	return &FastenClient{
		manualClient,
	}, err
}
