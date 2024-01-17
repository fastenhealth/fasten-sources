package catalog

import (
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	validator "github.com/go-playground/validator/v10"
	"regexp"
)

func PatientAccessBrandNameRegex(fl validator.FieldLevel) bool {
	return regexp.MustCompile("^[a-zA-Z0-9-+'()&.,#:!@/ _]+$").MatchString(fl.Field().String())
}

// TODO: generate via reflection
type PatientAccessBrand struct {
	// Fasten UUID for the brand - id should be a unique identifier for the brand. It is globally unique and should be a UUID
	Id string `json:"id" yaml:"id" validate:"required,uuid"`
	// List of identifiers for the organization, e.g., external system, etc NPI, etc
	// Identifiers SHOULD include a platform identifier, so we know where this entry came from, but not required
	Identifiers []datatypes.Identifier `json:"identifiers,omitempty" yaml:"identifiers,omitempty" validate:"omitempty,dive"`
	// RFC3339 Date and time the organization was last updated - Timestamp should be the last updated datetime for the data from this source, not the current date
	LastUpdated string `json:"last_updated" yaml:"last_updated" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	// Primary name for the organization to display on a card, e.g., “General Hospital”
	// Note this is not used within the app, only the Portal name should be used.
	Name string `json:"name" yaml:"name" validate:"required,min=2,patient-access-brand-name"`
	// URL for the organization’s primary website. note this is distinct from a patient portal, described under “Patient Access Details” below
	BrandWebsite string `json:"brand_website,omitempty" yaml:"brand_website,omitempty" validate:"omitempty,http_url"`
	// URL for the organization’s logo, which will be displayed on a card, Note this is a fallback logo, the primary logo will always be the Portal logo
	Logo string `json:"logo,omitempty" yaml:"logo,omitempty" validate:"omitempty,http_url"`
	// List of alternate names for the organization, e.g., “GH”, “General”, “GH Hospital”
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty" validate:"omitempty,unique,dive"`
	// List of locations for the organization
	// These should be the locations where the organization has a physical presence, e.g., a hospital or clinic"
	Locations []datatypes.Address `json:"locations,omitempty" yaml:"locations,omitempty" validate:"omitempty,dive"`
	// Patient Access Details
	// These must be references to Patient Access Portal resource Ids
	PortalsIds []string `json:"portal_ids" yaml:"portal_ids" validate:"required,unique,dive,uuid"`

	// list of brand ids that were merged together to creat this brand
	BrandIds []string `json:"brand_ids,omitempty" validate:"required,dive,uuid"`
}

func (o *PatientAccessBrand) Validate() error {
	valid := validator.New()
	valid.RegisterValidation("patient-access-brand-name", PatientAccessBrandNameRegex)
	err := valid.Struct(o)
	if err != nil {
		return err
	}

	return nil
}
