package utils

import (
	"context"
	"fmt"
	sourceModels "github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"maps"
	"os"
	"path/filepath"
	"time"
)

type StorageRepository struct {
	Logger        logrus.FieldLogger
	LocalFilepath string
	LocalFilename string
	JsonlWriter   JsonlWriter

	Cache map[string]bool

	ErrorData map[string]interface{}
}

func NewStorageRepository(logger logrus.FieldLogger) (*StorageRepository, error) {
	localFilename := fmt.Sprintf("%s-%s.jsonl", time.Now().Format(time.DateOnly), uuid.New().String())
	localFilepath := filepath.Join("/tmp", localFilename)

	storageFile, err := os.OpenFile(localFilepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open storage file: %w", err)
	}
	storageWriter := NewJsonlWriter(storageFile)

	return &StorageRepository{
		Logger:        logger,
		LocalFilepath: localFilepath,
		LocalFilename: localFilename,
		JsonlWriter:   storageWriter,
		Cache:         map[string]bool{},
		ErrorData:     map[string]interface{}{},
	}, nil
}

func (s *StorageRepository) Close() error {
	return s.JsonlWriter.Close()
}

func (s *StorageRepository) UpsertRawResource(ctx context.Context, sourceCredentials sourceModels.SourceCredential, rawResource sourceModels.RawResourceFhir) (bool, error) {
	resourceId := fmt.Sprintf("%s/%s", rawResource.SourceResourceType, rawResource.SourceResourceID)
	s.Logger.Infof("UpsertRawResource: %s", resourceId)

	if _, existsInCache := s.Cache[resourceId]; existsInCache {
		return false, nil
	} else {
		s.Cache[resourceId] = true

		err := s.JsonlWriter.Write(rawResource.ResourceRaw)
		return true, err
	}

}

func (s *StorageRepository) UpsertRawResourceAssociation(
	ctx context.Context,
	sourceId string,
	sourceResourceType string,
	sourceResourceId string,
	targetSourceId string,
	targetResourceType string,
	targetResourceId string,
) error {
	return nil
}

func (s *StorageRepository) BackgroundJobCheckpoint(ctx context.Context, checkpointData map[string]interface{}, errorData map[string]interface{}) {
	maps.Copy(s.ErrorData, errorData)
}
