package pkg

import (
	"os"
	"strings"
)

func GetCallbackEndpoint(sourceType string) string {
	return os.Getenv("FASTEN_CALLBACK_ENDPOINT") + sourceType
}

type FastenLighthouseEnvType string

const (
	FastenLighthouseEnvSandbox    FastenLighthouseEnvType = "sandbox"
	FastenLighthouseEnvProduction FastenLighthouseEnvType = "prod"
)

func GetFastenLighthouseEnv() FastenLighthouseEnvType {
	env := FastenLighthouseEnvSandbox
	if envString := os.Getenv("FASTEN_ENV"); len(envString) > 0 {
		env = FastenLighthouseEnvType(strings.ToLower(strings.TrimSpace(envString)))
	}
	return env
}
