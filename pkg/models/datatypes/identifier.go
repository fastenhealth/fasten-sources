package datatypes

import "github.com/go-playground/validator/v10"

type Identifier struct {
	//Identifier is a unique identifier of a thing.
	Use string `json:"use,omitempty" yaml:"use,omitempty" validate:"required,oneof=ext-platform-id ext-npi-number ext-npi-provider-taxonomy fasten-legacy-source-type fasten-endpoint-platform-override fasten-sandbox-mode"`
	// custom Fasten system or NPI http://hl7.org/fhir/sid/us-npi
	//NUCC taxonomy codes
	//https://npidb.org/taxonomy/
	// Taxonomy code mapping: http://www.wpc-edi.com/reference/codelists/healthcare/health-care-provider-taxonomy-code-set/
	// https://build.fhir.org/valueset-provider-taxonomy.html
	System string `json:"system" yaml:"system" validate:"required,min=2"`
	// Value is the unique identifier of the thing.
	Value string `json:"value" yaml:"value" validate:"required,min=2"`
}

func (o *Identifier) Validate() error {
	valid := validator.New()
	err := valid.Struct(o)
	if err != nil {
		return err
	}

	return nil
}
