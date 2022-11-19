package models

type ResourceInterface interface {
	ResourceRef() (string, *string)
}
