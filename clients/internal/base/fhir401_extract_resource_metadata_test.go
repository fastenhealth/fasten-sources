package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveDuplicateStr(t *testing.T) {
	t.Parallel()
	//setup
	listOfStrings := []string{"a", "b", "c", "a", "b", "c", "d", "e", "f", "g", "h"}

	//test
	uniqueStrings := removeDuplicateStr(listOfStrings)

	//assert
	require.Equal(t, 8, len(uniqueStrings))
	require.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h"}, uniqueStrings)
}

func TestNormalizeReferenceId(t *testing.T) {
	t.Parallel()
	//setup
	var tests = []struct {
		inputReferenceId              string
		internalReferenceMap          map[string]string
		expectedNormalizedReferenceId string
	}{
		// the table itself
		{
			"urn:uuid:801922ee-1eaa-70ab-96ef-27a226ba82d3",
			map[string]string{"urn:uuid:801922ee-1eaa-70ab-96ef-27a226ba82d3": "Patient/801922ee-1eaa-70ab-96ef-27a226ba82d3"},
			"Patient/801922ee-1eaa-70ab-96ef-27a226ba82d3",
		},
		{
			"Patient/4a085566-49d8-fa6f-d934-8494253b3148",
			map[string]string{},
			"Patient/4a085566-49d8-fa6f-d934-8494253b3148",
		},
		{
			"Patient/4a085566-49d8-fa6f-d934-8494253b3148/_history/1",
			map[string]string{},
			"Patient/4a085566-49d8-fa6f-d934-8494253b3148",
		},
		{
			"http://example.org/fhir/Observation/1x2/_history/2",
			map[string]string{},
			"http://example.org/fhir/Observation/1x2/_history/2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputReferenceId, func(t *testing.T) {
			//test
			ans := normalizeReferenceId(tt.inputReferenceId, tt.internalReferenceMap)

			//assert
			require.Equal(t, tt.expectedNormalizedReferenceId, ans)
		})
	}
}
