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
	require.LessOrEqual(t, len(brands), 100)

	for _, brand := range brands {
		if brand.Id == "db814755-2b62-4549-ba65-5138c0b80536" {
			require.Len(t, brand.PortalsIds, 2)
		}
	}
}

func TestCatalog_GetBrands_WithCache(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}
	endpoints, err := catalog.GetEndpoints(&opts)
	require.NoError(t, err)
	portals, err := catalog.GetPortals(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox, CachedEndpointsLookup: &endpoints})
	//test
	brands, err := catalog.GetBrands(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox, CachedPortalsLookup: &portals})

	//assert
	require.NoError(t, err)
	require.LessOrEqual(t, len(portals), 100)

	for _, brand := range brands {
		if brand.Id == "db814755-2b62-4549-ba65-5138c0b80536" {
			require.Len(t, brand.PortalsIds, 2)
		} else if brand.Id == "e5079d5c-4526-4b03-a5d9-55db63065f94" {
			require.Len(t, brand.PortalsIds, 3)
		} else {
			require.Len(t, brand.PortalsIds, 1)
		}
	}
}

func TestCatalog_GetPortals_WithSandboxMode(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	portals, err := catalog.GetPortals(&opts)

	//assert
	require.NoError(t, err)
	require.LessOrEqual(t, len(portals), 100)

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
		_, err := definitions.GetSourceDefinition(definitions.GetSourceConfigOptions{EndpointId: endpointId, Env: pkg.FastenLighthouseEnvProduction})
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
		_, err := definitions.GetSourceDefinition(definitions.GetSourceConfigOptions{EndpointId: endpointId, Env: pkg.FastenLighthouseEnvSandbox})
		require.NoError(t, err)
	}

	//assert
	require.True(t, len(endpointPlatformTypes) >= 1)
	require.True(t, foundAllEndpointPlatfromTypes)
}

func TestCatalog_GetEndpoints_SuspendedEndpointsShouldBeRemoved(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "0143a953-44e9-416f-a506-4172ed426e3a",
	}
	_, err := catalog.GetEndpoints(&opts)
	require.Error(t, err, "endpoint with id 0143a953-44e9-416f-a506-4172ed426e3a not found")
}

func TestCatalog_GetPortal_KaiserMultipleEndpoints(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7",
	}
	portals, err := catalog.GetPortals(&opts)
	require.NoError(t, err, "endpoint with id 0143a953-44e9-416f-a506-4172ed426e3a not found")

	require.Len(t, portals, 1)
	for _, portal := range portals {
		require.Len(t, portal.EndpointIds, 2)
	}
}

func TestCatalog_GetPortal_KaiserMultipleEndpoints_Sandbox(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id:                "59673c08-e4b5-44d5-b5ab-532e69e8f7e7",
		LighthouseEnvType: pkg.FastenLighthouseEnvSandbox,
	}
	portals, err := catalog.GetPortals(&opts)
	require.NoError(t, err, "endpoint with id 0143a953-44e9-416f-a506-4172ed426e3a not found")

	require.Len(t, portals, 1)
	for _, portal := range portals {
		require.Len(t, portal.EndpointIds, 1)
	}
}
