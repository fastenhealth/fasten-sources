package manual

import (
	"net/http"
	"os"
	"strings"
)

/*
GetFileContentType returns the content type of given file by sniffing out fist 512 bytes of the file
We can't depend on just http.DetectContentType because it doesn't detect the difference between JSON and XML.
Instead we'll try to determine the content type ourselves.
*/
func GetFileContentType(file *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	contentTypeParts := strings.Split(contentType, ";")

	if contentTypeParts[0] == "text/plain" {
		firstNonWS := 0
		//find first non-whitespace character
		for ; firstNonWS < len(buffer) && isWhitespaceChar(buffer[firstNonWS]); firstNonWS++ {
		}

		//determine the content-type using buffer
		if firstNonWS == len(buffer) {
			return "text/plain", nil
		} else if buffer[firstNonWS] == '<' {
			return "application/xml", nil
		} else if buffer[firstNonWS] == '{' || buffer[firstNonWS] == '[' {
			return "application/json", nil
		}
		return "text/plain", nil
	}
	return contentTypeParts[0], nil

	//if contentTypeParts[0] == "text/plain" {
	//	// If the content type is text/plain, we'll try to determine the content type ourselves.
	//	// We'll check if the first non-whitespace character is '<' or '{' or '['.
	//	// If it is '<', we'll assume it's XML.
	//	// If it is '{' or '[', we'll assume it's JSON.
	//	// Otherwise, we'll assume it's text/plain.
	//	firstNonWS := 0
	//	for ; firstNonWS < len(buffer) && isWhitespaceChar(buffer[firstNonWS]); firstNonWS++ {
	//		if firstNonWS == len(buffer) {
	//			// The file is empty.
	//			return "text/plain", nil
	//		} else if buffer[firstNonWS] == '<' {
	//			return "application/xml", nil
	//		} else if buffer[firstNonWS] == '{' || buffer[firstNonWS] == '[' {
	//			return "application/json", nil
	//		} else {
	//			return "text/plain", nil
	//		}
	//
	//	}
	//}

	//return contentType, nil
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
