package models

import (
	"encoding/json"
	"time"
)

//go:generate mockgen -source=resource.go -destination=mock/mock_resource.go
type ResourceInterface interface {
	ResourceRef() (string, *string)
}

type RawResourceFhir struct {
	SourceResourceType string          `json:"source_resource_type"`
	SourceResourceID   string          `json:"source_resource_id"`
	ResourceRaw        json.RawMessage `json:"resource_raw"`

	SortTitle           *string    `json:"sort_title"`
	SortDate            *time.Time `json:"sort_date"`
	ReferencedResources []string   `json:"referenced_resources"`
}
