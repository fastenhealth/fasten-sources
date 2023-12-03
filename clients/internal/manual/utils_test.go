package manual

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetFileDocumentType(t *testing.T) {

	var documentTypeTests = []struct {
		fixturePath          string           // input
		expectedSuccess      bool             // expected result
		expectedDocumentType pkg.DocumentType // expected result
	}{
		{"testdata/fixtures/401-R4/bundle/synthea_Tania553_Harris789_545c2380-b77f-4919-ab5d-0f615f877250.json", true, pkg.DocumentTypeFhirBundle},
		{"testdata/fixtures/401-R4/list/family-history.json", true, pkg.DocumentTypeFhirList},
		{"testdata/fixtures/401-R4/ccda/MaryGrant-ClinicalSummary.xml", true, pkg.DocumentTypeCCDA},
		{"testdata/fixtures/401-R4/international-patient-summary/IPS-bundle-01.json", true, pkg.DocumentTypeFhirBundle},
		{"testdata/fixtures/401-R4/phr-ndjson-jsonl/JohnDoe.phr", true, pkg.DocumentTypeFhirNDJSON},
		{"testdata/fixtures/401-R4/phr-ndjson-jsonl/TimmySmart-FosterCareTimeline.phr", true, pkg.DocumentTypeFhirNDJSON},
		{"testdata/fixtures/README.md", false, pkg.DocumentType("")},
	}

	for _, tt := range documentTypeTests {
		file, err := os.OpenFile(tt.fixturePath, os.O_RDONLY, 0644)
		require.NoError(t, err)

		//test
		contentType, err := GetFileDocumentType(file)
		if tt.expectedSuccess {
			require.NoError(t, err, "fixture: %s", tt.fixturePath)
			require.Equal(t, contentType, tt.expectedDocumentType, "fixture: %s", tt.fixturePath)
		} else {
			require.Error(t, err, "fixture: %s", tt.fixturePath)
		}
	}
}
