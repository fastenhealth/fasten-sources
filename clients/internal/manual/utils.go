package manual

import (
	"encoding/json"
	"fmt"
	"github.com/fastenhealth/fasten-sources/pkg"
	fhir401utils "github.com/fastenhealth/gofhir-models/fhir401/utils"
	"io"
	"net/http"
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

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	defer file.Seek(0, io.SeekStart)

	contentType := http.DetectContentType(buffer)
	contentTypeParts := strings.Split(contentType, ";")

	if contentTypeParts[0] == "text/plain" {
		firstNonWS := 0
		//find first non-whitespace character
		for ; firstNonWS < len(buffer) && isWhitespaceChar(buffer[firstNonWS]); firstNonWS++ {
		}

		//determine the content-type using buffer
		if firstNonWS == len(buffer) {
			return pkg.DocumentType(contentTypeParts[0]), fmt.Errorf("file is empty")
		} else if buffer[firstNonWS] == '<' {
			//TODO: we should do some additional parsing to confirm
			return pkg.DocumentTypeCCDA, nil
		} else if buffer[firstNonWS] == '{' {
			//reset file to beginning
			file.Seek(0, io.SeekStart)

			//this is 1 or more json documents
			documentCount := 0
			containsFhirResource := false
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
					_, err := fhir401utils.MapToResource(resource, false)
					if err == nil {
						containsFhirResource = true
					}
				}

				documentCount++
			}

			if containsFhirResource && documentCount > 1 {
				return pkg.DocumentTypeFhirNDJSON, nil
			} else if containsFhirResource && documentCount == 1 {
				return pkg.DocumentTypeFhirBundle, nil
			}
		}
	}
	return pkg.DocumentType(contentTypeParts[0]), fmt.Errorf("unknown document type, content-type: %s", contentTypeParts[0])
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
