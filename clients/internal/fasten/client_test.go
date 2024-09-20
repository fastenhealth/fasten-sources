package fasten

import (
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSourceClientFasten_ImplementsInterface(t *testing.T) {
	t.Parallel()

	//assert
	require.Implements(t, (*models.SourceClient)(nil), &FastenClient{}, "should implement the models.SourceClient interface")
}
