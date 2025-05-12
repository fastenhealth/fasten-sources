package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/catalog"
	modelsCatalog "github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	"github.com/samber/lo"
	"log"
	"os"
	"strings"
)

// states they care about
// Georgia, Florida, Ohio, Texas and Kansas.
// GA
// FL
// OH
// TX
// KS

func main() {

	brands, err := catalog.GetBrands(&modelsCatalog.CatalogQueryOptions{})
	if err != nil {
		panic(err)
	}

	GABrands := map[string]modelsCatalog.PatientAccessBrand{}
	FLBrands := map[string]modelsCatalog.PatientAccessBrand{}
	OHBrands := map[string]modelsCatalog.PatientAccessBrand{}
	TXBrands := map[string]modelsCatalog.PatientAccessBrand{}
	KSBrands := map[string]modelsCatalog.PatientAccessBrand{}
	NationWideBrands := map[string]modelsCatalog.PatientAccessBrand{}

	//this is the number of health systems, that have a name of a certain length
	for bndx, _ := range brands {
		brand := brands[bndx]

		locations := lo.WithoutEmpty(lo.Uniq(lo.Map(brand.Locations, func(item datatypes.Address, index int) string {
			return item.State
		})))

		if len(locations) == 0 || len(locations) >= 50 {
			//if all states provided, or no states provided, then we don't need to include them
			locations = []string{"ALL"}
		}

		if lo.Contains(locations, "GA") {
			GABrands[brand.Id] = brand
		} else if lo.Contains(locations, "FL") {
			FLBrands[brand.Id] = brand
		} else if lo.Contains(locations, "OH") {
			OHBrands[brand.Id] = brand
		} else if lo.Contains(locations, "TX") {
			TXBrands[brand.Id] = brand
		} else if lo.Contains(locations, "KS") {
			KSBrands[brand.Id] = brand
		} else if lo.Contains(locations, "ALL") {
			NationWideBrands[brand.Id] = brand
		}
	}

	stats := map[string]int{
		"GA":         len(GABrands),
		"FL":         len(FLBrands),
		"OH":         len(OHBrands),
		"TX":         len(TXBrands),
		"KS":         len(KSBrands),
		"NationWide": len(NationWideBrands),
	}

	statsBytes, err := json.MarshalIndent(stats, "", "  ")
	log.Printf("Health System Name Stats: %s", string(statsBytes))

	if err := generateCSVFile("GA", GABrands); err != nil {
		log.Fatalf("failed to generate CSV file for GA: %s", err)
	}
	if err := generateCSVFile("FL", FLBrands); err != nil {
		log.Fatalf("failed to generate CSV file for FL: %s", err)
	}
	if err := generateCSVFile("OH", OHBrands); err != nil {
		log.Fatalf("failed to generate CSV file for OH: %s", err)
	}
	if err := generateCSVFile("TX", TXBrands); err != nil {
		log.Fatalf("failed to generate CSV file for TX: %s", err)
	}
	if err := generateCSVFile("KS", KSBrands); err != nil {
		log.Fatalf("failed to generate CSV file for KS: %s", err)
	}
	if err := generateCSVFile("NationWide", NationWideBrands); err != nil {
		log.Fatalf("failed to generate CSV file for NationWide: %s", err)
	}
}

func generateCSVFile(state string, brands map[string]modelsCatalog.PatientAccessBrand) error {
	records := [][]string{
		{"brand_id", "brand_name", "aliases"},
	}

	for brandId, _ := range brands {
		brand := brands[brandId]
		records = append(records, []string{
			brand.Id,
			brand.Name,
			strings.Join(lo.Compact(brand.Aliases), ", "),
		})
	}

	csvFile, err := os.OpenFile(fmt.Sprintf("%s_brands.csv", state), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return err
	}
	w := csv.NewWriter(csvFile)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	if err := csvFile.Close(); err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Printf("CSV file %s_brands.csv created successfully", state)
		return nil
	}
}
