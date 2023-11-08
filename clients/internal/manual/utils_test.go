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
	contentType, err := GetFileDocumentType(file)

	//assert
	require.NoError(t, err)
	require.Equal(t, contentType, DocumentTypeCCDA)
}

func TestGetFileDocumentType(t *testing.T) {

	var documentTypeTests = []struct {
		fixturePath          string       // input
		expectedSuccess      bool         // expected result
		expectedDocumentType DocumentType // expected result
	}{
		{"testdata/fixtures/401-R4/bundle/synthea_Tania553_Harris789_545c2380-b77f-4919-ab5d-0f615f877250.json", true, DocumentTypeFhirBundle},
		{"testdata/fixtures/401-R4/ccda/MaryGrant-ClinicalSummary.xml", true, DocumentTypeCCDA},
		{"testdata/fixtures/401-R4/international-patient-summary/IPS-bundle-01.json", true, DocumentTypeFhirBundle},
		{"testdata/fixtures/401-R4/phr-ndjson-jsonl/JohnDoe.phr", true, DocumentTypeFhirNDJSON},
		{"testdata/fixtures/401-R4/phr-ndjson-jsonl/TimmySmart-FosterCareTimeline.phr", true, DocumentTypeFhirNDJSON},
		{"testdata/fixtures/README.md", false, DocumentType("")},
	}

	for _, tt := range documentTypeTests {
		file, err := os.OpenFile(tt.fixturePath, os.O_RDONLY, 0644)
		require.NoError(t, err)

		//test
		contentType, err := GetFileDocumentType(file)
		if tt.expectedSuccess {
			require.NoError(t, err)
			require.Equal(t, contentType, tt.expectedDocumentType, "fixture: %s", tt.fixturePath)
		} else {
			require.Error(t, err, "fixture: %s", tt.fixturePath)
		}
	}
}
