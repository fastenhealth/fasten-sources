package catalog_test

import (
	"fmt"
	"github.com/fastenhealth/fasten-sources/catalog"
	"github.com/fastenhealth/fasten-sources/definitions"
	"github.com/fastenhealth/fasten-sources/pkg"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCatalog_GetBrands(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{Id: "0000f30d-3987-4539-b743-5c263416a6cf"}

	//test
	brands, err := catalog.GetBrands(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, brands, 1)
}

func TestCatalog_GetBrands_WithInvalidId(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{Id: "1"}

	//test
	_, err := catalog.GetBrands(&opts)

	//assert
	require.EqualError(t, err, fmt.Sprintf("brand with id %s not found", opts.Id))
}

func TestCatalog_GetBrands_WithSandboxMode(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	brands, err := catalog.GetBrands(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, brands, 20)

	for _, brand := range brands {
		require.Len(t, brand.PortalsIds, 1)
	}
}

func TestCatalog_GetPortals_WithSandboxMode(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	portals, err := catalog.GetPortals(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, portals, 20)

	for _, portal := range portals {
		require.Len(t, portal.EndpointIds, 1)
	}
}

func TestCatalog_GetEndpoints_WithSandboxMode(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	endpoints, err := catalog.GetEndpoints(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, endpoints, 20)
}

func TestCatalog_GetEndpoints_HaveKnownPlatformType_Production(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType: pkg.FastenLighthouseEnvProduction,
	}
	endpoints, err := catalog.GetEndpoints(&opts)
	require.NoError(t, err)

	endpointPlatformTypes := map[pkg.PlatformType]string{}
	knownPlatformTypes := pkg.GetPlatformTypes()

	//test
	for _, endpoint := range endpoints {
		if _, exists := endpointPlatformTypes[endpoint.GetPlatformType()]; !exists {
			endpointPlatformTypes[endpoint.GetPlatformType()] = endpoint.Id
		}
	}
	foundAllEndpointPlatfromTypes := lo.EveryBy(lo.Keys(endpointPlatformTypes), func(platformType pkg.PlatformType) bool {
		return lo.Contains(knownPlatformTypes, platformType)
	})

	for _, endpointId := range endpointPlatformTypes {
		_, err := definitions.GetSourceDefinition(pkg.FastenLighthouseEnvProduction, map[pkg.PlatformType]string{}, definitions.GetSourceConfigOptions{EndpointId: endpointId})
		require.NoError(t, err)
	}

	//assert
	require.True(t, len(endpointPlatformTypes) >= 1)
	require.True(t, foundAllEndpointPlatfromTypes)
}

func TestCatalog_GetEndpoints_HaveKnownPlatformType_Sandbox(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType: pkg.FastenLighthouseEnvSandbox,
	}
	endpoints, err := catalog.GetEndpoints(&opts)
	require.NoError(t, err)

	endpointPlatformTypes := map[pkg.PlatformType]string{}
	knownPlatformTypes := pkg.GetPlatformTypes()

	//test
	for _, endpoint := range endpoints {
		if _, exists := endpointPlatformTypes[endpoint.GetPlatformType()]; !exists {
			endpointPlatformTypes[endpoint.GetPlatformType()] = endpoint.Id
		}
	}
	foundAllEndpointPlatfromTypes := lo.EveryBy(lo.Keys(endpointPlatformTypes), func(platformType pkg.PlatformType) bool {
		return lo.Contains(knownPlatformTypes, platformType)
	})

	for _, endpointId := range endpointPlatformTypes {
		_, err := definitions.GetSourceDefinition(pkg.FastenLighthouseEnvSandbox, map[pkg.PlatformType]string{}, definitions.GetSourceConfigOptions{EndpointId: endpointId})
		require.NoError(t, err)
	}

	//assert
	require.True(t, len(endpointPlatformTypes) >= 1)
	require.True(t, foundAllEndpointPlatfromTypes)
}
