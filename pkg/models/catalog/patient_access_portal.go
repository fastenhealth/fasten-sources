package catalog

import (
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	validator "github.com/go-playground/validator/v10"
)

// TODO: generate via reflection
type PatientAccessPortal struct {
	// Fasten UUID for the portal
	Id string `json:"id" yaml:"id" validate:"required,uuid"`
	// List of identifiers for the organization, e.g., “GH1234”
	Identifiers []datatypes.Identifier `json:"identifiers,omitempty" yaml:"identifiers,omitempty" validate:"omitempty,dive"`
	// RFC3339 date & time of the last update to the patient portal’s information
	LastUpdated string `json:"last_updated" yaml:"last_updated" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	// Name of the patient portal, e.g., “MyChart”
	Name string `json:"name" yaml:"name" validate:"required,min=2,patient-access-brand-name"`
	// URL for the patient portal’s logo, which will be displayed on a card
	Logo string `json:"logo,omitempty" yaml:"logo,omitempty" validate:"omitempty,http_url"`
	// URL for the patient portal, where patients can manage accounts with this provider.
	PortalWebsite string `json:"portal_website,omitempty" yaml:"portal_website,omitempty" validate:"omitempty,http_url"`
	// Description of the patient portal, e.g., “Manage your health information with General Hospital”
	Description string `json:"description,omitempty" yaml:"description,omitempty" validate:"omitempty"`
	// List of endpoint IDs for the patient portal. This is used to associate the patient portal with the endpoints that are used to access it.
	EndpointIds []string `json:"endpoint_ids" yaml:"endpoint_ids" validate:"required,unique,dive,uuid"`
}

func (o *PatientAccessPortal) Validate() error {
	valid := validator.New()
	valid.RegisterValidation("patient-access-brand-name", PatientAccessBrandNameRegex)
	err := valid.Struct(o)
	if err != nil {
		return err
	}

	return nil
}
