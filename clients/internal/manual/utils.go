package manual

import (
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/pkg"
	fhir401utils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"github.com/gabriel-vasile/mimetype"
	"io"
	"os"
	"strings"
)

/*
GetFileDocumentType returns the document type of given file.
We'll use various techniques to determine if the file is XML or JSON-like first.
If the file is JSON, we'll determine if it's a FHIR bundle.
If the file is NDJson, we'll parse the first entry to ensure it's a valid fhir resource.
We can't depend on just http.DetectContentType because it doesn't detect the difference between JSON, NDJSON, FHIR, XML and CCDAs.
*/
func GetFileDocumentType(file *os.File) (pkg.DocumentType, error) {

	defer file.Seek(0, io.SeekStart)

	file.Seek(0, io.SeekStart)
	mimetype.SetLimit(0)
	contentType, err := mimetype.DetectReader(file)
	if err != nil {
		return pkg.DocumentType("unknown"), err
	}
	contentTypeParts := strings.Split(contentType.String(), ";")

	if contentTypeParts[0] == "application/xml" || contentTypeParts[0] == "text/xml" {
		return pkg.DocumentTypeCCDA, nil
	} else if contentTypeParts[0] == "application/x-ndjson" {
		return pkg.DocumentTypeFhirNDJSON, nil
	} else if contentTypeParts[0] == "application/json" {
		//reset file to beginning
		file.Seek(0, io.SeekStart)

		//this is 1 or more json documents
		documentCount := 0
		containsFhirResource := false
		var parsedResource models.ResourceInterface
		d := json.NewDecoder(file)
		for {

			var resource json.RawMessage
			err := d.Decode(&resource)
			if err != nil {
				// io.EOF is expected at end of stream.
				if err == io.EOF {
					break //we're done
				} else {
					continue //skip this document, don't add it to the documentCount
				}
			}

			//we have a valid json document, lets check if it's a FHIR resource
			if !containsFhirResource {
				unknownResource, err := fhir401utils.MapToResource(resource, false)
				if err == nil {
					containsFhirResource = true
				}
				if castedResource, ok := unknownResource.(models.ResourceInterface); ok {
					parsedResource = castedResource
				}

			}

			documentCount++
		}

		if containsFhirResource && documentCount > 1 {
			return pkg.DocumentTypeFhirNDJSON, nil
		} else if containsFhirResource && documentCount == 1 && parsedResource != nil {
			primaryResourceType, _ := parsedResource.ResourceRef()
			if primaryResourceType == "Bundle" {
				return pkg.DocumentTypeFhirBundle, nil
			} else if primaryResourceType == "List" {
				return pkg.DocumentTypeFhirList, nil
			} else {
				return pkg.DocumentType("unknown"), fmt.Errorf("unknown FHIR Resource collection type: %s", primaryResourceType)
			}

		}
	}

	return pkg.DocumentType(contentTypeParts[0]), fmt.Errorf("unknown document type, content-type: %s", contentType)
}

// isWhitespaceChar reports whether the provided byte is a whitespace byte (0xWS)
// as defined in https://mimesniff.spec.whatwg.org/#terminology.
func isWhitespaceChar(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}
