package pkg

type DocumentType string

const (
	DocumentTypeCCDA       DocumentType = "CCDA"
	DocumentTypeFhirBundle DocumentType = "FHIR_BUNDLE"
	DocumentTypeFhirList   DocumentType = "FHIR_LIST"
	DocumentTypeFhirNDJSON DocumentType = "NDJSON"
)
