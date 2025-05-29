package pkg

type SourceCredentialType string

const (
	SourceCredentialTypeSmartOnFhir      SourceCredentialType = "smart_on_fhir"     // SMART on FHIR authorization (default)
	SourceCredentialTypeTefcaDirect      SourceCredentialType = "tefca_direct"      // TEFCA Direct authorization
	SourceCredentialTypeTefcaFacilitated SourceCredentialType = "tefca_facilitated" // TEFCA Facilitated FHIR authorization (SMART on FHIR under the hood, but initially populated by TEFCA RLS response)

)
