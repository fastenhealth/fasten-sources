package catalog

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"testing"
)

func benchmarkGetEndpointsWithOptions(opts *modelsCatalog.CatalogQueryOptions, b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetEndpoints(opts)
	}
}

func benchmarkGetPortalsWithOptions(opts *modelsCatalog.CatalogQueryOptions, b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPortals(opts)
	}
}

func benchmarkGetBrandsWithOptions(opts *modelsCatalog.CatalogQueryOptions, b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetBrands(opts)
	}
}

// 56487095 ns/op
func Benchmark_GetEndpoint_NoFilter(b *testing.B) { benchmarkGetEndpointsWithOptions(nil, b) }

// 56367765 ns/op
func Benchmark_GetEndpoint_SingleEndpoint(b *testing.B) {
	benchmarkGetEndpointsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "9d0fa28a-0c5b-4065-9ee6-284ec9577a57"}, b)
}

// 55848260 ns/op
func Benchmark_GetEndpoint_Environment(b *testing.B) {
	benchmarkGetEndpointsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 182056250 ns/op
func Benchmark_GetPortals_NoFilter(b *testing.B) { benchmarkGetPortalsWithOptions(nil, b) }

// 181978688 ns/op
func Benchmark_GetPortals_SingleEndpoint(b *testing.B) {
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7"}, b)
}

// 300082958 ns/op - double time
func Benchmark_GetPortals_Environment(b *testing.B) {
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 194024736 ns/op
func Benchmark_GetPortals_Environment_WithEndpointCache(b *testing.B) {
	cache, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox})
	benchmarkGetPortalsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:     pkg.FastenLighthouseEnvSandbox,
		CachedEndpointsLookup: &cache,
	}, b)
}

// 203611925 ns/op
func Benchmark_GetBrands_NoFilter(b *testing.B) { benchmarkGetBrandsWithOptions(nil, b) }

// 205480108 ns/op
func Benchmark_GetBrands_SingleEndpoint(b *testing.B) {
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{Id: "59673c08-e4b5-44d5-b5ab-532e69e8f7e7"}, b)
}

// 3094676666 ns/op - 12 times slower
func Benchmark_GetBrands_Environment_Sandbox(b *testing.B) {
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvSandbox}, b)
}

// 267615969 ns/op
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

// 1413393208 ns/op
func Benchmark_GetBrands_Environment_Prod_WithEndpointCache(b *testing.B) {
	cacheEndpoints, _ := GetEndpoints(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction})
	cachePortals, _ := GetPortals(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: pkg.FastenLighthouseEnvProduction, CachedEndpointsLookup: &cacheEndpoints})
	benchmarkGetBrandsWithOptions(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:   pkg.FastenLighthouseEnvProduction,
		CachedPortalsLookup: &cachePortals,
	}, b)
}
