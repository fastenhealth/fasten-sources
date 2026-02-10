package catalog_details

import (
	"embed"
	"fmt"
	"github.com/fastenhealth/fasten-sources/catalog"
	"github.com/fastenhealth/fasten-sources/pkg"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	"strings"

	"github.com/samber/lo"
)

//go:embed brand_details.json
var brandDetailsFs embed.FS

var brandDetailsCache map[string]modelsCatalog.PatientAccessBrandDetails

func GetBrandDetails(opts *modelsCatalog.CatalogQueryOptions) (map[string]modelsCatalog.PatientAccessBrandDetails, error) {
	var err error
	if brandDetailsCache == nil {
		brandDetailsCache, err = catalog.StrictUnmarshalEmbeddedFile[modelsCatalog.PatientAccessBrandDetails](brandDetailsFs, "brand_details.json")
		if err != nil {
			return nil, fmt.Errorf("failed: %w", err)
		}
	}
	brands := brandDetailsCache

	if opts != nil {
		// filter by id if provided
		if len(opts.Id) > 0 {
			if brand, brandOk := brands[opts.Id]; brandOk {
				brands = map[string]modelsCatalog.PatientAccessBrandDetails{brand.Id: brand}
			} else {

				//this brand id may have been merged into another brand, so we need to check the merged brand ids
				//fallback to looping though all the brands, and searching in the BrandIds array

				matchingBrand := lo.PickBy(brands, func(key string, value modelsCatalog.PatientAccessBrandDetails) bool {
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
			brands = lo.PickBy(brands, func(key string, value modelsCatalog.PatientAccessBrandDetails) bool {

				_, isMatchingEnv := lo.Find(value.Identifiers, func(identifier datatypes.Identifier) bool {
					return identifier.Use == "fasten-sandbox-mode" && identifier.Value == sandboxModeExpectedValue
				})
				return isMatchingEnv
			})

			// pass 2: if the brand has multiple portals, we need to filter out any endpoints that don't match the expected sandbox mode identifier

			filteredPortals, err := catalog.GetPortals(&modelsCatalog.CatalogQueryOptions{LighthouseEnvType: opts.LighthouseEnvType})
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

// the portal doesn't really matter here. We just care about Brand and associated Endpoint
// if we are able to successfully dtermine the endpoint, ALWAYS return it, even if we don't know the brand or portal. Endpoint is all we need for TEFCA Facilitated connections
func GetBrandPortalEndpointUsingTEFCAIdentifiers(platformType pkg.PlatformType, tefcaBrandName, tefcaUrl string) (*modelsCatalog.PatientAccessBrandDetails, *modelsCatalog.PatientAccessPortal, *modelsCatalog.PatientAccessEndpoint, bool, error) {
	foundEndpoint := false
	var err error

	endpoints, err := catalog.GetEndpoints(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType: pkg.FastenLighthouseEnvProduction, //always production for RLS.
	})

	//get an endpoint (since they are unique on the URL) that matches the provided url
	tefcaUrl = normalizeEndpointURL(tefcaUrl)
	endpoints = lo.PickBy(endpoints, func(key string, value modelsCatalog.PatientAccessEndpoint) bool {
		if value.GetPlatformType() != platformType {
			return false
		}
		return normalizeEndpointURL(value.Url) == tefcaUrl
	})

	if len(endpoints) == 0 {
		return nil, nil, nil, foundEndpoint, fmt.Errorf("no endpoints found matching url: %s", tefcaUrl)
	}
	foundEndpoint = true
	if len(endpoints) > 1 {
		return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("multiple endpoints found matching url, returning first: %s", tefcaUrl)
	}

	//if we found an endpoint, lets try to find an associated portal and brand.
	portals, err := catalog.GetPortals(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:     pkg.FastenLighthouseEnvProduction, //always production for RLS.
		CachedEndpointsLookup: &endpoints,
	})
	if err != nil {
		return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("error getting portals from catalog: %w", err)
	}

	brands, err := GetBrandDetails(&modelsCatalog.CatalogQueryOptions{
		LighthouseEnvType:   pkg.FastenLighthouseEnvProduction, //always production for RLS.
		CachedPortalsLookup: &portals,
	})
	if err != nil {
		return nil, nil, &lo.Values(endpoints)[0], foundEndpoint, fmt.Errorf("error getting brands from catalog: %w", err)
	}

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

func normalizeEndpointURL(url string) string {
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
