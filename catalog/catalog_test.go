package catalog

import (
	"fmt"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCatalog_GetBrands(t *testing.T) {
	//setup
	opts := catalog.CatalogQueryOptions{Id: "0000f30d-3987-4539-b743-5c263416a6cf"}

	//test
	brands, err := GetBrands(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, brands, 1)
}

func TestCatalog_GetBrands_WithInvalidId(t *testing.T) {
	//setup
	opts := catalog.CatalogQueryOptions{Id: "1"}

	//test
	_, err := GetBrands(&opts)

	//assert
	require.EqualError(t, err, fmt.Sprintf("brand with id %s not found", opts.Id))
}

func TestCatalog_GetBrands_WithSandboxMode(t *testing.T) {
	//setup
	opts := catalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	brands, err := GetBrands(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, brands, 20)

	for _, brand := range brands {
		require.Len(t, brand.PortalsIds, 1)
	}
}

func TestCatalog_GetPortals_WithSandboxMode(t *testing.T) {
	//setup
	opts := catalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	portals, err := GetPortals(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, portals, 20)

	for _, portal := range portals {
		require.Len(t, portal.EndpointIds, 1)
	}
}

func TestCatalog_GetEndpoints_WithSandboxMode(t *testing.T) {
	//setup
	opts := catalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}

	//test
	endpoints, err := GetEndpoints(&opts)

	//assert
	require.NoError(t, err)
	require.Len(t, endpoints, 20)
}
