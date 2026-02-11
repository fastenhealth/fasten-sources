package catalog

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	"github.com/samber/lo"
)

//go:embed brands.json
var brandsFs embed.FS

//go:embed portals.json
var portalsFs embed.FS

//go:embed endpoints.json
var endpointsFs embed.FS

// cached maps -- THESE MUST BE READ-ONLY after initial load
var brandsCache map[string]catalog.PatientAccessBrand
var portalsCache map[string]catalog.PatientAccessPortal
var endpointsCache map[string]catalog.PatientAccessEndpoint

func GetBrands(opts *catalog.CatalogQueryOptions) (map[string]catalog.PatientAccessBrand, error) {
	var err error
	if brandsCache == nil {
		brandsCache, err = strictUnmarshalEmbeddedFile[catalog.PatientAccessBrand](brandsFs, "brands.json")
		if err != nil {
			return nil, fmt.Errorf("failed: %w", err)
		}
	}
	brands := brandsCache

	if opts != nil {
		// filter by id if provided
		if len(opts.Id) > 0 {
			if brand, brandOk := brands[opts.Id]; brandOk {
				brands = map[string]catalog.PatientAccessBrand{brand.Id: brand}
			} else {

				//this brand id may have been merged into another brand, so we need to check the merged brand ids
				//fallback to looping though all the brands, and searching in the BrandIds array

				matchingBrand := lo.PickBy(brands, func(key string, value catalog.PatientAccessBrand) bool {
					return lo.Contains(value.BrandIds, opts.Id)
				})
				if len(matchingBrand) == 0 {
					return nil, fmt.Errorf("brand with id %s not found", opts.Id)
				} else if len(matchingBrand) > 1 {
					return nil, fmt.Errorf("multiple brands found with id %s", opts.Id)
				} else {
					brands = matchingBrand
				}
			}
		}

		// filter by environment if provided
		if len(opts.LighthouseEnvType) > 0 {
			//we need to request the endpoints, and then find the associated portals (stripping out any portals that don't have endpoints)

			sandboxModeExpectedValue := "false"
			if opts.LighthouseEnvType == pkg.FastenLighthouseEnvSandbox {
				sandboxModeExpectedValue = "true"
			}

			// pass 1: filter out brands that don't have the expected sandbox mode identifier
			brands = lo.PickBy(brands, func(key string, value catalog.PatientAccessBrand) bool {

				_, isMatchingEnv := lo.Find(value.Identifiers, func(identifier datatypes.Identifier) bool {
					return identifier.Use == "fasten-sandbox-mode" && identifier.Value == sandboxModeExpectedValue
				})
				return isMatchingEnv
			})

			// pass 2: if the brand has multiple portals, we need to filter out any endpoints that don't match the expected sandbox mode identifier

			filteredPortals, err := GetPortals(&catalog.CatalogQueryOptions{LighthouseEnvType: opts.LighthouseEnvType})
			if err != nil {
				return nil, fmt.Errorf("error getting portals from catalog: %w", err)
			}
			for brandId, _ := range brands {
				filteredBrand := brands[brandId]

				if len(filteredBrand.PortalsIds) == 1 {
					continue
				}

				if opts != nil && opts.CachedPortalsLookup != nil {
					//if we have a cached endpoints lookup, we can use that to filter out the endpoints
					filteredBrand.PortalsIds = lo.Keys(lo.PickByKeys(*opts.CachedPortalsLookup, filteredBrand.PortalsIds))
				} else {
					filteredBrand.PortalsIds = lo.Filter(filteredBrand.PortalsIds, func(portalId string, ndx int) bool {
						_, foundPortal := filteredPortals[portalId] // we only care if it exists in the filtered portals map
						return foundPortal
					})
				}

				//update
				brands[brandId] = filteredBrand
			}
		}
	}

	if len(brands) == 0 {
		return nil, fmt.Errorf("no brands found")
	}
	return brands, nil
}

