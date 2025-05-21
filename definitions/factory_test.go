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
		_, err := getPlatformDefinition(platformType)
		//assert
		require.NoError(t, err)

	}

}

func TestGetSourceDefinition_WithKnownMergedEndpointId(t *testing.T) {
	//setup
	endpointDefinition, err := GetSourceDefinition(
		WithEndpointId("fc94bfc7-684d-4e4d-aa6e-ceec01c21c81"),
		WithEnv(pkg.FastenLighthouseEnvSandbox),
		WithClientIdLookup(map[pkg.PlatformType]string{
			pkg.PlatformTypeEpicLegacy: "placeholder-clientid",
		}),
	)
	require.NoError(t, err)
	require.Equal(t, "8e2f5de7-46ac-4067-96ba-5e3f60ad52a4", endpointDefinition.Id)
}
