package datatypes

import "github.com/go-playground/validator/v10"

type Address struct {
	// the lines of the address. For example, "123 Governors Ln".
	Line []string `json:"line,omitempty" yaml:"line,omitempty" validate:"omitempty,dive,min=1"`
	// the five-digit zip code.
	PostalCode string `json:"postal_code,omitempty" yaml:"postal_code,omitempty" validate:"omitempty,min=2"`
	City       string `json:"city,omitempty" yaml:"city,omitempty" validate:"omitempty,min=2"`
	// the two-letter state or possession abbreviation as defined in https://pe.usps.com/text/pub28/28apb.htm.
	State string `json:"state,omitempty" yaml:"state,omitempty" validate:"omitempty,max=3"` // iso3166_2 without prefix
	// the two-letter country code
	Country string `json:"country" yaml:"country" validate:"required,iso3166_1_alpha2,max=2"`
}

func (address *Address) Validate() error {
	valid := validator.New()
	err := valid.Struct(address)
	if err != nil {
		return err
	}

	return nil
}
