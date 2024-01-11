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
					return matchingBrand, nil
				}
			}
		}

		// filter by environment if provided
		if len(opts.LighthouseEnvType) > 0 {
			//we need to request the filtered portals, and then find the associated brands (stripping out any brands that don't have portals)

			// filtered portals is a map of sandbox or production portals
			filteredPortals, err := GetPortals(opts)
			if err != nil {
				return nil, fmt.Errorf("failed: %w", err)
			}

			//loop though all the portals, and search for their endpoint ids in the filtered endpoints list
			filteredBrands := map[string]catalog.PatientAccessBrand{}
			for brandId, brand := range brands {

				foundBrandPortals := lo.PickByKeys(filteredPortals, brand.PortalsIds)

				//if we have found (sandbox or prod) endpoints for this portal, then we update the endpoint ids, and add this portal to the filtered portals map
				if len(foundBrandPortals) > 0 {
					//we have endpoints for this portal, so we can add it to the filtered portals
					brand.PortalsIds = lo.Keys(foundBrandPortals)
					filteredBrands[brandId] = brand
				}
			}

			//reassign portals to the filtered portals
			brands = filteredBrands
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

			// filtered endpoints is a map of sandbox or production endpoints
			filteredEndpoints, err := GetEndpoints(opts)
			if err != nil {
				return nil, fmt.Errorf("failed getting endpoints for portal: %w", err)
			}

			//loop though all the portals, and search for their endpoint ids in the filtered endpoints list
			filteredPortals := map[string]catalog.PatientAccessPortal{}
			for portalId, portal := range portals {

				foundPortalEndpoints := lo.PickByKeys(filteredEndpoints, portal.EndpointIds)

				//if we have found (sandbox or prod) endpoints for this portal, then we update the endpoint ids, and add this portal to the filtered portals map
				if len(foundPortalEndpoints) > 0 {
					//we have endpoints for this portal, so we can add it to the filtered portals
					portal.EndpointIds = lo.Keys(foundPortalEndpoints)
					filteredPortals[portalId] = portal
				}
			}

			//reassign portals to the filtered portals
			portals = filteredPortals
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
			endpoints = lo.PickBy(endpoints, func(key string, value catalog.PatientAccessEndpoint) bool {
				_, sandboxMode := lo.Find(value.Identifiers, func(identifier datatypes.Identifier) bool {
					return identifier.Use == "fasten-sandbox-mode" && identifier.Value == "true"
				})

				if opts.LighthouseEnvType == pkg.FastenLighthouseEnvSandbox {
					return sandboxMode
				} else {
					return !sandboxMode
				}

			})
		}
	}

	if len(endpoints) == 0 {
		return nil, fmt.Errorf("no endpoints found")
	}
	return endpoints, nil
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
