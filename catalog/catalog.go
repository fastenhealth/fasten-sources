package catalog

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	"github.com/samber/lo"
	"strings"
)

//go:embed brands.json
var brandsFs embed.FS

//go:embed portals.json
var portalsFs embed.FS

//go:embed endpoints.json
var endpointsFs embed.FS

func GetBrands(opts *catalog.CatalogQueryOptions) (map[string]catalog.PatientAccessBrand, error) {
	brands, err := strictUnmarshalEmbeddedFile[catalog.PatientAccessBrand](brandsFs, "brands.json")
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}

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
			for brandId, _ := range brands {
				filteredBrand := brands[brandId]

				if len(filteredBrand.PortalsIds) == 1 {
					continue
				}

				filteredBrand.PortalsIds = lo.Filter(filteredBrand.PortalsIds, func(portalId string, ndx int) bool {
					_, err := GetPortals(&catalog.CatalogQueryOptions{Id: portalId, LighthouseEnvType: opts.LighthouseEnvType})
					if err != nil {
						return false
					} else {
						return true
					}
				})

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
	portals, err := strictUnmarshalEmbeddedFile[catalog.PatientAccessPortal](portalsFs, "portals.json")
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}

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
			for portalId, _ := range portals {
				filteredPortal := portals[portalId]

				if len(filteredPortal.EndpointIds) == 1 {
					continue
				}

				filteredPortal.EndpointIds = lo.Filter(filteredPortal.EndpointIds, func(endpointId string, ndx int) bool {
					_, err := GetEndpoints(&catalog.CatalogQueryOptions{Id: endpointId, LighthouseEnvType: opts.LighthouseEnvType})
					if err != nil {
						return false
					} else {
						return true
					}
				})

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

func GetEndpoints(opts *catalog.CatalogQueryOptions) (map[string]catalog.PatientAccessEndpoint, error) {
	endpoints, err := strictUnmarshalEmbeddedFile[catalog.PatientAccessEndpoint](endpointsFs, "endpoints.json")
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}

	if opts != nil {

		// filter by id if provided
		if len(opts.Id) > 0 {
			if endpoint, endpointOk := endpoints[opts.Id]; endpointOk {
				endpoints = map[string]catalog.PatientAccessEndpoint{endpoint.Id: endpoint}
			} else {
				return nil, fmt.Errorf("endpoint with id %s not found", opts.Id)
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

func GetPatientAccessInfoForLegacySourceType(legacySourceType string, legacyApiEndpoint string) (*catalog.PatientAccessBrand, *catalog.PatientAccessPortal, *catalog.PatientAccessEndpoint, pkg.FastenLighthouseEnvType, error) {
	brands, err := GetBrands(&catalog.CatalogQueryOptions{})
	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("error getting brands from catalog: %w", err)
	}

	portals, err := GetPortals(&catalog.CatalogQueryOptions{})
	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("error getting portals from catalog: %w", err)
	}

	endpoints, err := GetEndpoints(&catalog.CatalogQueryOptions{})
	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("error getting endpoints from catalog: %w", err)
	}

	// Find Portal

	matchingPortals := lo.PickBy(portals, func(key string, portal catalog.PatientAccessPortal) bool {
		_, found := lo.Find(portal.Identifiers, func(identifier datatypes.Identifier) bool {
			return identifier.Use == "fasten-legacy-source-type" && identifier.Value == legacySourceType
		})
		return found
	})

	if len(matchingPortals) == 0 {
		errMessage := fmt.Sprintf("No matching portal found for legacy source type: %s", legacySourceType)
		return nil, nil, nil, "", fmt.Errorf(errMessage)
	}

	if len(matchingPortals) > 1 {
		errMessage := fmt.Sprintf("Multiple matching portals found for legacy source type: %s vs %v", legacySourceType, lo.Keys(matchingPortals))
		return nil, nil, nil, "", fmt.Errorf(errMessage)
	}

	//found a portal, store it in the source credential
	matchingPortal := lo.Values(matchingPortals)[0]

	// lets find associated brand. if more than 1 brand is found, we will pick the first one
	matchingBrands := lo.PickBy(brands, func(key string, brand catalog.PatientAccessBrand) bool {
		return lo.Contains(brand.PortalsIds, matchingPortal.Id)
	})
	if len(matchingBrands) == 0 {
		errMessage := fmt.Sprintf("No matching brand found for portal: %s", matchingPortal.Id)
		return nil, nil, nil, "", fmt.Errorf(errMessage)
	}
	matchingBrand := lo.Values(matchingBrands)[0]

	//lets find the associated endpoint.
	matchingEndpoints := lo.PickByKeys(endpoints, matchingPortal.EndpointIds)
	if len(matchingEndpoints) == 0 {
		errMessage := fmt.Sprintf("No matching endpoint found for portal: %s", matchingPortal.Id)
		return nil, nil, nil, "", fmt.Errorf(errMessage)
	}

	//find endpoint by matching the sourceCredetial.ApiEndpointUrl with the endpoint url
	if len(matchingEndpoints) > 1 {
		filteredMatchingEndpoints := lo.PickBy(matchingEndpoints, func(key string, endpoint catalog.PatientAccessEndpoint) bool {
			return normalizeEndpointURL(endpoint.Url) == normalizeEndpointURL(legacyApiEndpoint)
		})
		if len(filteredMatchingEndpoints) == 1 {
			matchingEndpoints = filteredMatchingEndpoints
		}
	}
	// if more than 1 endpoint is found, we will filter any inactive & non-production endpoints
	if len(matchingEndpoints) > 1 {
		filteredMatchingEndpoints := lo.PickBy(matchingEndpoints, func(key string, endpoint catalog.PatientAccessEndpoint) bool {
			return endpoint.Status == "active" && lo.NoneBy(endpoint.Identifiers, func(identifier datatypes.Identifier) bool {
				return identifier.Use == "fasten-sandbox-mode" && identifier.Value == "true"
			})
		})
		if len(filteredMatchingEndpoints) == 1 {
			matchingEndpoints = filteredMatchingEndpoints
		}
	}

	//select the first endpoint
	matchingEndpoint := lo.Values(matchingEndpoints)[0]

	//find the environment
	var endpointEnv pkg.FastenLighthouseEnvType
	if _, isSandbox := lo.Find(matchingEndpoint.Identifiers, func(identifier datatypes.Identifier) bool {
		return identifier.Use == "fasten-sandbox-mode" && identifier.Value == "true"
	}); isSandbox {
		endpointEnv = pkg.FastenLighthouseEnvSandbox
	} else {
		endpointEnv = pkg.FastenLighthouseEnvProduction
	}

	return &matchingBrand, &matchingPortal, &matchingEndpoint, endpointEnv, nil
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

func normalizeEndpointURL(url string) string {
	normalized := url
	// for cases such as foobar.com
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		normalized = "https://" + normalized
	}

	if !strings.HasSuffix(url, "/") {
		normalized = normalized + "/"
	}
	return normalized
}
