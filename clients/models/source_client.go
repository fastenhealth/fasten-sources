package models

import (
	"os"
)

//go:generate mockgen -source=interface.go -destination=mock/mock_client.go
type SourceClient interface {
	GetRequest(resourceSubpath string, decodeModelPtr interface{}) error
	SyncAll(db DatabaseRepository) error

	//Manual client ONLY functions
	SyncAllBundle(db DatabaseRepository, bundleFile *os.File) error
}
