package definitions

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPlatformDefinition(t *testing.T) {
	//setup
	platformTypes := pkg.GetPlatformTypes()

	//test
	for _, platformType := range platformTypes {
		_, err := getPlatformDefinition(platformType, pkg.FastenLighthouseEnvSandbox, map[pkg.PlatformType]string{})
		//assert
		require.NoError(t, err)

	}

}
