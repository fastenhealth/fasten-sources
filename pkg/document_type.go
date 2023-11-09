package pkg

type DocumentType string

const (
	DocumentTypeCCDA       DocumentType = "CCDA"
	DocumentTypeFhirBundle DocumentType = "FHIR_BUNDLE"
	DocumentTypeFhirNDJSON DocumentType = "NDJSON"
)
