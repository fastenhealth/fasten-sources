package catalog

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func benchmarkGetEndpointsWithOptions(opts *modelsCatalog.CatalogQueryOptions, b *testing.B) {
	for n := 0; n < b.N; n++ {
		g, _ := GetEndpoints(opts)
		b.Log("Retrieved endpoints:", len(g))
	}
}

func benchmarkGetPortalsWithOptions(opts *modelsCatalog.CatalogQueryOptions, b *testing.B) {
	for n := 0; n < b.N; n++ {
		g, _ := GetPortals(opts)
		b.Log("Retrieved portals:", len(g))
	}
}

func benchmarkGetBrandsWithOptions(opts *modelsCatalog.CatalogQueryOptions, b *testing.B) {
	for n := 0; n < b.N; n++ {
		g, _ := GetBrands(opts)
		b.Log("Retrieved brands:", len(g))
	}
}

// 5929 ns/op
// 1 of each call
func Benchmark_GetEndpoint_MultipleEndpoints(b *testing.B) {
	//we want to make sure the cache isn't being modified between runs
	for n := 0; n < b.N; n++ {
		g1, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{Id: "9d0fa28a-0c5b-4065-9ee6-284ec9577a57"})
		b.Log("Retrieved endpoints - CALL 1 (Epic):", len(g1))
		assert.NotZero(b, len(g1))
		g2, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{Id: "3290e5d7-978e-42ad-b661-1cf8a01a989c"})
		b.Log("Retrieved endpoints - CALL 2 (Cerner):", len(g2))
		assert.NotZero(b, len(g2))
	}
}

// 2682 ns/op
// 13938 results
func Benchmark_GetEndpoint_NoFilter(b *testing.B) { benchmarkGetEndpointsWithOptions(nil, b) }

// 2659 ns/op
// 1 result
func Benchmark_GetEndpoint_SingleEndpoint(b *testing.B) {
	benchmarkGetEndpointsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "9d0fa28a-0c5b-4065-9ee6-284ec9577a57"}, b)
}

// 1221454 ns/op
// 36 results
func Benchmark_GetEndpoint_Environment(b *testing.B) {
	benchmarkGetEndpointsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 6038 ns/op
// 1 of each call
func Benchmark_GetPortals_MultiplePortals(b *testing.B) {
	//we want to make sure the cache isn't being modified between runs
	for n := 0; n < b.N; n++ {
		g1, _ := GetPortals(&modelsCatalog.CatalogQueryOptions{Id: "2727ec27-67e9-475a-bea1-423102beaa1d"})
		b.Log("Retrieved endpoints - CALL 1 (Epic):", len(g1))
		assert.NotZero(b, len(g1))
		g2, _ := GetPortals(&modelsCatalog.CatalogQueryOptions{Id: "00a83214-7b14-4a12-ad95-5198b70dbb63"})
		b.Log("Retrieved endpoints - CALL 2 (Cerner):", len(g2))
		assert.NotZero(b, len(g2))
	}
}

// 2581 ns/op
// 40167 results
func Benchmark_GetPortals_NoFilter(b *testing.B) { benchmarkGetPortalsWithOptions(nil, b) }

// 2689 ns/op
// 1 result
func Benchmark_GetPortals_SingleEndpoint(b *testing.B) {
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7"}, b)
}

// 4968109 ns/op
// 37 results
func Benchmark_GetPortals_Environment(b *testing.B) {
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 4748811 ns/op
// 37 results
func Benchmark_GetPortals_Environment_WithEndpointCache(b *testing.B) {
	cache, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox})
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:     pkg.FastenLighthouseEnvSandbox,
		CachedEndpointsLookup: &cache,
	}, b)
}

// 2794283 ns/op
// 1 of each call
func Benchmark_GetBrands_MultipleBrands(b *testing.B) {
	//we want to make sure the cache isn't being modified between runs
	for n := 0; n < b.N; n++ {
		g1, _ := GetBrands(&modelsCatalog.CatalogQueryOptions{Id: "e16b9952-8885-4905-b2e3-b0f04746ed5c"})
		b.Log("Retrieved endpoints - CALL 1 (Epic):", len(g1))
		assert.NotZero(b, len(g1))
		g2, _ := GetBrands(&modelsCatalog.CatalogQueryOptions{Id: "a9da9380-0510-4026-9161-8ec238695c49"})
		b.Log("Retrieved endpoints - CALL 2 (Cerner):", len(g2))
		assert.NotZero(b, len(g2))
	}
}

// 2562 ns/op
// 33329 results
func Benchmark_GetBrands_NoFilter(b *testing.B) { benchmarkGetBrandsWithOptions(nil, b) }

// 2826065 ns/op
// 0 results
func Benchmark_GetBrands_SingleEndpoint(b *testing.B) {
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7"}, b)
}

// 8634036542 ns/op - 12 times slower
// 33 results
func Benchmark_GetBrands_Environment_Sandbox(b *testing.B) {
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 1022472167 ns/op
// 33 results
func Benchmark_GetBrands_Environment_Sandbox_WithEndpointCache(b *testing.B) {
	cacheEndpoints, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox})
	cachePortals, _ := GetPortals(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox, CachedEndpointsLookup: &cacheEndpoints})
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:   pkg.FastenLighthouseEnvSandbox,
		CachedPortalsLookup: &cachePortals,
	}, b)
}

// TODO: cannot run this, it takes ~10 minutes to run and it's not feasible
// 1033727667 ns/op
//func Benchmark_GetBrands_Environment_Prod(b *testing.B) {
//	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction}, b)
//}

// 4665917167 ns/op
// 33307 results
func Benchmark_GetBrands_Environment_Prod_WithEndpointCache(b *testing.B) {
	cacheEndpoints, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction})
	cachePortals, _ := GetPortals(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction, CachedEndpointsLookup: &cacheEndpoints})
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:   pkg.FastenLighthouseEnvProduction,
		CachedPortalsLookup: &cachePortals,
	}, b)
}

// 8561518084 ns/op
func Benchmark_GetBrandPortalEndpointUsingTEFCAIdentifiers_Multiple_Lookups(b *testing.B) {
	platformType := pkg.PlatformTypeEpic
	for n := 0; n < b.N; n++ {
		_, _, _, _, err := GetBrandPortalEndpointUsingTEFCAIdentifiers(platformType, "The Portland Clinic", "https://epicproxy.et4001.epichosted.com/APIProxyPRD/TPC/api/FHIR/R4/")
		if err != nil {
			b.Errorf("Error in GetBrandPortalEndpointUsingTEFCAIdentifiers: %v", err)
		}

		_, _, _, _, err2 := GetBrandPortalEndpointUsingTEFCAIdentifiers(platformType, "Midwestern University Clinics", "https://epicproxy.et1329.epichosted.com/APIProxyPRD/api/FHIR/R4/")
		if err2 != nil {
			b.Errorf("Error in GetBrandPortalEndpointUsingTEFCAIdentifiers: %v", err)
		}
	}
}
