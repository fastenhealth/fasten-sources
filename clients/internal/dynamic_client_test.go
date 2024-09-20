package internal

import (
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDynamicSourceClient_ImplementsInterface(t *testing.T) {
	t.Parallel()

	//assert
	require.Implements(t, (*models.SourceClient)(nil), &dynamicSourceClient{}, "should implement the models.SourceClient interface")
}
