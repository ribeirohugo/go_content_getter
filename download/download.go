// Package download holds request download content methods
package download

import (
	"fmt"
	"io"
	"net/http"
)

// ContentBytes - Makes an HTTP request to a URL and
// gets the content in bytes format, according to a given content regex pattern.
func ContentBytes(contentURL string) ([]byte, error) {
	response, err := http.Get(contentURL)
	if err != nil {
		return []byte{}, fmt.Errorf("error making HTTP request to \"%s\": %s", contentURL, err.Error())
	}

	if response.StatusCode == http.StatusOK {
		// Read the response body into a byte slice
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return []byte{}, fmt.Errorf("error reading response body: %s", err.Error())
		}

		return bodyBytes, nil
	}

	return []byte{}, fmt.Errorf("error response status: %s", response.Status)
}
