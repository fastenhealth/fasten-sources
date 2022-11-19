package models

import (
	"context"
)

type DatabaseRepository interface {
	WrapRawResource(rawResourcePtr interface{})
	UpsertRawResource(ctx context.Context, sourceCredentials SourceCredential, rawResource ResourceInterface) error
}
