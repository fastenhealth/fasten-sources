package models

//go:generate mockgen -source=resource.go -destination=mock/mock_resource.go
type ResourceInterface interface {
	ResourceRef() (string, *string)
}
