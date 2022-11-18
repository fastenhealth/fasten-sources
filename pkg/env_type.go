package pkg

import (
	"os"
	"strings"
)

func GetCallbackEndpoint(sourceType string) string {
	return os.Getenv("FASTEN_CALLBACK_ENDPOINT") + sourceType
}

type FastenEnvType string

const (
	FastenEnvSandbox    FastenEnvType = "sandbox"
	FastenEnvProduction FastenEnvType = "prod"
)

func GetFastenEnv() FastenEnvType {
	env := FastenEnvSandbox
	if envString := os.Getenv("FASTEN_ENV"); len(envString) > 0 {
		env = FastenEnvType(strings.ToLower(strings.TrimSpace(envString)))
	}
	return env
}