func GetPortals(opts *catalog.CatalogQueryOptions) (map[string]catalog.PatientAccessPortal, error) {
	var err error
	if portalsCache == nil {
		portalsCache, err = strictUnmarshalEmbeddedFile[catalog.PatientAccessPortal](portalsFs, "portals.json")
		if err != nil {
			return nil, fmt.Errorf("failed: %w", err)
		}
	}
	portals := portalsCache

	if opts != nil {
		// filter by id if provided
		if len(opts.Id) > 0 {
			if portal, portalOk := portals[opts.Id]; portalOk {
				portals = map[string]catalog.PatientAccessPortal{portal.Id: portal}
			} else {
				return nil, fmt.Errorf("portal with id %s not found", opts.Id)
			}
		}

		// filter by environment if provided
		if len(opts.LighthouseEnvType) > 0 {
			//we need to request the endpoints, and then find the associated portals (stripping out any portals that don't have endpoints)

			sandboxModeExpectedValue := "false"
			if opts.LighthouseEnvType == pkg.FastenLighthouseEnvSandbox {
				sandboxModeExpectedValue = "true"
			}

			// pass 1: filter out portals that don't have the expected sandbox mode identifier
			portals = lo.PickBy(portals, func(key string, value catalog.PatientAccessPortal) bool {

				_, isMatchingEnv := lo.Find(value.Identifiers, func(identifier datatypes.Identifier) bool {
					return identifier.Use == "fasten-sandbox-mode" && identifier.Value == sandboxModeExpectedValue
				})
				return isMatchingEnv
			})

			// pass 2: if the portal has multiple endpoints, we need to filter out any endpoints that don't match the expected sandbox mode identifier
			filteredEndpoints, err := GetEndpoints(&catalog.CatalogQueryOptions{LighthouseEnvType: opts.LighthouseEnvType})
			if err != nil {
				return nil, fmt.Errorf("error getting endpoints from catalog: %w", err)
			}
			for portalId, _ := range portals {
				filteredPortal := portals[portalId]

				if len(filteredPortal.EndpointIds) == 1 {
					continue
				}

				if opts != nil && opts.CachedEndpointsLookup != nil {
					//if we have a cached endpoints lookup, we can use that to filter out the endpoints
					filteredPortal.EndpointIds = lo.Keys(lo.PickByKeys(*opts.CachedEndpointsLookup, filteredPortal.EndpointIds))
				} else {
					filteredPortal.EndpointIds = lo.Filter(filteredPortal.EndpointIds, func(endpointId string, ndx int) bool {
						_, found := filteredEndpoints[endpointId]
						return found
					})
				}

				//update
				portals[portalId] = filteredPortal
			}

		}

	}
	if len(portals) == 0 {
		return nil, fmt.Errorf("no portals found")
	}
	return portals, nil
}

// GetEndpoints returns a map of endpoints, filtered by the provided options.
// Note, the map keys may not match the endpoint IDs, as they may be aliases for a merged endpoint.
// eg.
// {"endpoint1": {"id": "endpoint2", "url": "https://endpoint2.com", "endpoint_ids": ["endpoint1"]}}
func GetEndpoints(opts *catalog.CatalogQueryOptions) (map[string]catalog.PatientAccessEndpoint, error) {
	var err error
	if endpointsCache == nil {
		endpointsCache, err = strictUnmarshalEmbeddedFile[catalog.PatientAccessEndpoint](endpointsFs, "endpoints.json")
		if err != nil {
			return nil, fmt.Errorf("failed: %w", err)
		}
	}
	endpoints := endpointsCache

	if opts != nil {
		// filter by id if provided
		if len(opts.Id) > 0 {
			if endpoint, endpointOk := endpoints[opts.Id]; endpointOk {
				endpoints = map[string]catalog.PatientAccessEndpoint{endpoint.Id: endpoint}
			} else {
				matchingEndpoint := lo.PickBy(endpoints, func(key string, value catalog.PatientAccessEndpoint) bool {
					return lo.Contains(value.EndpointIds, opts.Id)
				})
				if len(matchingEndpoint) == 0 {
					return nil, fmt.Errorf("endpoint with id %s not found", opts.Id)
				} else if len(matchingEndpoint) > 1 {
					return nil, fmt.Errorf("multiple endpoints found with id %s", opts.Id)
				} else {
					//since we're returning a map, we should also set an alias for the endpoint that was queried
					matchingEndpointList := lo.Values(matchingEndpoint) // this should only have 1 value, and now we can set the alias
					endpoints = map[string]catalog.PatientAccessEndpoint{opts.Id: matchingEndpointList[0]}
				}
			}
		}

		// filter by environment if provided
		if len(opts.LighthouseEnvType) > 0 {
			sandboxModeExpectedValue := "false"
			if opts.LighthouseEnvType == pkg.FastenLighthouseEnvSandbox {
				sandboxModeExpectedValue = "true"
			}

			endpoints = lo.PickBy(endpoints, func(key string, value catalog.PatientAccessEndpoint) bool {

				_, isMatchingEnv := lo.Find(value.Identifiers, func(identifier datatypes.Identifier) bool {
					return identifier.Use == "fasten-sandbox-mode" && identifier.Value == sandboxModeExpectedValue
				})

				return isMatchingEnv
			})
		}
	}

	if len(endpoints) == 0 {
		return nil, fmt.Errorf("no endpoints found")
	}
	return endpoints, nil
}

