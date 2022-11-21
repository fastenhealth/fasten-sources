package models

import "encoding/json"

//go:generate mockgen -source=resource.go -destination=mock/mock_resource.go
type ResourceInterface interface {
	ResourceRef() (string, *string)
}

type RawResourceFhir struct {
	SourceResourceType string          `json:"source_resource_type"`
	SourceResourceID   string          `json:"source_resource_id"`
	RawResource        json.RawMessage `json:"raw_resource"`
}
