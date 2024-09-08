package main

import (
	"encoding/json"
	"github.com/fastenhealth/fasten-sources/catalog"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"log"
)

func main() {

	brands, err := catalog.GetBrands(&modelsCatalog.CatalogQueryOptions{})
	if err != nil {
		panic(err)
	}

	portals, err := catalog.GetPortals(&modelsCatalog.CatalogQueryOptions{})
	if err != nil {
		panic(err)
	}

	// Print the brands
	log.Printf("Total Brands: %v", len(brands))
	log.Printf("Total Portals: %v", len(portals))

	//this is the number of health systems, that have a name of a certain length
	healthSystemNameStats := map[int]int{}
	for bndx, _ := range brands {
		brand := brands[bndx]

		if _, ok := healthSystemNameStats[len(brand.Name)]; !ok {
			healthSystemNameStats[len(brand.Name)] = 0
		}
		healthSystemNameStats[len(brand.Name)]++

		for andx, _ := range brand.Aliases {
			alias := brand.Aliases[andx]
			if _, ok := healthSystemNameStats[len(alias)]; !ok {
				healthSystemNameStats[len(alias)] = 0
			}
			healthSystemNameStats[len(alias)]++
		}
	}

	brandBytes, err := json.MarshalIndent(healthSystemNameStats, "", "  ")
	log.Printf("Health System Name Stats: %s", string(brandBytes))

}
