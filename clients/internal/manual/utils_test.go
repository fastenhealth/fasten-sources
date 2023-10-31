package manual

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetFileContentType_Xml(t *testing.T) {
	t.Parallel()
	//setup
	file, err := os.OpenFile("testdata/fixtures/401-R4/ccda/MaryGrant-ClinicalSummary.xml", os.O_RDONLY, 0644)
	require.NoError(t, err)

	//test
	contentType, err := GetFileContentType(file)

	//assert
	require.NoError(t, err)
	require.Equal(t, contentType, "application/xml")
}
