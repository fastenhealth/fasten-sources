package catalog

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/fasten-sources/pkg/models/datatypes"
	validator "github.com/go-playground/validator/v10"
)

// TODO: generate via reflection
type PatientAccessEndpoint struct {
	// Fasten UUID for the endpoint
	Id string `json:"id" yaml:"id" validate:"required,uuid"`
	// List of identifiers for the endpoint, e.g., “GH1234”
	Identifiers []datatypes.Identifier `json:"identifiers,omitempty" yaml:"identifiers,omitempty" validate:"omitempty,dive"`
	// RFC3339 Date and time the endpoint was last updated
	LastUpdated string `json:"last_updated" yaml:"last_updated" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	// Status of the endpoint, e.g., “active” - http://terminology.hl7.org/CodeSystem/endpoint-status
	Status string `json:"status" yaml:"status" validate:"required,oneof=active suspended"`
	// Connection type for the endpoint, e.g., “hl7-fhir-rest” - http://terminology.hl7.org/CodeSystem/endpoint-connection-type
	ConnectionType string `json:"connection_type" yaml:"connection_type" validate:"required,oneof=hl7-fhir-rest dicom-wado-rs"`
	// Platform type for the endpoint, e.g., “epic”, "cerner"
	PlatformType string `json:"platform_type" yaml:"platform_type" validate:"required"`
	// URL for the endpoint, must have trailing slash
	Url string `json:"url" yaml:"url" validate:"required,http_url,endswith=/"`

	//oauth endpoints
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty" yaml:"authorization_endpoint,omitempty" validate:"required_if=PatientAccessEndpoint.Status active,omitempty,http_url"`
	TokenEndpoint         string `json:"token_endpoint,omitempty" yaml:"token_endpoint,omitempty" validate:"required_if=PatientAccessEndpoint.Status active,omitempty,http_url"`
	IntrospectionEndpoint string `json:"introspection_endpoint,omitempty" yaml:"introspection_endpoint,omitempty" validate:"omitempty,http_url"`
	UserInfoEndpoint      string `json:"userinfo_endpoint,omitempty" yaml:"userinfo_endpoint,omitempty" validate:"omitempty,http_url"`
	//optional - required when Dynamic Client Registration mode is set
	RegistrationEndpoint string `json:"registration_endpoint,omitempty" yaml:"registration_endpoint,omitempty" validate:"omitempty,http_url"`

	//Fasten custom configuration
	FhirVersion           string `json:"fhir_version,omitempty" yaml:"fhir_version,omitempty" validate:"omitempty"`
	SmartConfigurationUrl string `json:"smart_configuration_url,omitempty" yaml:"smart_configuration_url,omitempty" validate:"omitempty,http_url"`
	FhirCapabilitiesUrl   string `json:"fhir_capabilities_url,omitempty" yaml:"fhir_capabilities_url,omitempty" validate:"omitempty,http_url"`

	//Software info
	SoftwareName        string `json:"software_name,omitempty" yaml:"software_name,omitempty" validate:"omitempty"`
	SoftwareVersion     string `json:"software_version,omitempty" yaml:"software_version,omitempty" validate:"omitempty"`
	SoftwareReleaseDate string `json:"software_release_date,omitempty" yaml:"software_release_date,omitempty" validate:"omitempty"`
}

func (o *PatientAccessEndpoint) Validate() error {
	valid := validator.New()
	err := valid.Struct(o)
	if err != nil {
		return err
	}

	return nil
}

// GetPlatformType returns the platform type for the endpoint, taking into account overrides
func (o *PatientAccessEndpoint) GetPlatformType() pkg.SourceType {
	platformType := o.PlatformType

	//check if there's an overrride
	if o.Identifiers != nil {
		for _, identifier := range o.Identifiers {
			if identifier.Use == "fasten-endpoint-platform-override" {
				platformType = identifier.Value
			}
		}
	}
	return pkg.SourceType(platformType)
}
