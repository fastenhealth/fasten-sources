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

// 252661950 ns/op
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

// 56487095 ns/op
// 13938 results
func Benchmark_GetEndpoint_NoFilter(b *testing.B) { benchmarkGetEndpointsWithOptions(nil, b) }

// 56367765 ns/op
// 1 result
func Benchmark_GetEndpoint_SingleEndpoint(b *testing.B) {
	benchmarkGetEndpointsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "9d0fa28a-0c5b-4065-9ee6-284ec9577a57"}, b)
}

// 55848260 ns/op
// 36 results
func Benchmark_GetEndpoint_Environment(b *testing.B) {
	benchmarkGetEndpointsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 559879958 ns/op
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

// 182056250 ns/op
// 40167
func Benchmark_GetPortals_NoFilter(b *testing.B) { benchmarkGetPortalsWithOptions(nil, b) }

// 278172469 ns/op
// 1 result
func Benchmark_GetPortals_SingleEndpoint(b *testing.B) {
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7"}, b)
}

// 539327104 ns/op - double time
// 37 results
func Benchmark_GetPortals_Environment(b *testing.B) {
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 316843417 ns/op
// 37 results
func Benchmark_GetPortals_Environment_WithEndpointCache(b *testing.B) {
	cache, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox})
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:     pkg.FastenLighthouseEnvSandbox,
		CachedEndpointsLookup: &cache,
	}, b)
}

// 1079450167 ns/op
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

// 504037062 ns/op
// 33329 results
func Benchmark_GetBrands_NoFilter(b *testing.B) { benchmarkGetBrandsWithOptions(nil, b) }

// 519999334 ns/op
// 0 results
func Benchmark_GetBrands_SingleEndpoint(b *testing.B) {
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7"}, b)
}

// 8634036542 ns/op - 12 times slower
// 33 results
func Benchmark_GetBrands_Environment_Sandbox(b *testing.B) {
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 1007805916 ns/op
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
//func Benchmark_GetBrands_Environment_Prod(b *testing.B) {
//	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction}, b)
//}

// 4728215542 ns/op
// 33307 results
func Benchmark_GetBrands_Environment_Prod_WithEndpointCache(b *testing.B) {
	cacheEndpoints, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction})
	cachePortals, _ := GetPortals(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction, CachedEndpointsLookup: &cacheEndpoints})
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:   pkg.FastenLighthouseEnvProduction,
		CachedPortalsLookup: &cachePortals,
	}, b)
}
