package catalog_test

import (
	"fmt"
	"testing"

	"github.com/fastenhealth/fasten-sources/catalog"
	"github.com/fastenhealth/fasten-sources/definitions"
	"github.com/fastenhealth/fasten-sources/pkg"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
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
		require.LessOrEqual(t, len(portal.EndpointIds), 2)
	}
}

func TestCatalog_GetEndpoints_WithSandboxMode(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	endpoints, err := catalog.GetEndpoints(&opts)

	//assert
	require.NoError(t, err)
	require.LessOrEqual(t, len(endpoints), 100)
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
		_, err := definitions.GetSourceDefinition(definitions.WithEndpointId(endpointId), definitions.WithEnv(pkg.FastenLighthouseEnvProduction))
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
		_, err := definitions.GetSourceDefinition(
			definitions.WithEndpointId(endpointId),
			definitions.WithEnv(pkg.FastenLighthouseEnvSandbox),
		)
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

func TestCatalog_GetPortal_CernerMultipleEndpoints(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "00a54d25-9a00-4ebd-bc98-5a1b62958d91",
	}
	portals, err := catalog.GetPortals(&opts)
	require.NoError(t, err, "endpoint with id 1304ec75-b6ac-45ba-beef-f3657557c027, 4df464d1-00c6-4f53-bb71-2b53b7bd89d3 not found")

	require.Len(t, portals, 1)
	for _, portal := range portals {
		require.Len(t, portal.EndpointIds, 2)
	}
}

func TestCatalog_GetEndpoints_WithValidEndpoint_Id(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "33c76b01-4ac1-4889-8403-2be0d44f88b6",
	}
	endpoints, err := catalog.GetEndpoints(&opts)
	require.NoError(t, err)
	require.Len(t, endpoints, 1)
}

func TestCatalog_GetEndpoints_WithValidEndpointId(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "57209ef4-0e21-4e27-ae52-b23f395a30f5",
	}
	// 030f4652-65b5-4fa8-a005-4a2a2ed124b8
	endpoints, err := catalog.GetEndpoints(&opts)
	require.NoError(t, err)
	require.Len(t, endpoints, 1)
}

func TestCatalog_GetEndpoints_InValidEndpointId(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "57b8f926-f358-4bfe-b71b-e6eff720fbe4",
	}
	endpoints, err := catalog.GetEndpoints(&opts)
	require.Error(t, err, "endpoint with id 57b8f926-f358-4bfe-b71b-e6eff720fbe4 not found")
	require.Len(t, endpoints, 0)
}

func TestCatalog_GetEndpoints_WithKnownMergedId(t *testing.T) {
	//setup
	opts := modelsCatalog.CatalogQueryOptions{
		Id: "fc94bfc7-684d-4e4d-aa6e-ceec01c21c81",
	}
	endpoints, err := catalog.GetEndpoints(&opts)
	require.NoError(t, err, "endpoint with id fc94bfc7-684d-4e4d-aa6e-ceec01c21c81 not found")
	require.Len(t, endpoints, 1)
	require.NotEmpty(t, endpoints[opts.Id])
}

func TestGetBrandPortalEndpointUsingRLSIdentifiers(t *testing.T) {
	brand, portal, endpoint, foundEndpoint, err := catalog.GetBrandPortalEndpointUsingTEFCAIdentifiers(pkg.PlatformTypeEpic, "The Portland Clinic", "https://tpc-shield.tpcllp.com/FHIR/api/FHIR/R4/")

	require.NoError(t, err)
	require.True(t, foundEndpoint)
	require.Equal(t, "45c1a83e-fcdf-4bd5-8c87-8aaa3d4d7d5e", endpoint.Id)
	require.Equal(t, "96da4ce5-fc35-4e4e-a136-d60a724459ed", portal.Id)
	require.Equal(t, "eb9e8dea-d77a-45c8-b86a-76fa27cd19e6", brand.Id)

}