// the portal doesn't really matter here. We just care about Brand and associated Endpoint
// if we are able to successfully dtermine the endpoint, ALWAYS return it, even if we don't know the brand or portal. Endpoint is all we need for TEFCA Facilitated connections
func GetBrandPortalEndpointUsingTEFCAIdentifiers(platformType pkg.PlatformType, tefcaBrandName, tefcaUrl string) (*catalog.PatientAccessBrand, *catalog.PatientAccessPortal, *catalog.PatientAccessEndpoint, bool, error) {
	foundEndpoint := false
	var err error
	if endpointsCache == nil {
		endpointsCache, err = strictUnmarshalEmbeddedFile[catalog.PatientAccessEndpoint](endpointsFs, "endpoints.json")
		if err != nil {
			return nil, nil, nil, foundEndpoint, fmt.Errorf("failed: %w", err)
		}
	}
	endpoints := endpointsCache

	//get an endpoint (since they are unique on the URL) that matches the provided url
	tefcaUrl = NormalizeEndpointURL(tefcaUrl)
	endpoints = lo.PickBy(endpoints, func(key string, value catalog.PatientAccessEndpoint) bool {
		if value.GetPlatformType() != platformType {
			return false
		}
		return NormalizeEndpointURL(value.Url) == tefcaUrl
	})

	if len(endpoints) == 0 {
		return nil, nil, nil, foundEndpoint, fmt.Errorf("no endpoints found matching url: %s", tefcaUrl)
	}
	foundEndpoint = true
	if len(endpoints) > 1 {
		return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("multiple endpoints found matching url, returning first: %s", tefcaUrl)
	}

	//if we found an endpoint, lets try to find an associated portal and brand.
	if portalsCache == nil {
		portalsCache, err = GetPortals(&catalog.CatalogQueryOptions{
			LighthouseEnvType:     pkg.FastenLighthouseEnvProduction, //always production for RLS.
			CachedEndpointsLookup: &endpoints,
		})
		if err != nil {
			return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("error getting portals from catalog: %w", err)
		}
	}
	portals := portalsCache

	if brandsCache == nil {
		brandsCache, err = GetBrands(&catalog.CatalogQueryOptions{
			LighthouseEnvType:   pkg.FastenLighthouseEnvProduction, //always production for RLS.
			CachedPortalsLookup: &portals,
		})
		if err != nil {
			return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("error getting brands from catalog: %w", err)
		}
	}
	brands := brandsCache

	//now loop through the brands to find a matching brand name
	for brandId, _ := range brands {
		brand := brands[brandId]
		brandNames := lo.Uniq(append([]string{brand.Name}, brand.Aliases...))
		foundMatch := brandNamesMatch(brandNames, tefcaBrandName)
		if foundMatch {
			//we found a matching brand, now we need to find a portal that is associated with the endpoint
			for portalIdNdx, _ := range brand.PortalsIds {
				portalId := brand.PortalsIds[portalIdNdx]
				if portal, portalOk := portals[portalId]; portalOk {
					//check if the portal has the endpoint
					if lo.Contains(portal.EndpointIds, lo.Values(endpoints)[0].Id) {
						//we found a matching portal and endpoint
						return &brand, &portal, &lo.Values(endpoints)[0], foundEndpoint, nil
					}
				}
			}
		}
	}
	// no exact match found.
	// this is ok, because in the UI we'll use a fallback TEFCA brand if no exact match is found.
	// we'll still return the valid endpoint and the fact that it was found.

	// no brands found for this endpoint at all
	return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("no brand found matching name: %s", tefcaBrandName)
}

func brandNamesMatch(brandNames []string, tefcaBrandName string) bool {
	for ndx, _ := range brandNames {
		brandName := brandNames[ndx]
		if strings.EqualFold(tefcaBrandName, brandName) {
			return true
		}
	}
	return false
	//if !strings.EqualFold(tefcaBrandName, brandName) {
	//	return false
	//}
	//_, stateFound := lo.Find(brand.Locations, func(location datatypes.Address) bool {
	//	return strings.EqualFold(location.State, tefcaBrand.State)
	//})
	//return stateFound

}

//helpers

func strictUnmarshalEmbeddedFile[T catalog.PatientAccessBrand | catalog.PatientAccessPortal | catalog.PatientAccessEndpoint](embeddedFile embed.FS, embeddedFilename string) (map[string]T, error) {
	fileBytes, err := embeddedFile.ReadFile(embeddedFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded %s: %w", embeddedFilename, err)
	}

	//unmarshal the json into a map
	var patientAccessType map[string]T
	decoder := json.NewDecoder(bytes.NewReader(fileBytes))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&patientAccessType)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", embeddedFilename, err)
	}
	return patientAccessType, nil
}

func NormalizeEndpointURL(url string) string {
	normalized := url
	// for cases such as foobar.com
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		normalized = "https://" + normalized
	}

	if !strings.HasSuffix(url, "/") {
		normalized = normalized + "/"
	}
	normalized = strings.ToLower(normalized)
	return normalized
}
