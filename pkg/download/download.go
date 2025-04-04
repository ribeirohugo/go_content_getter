// Package download holds request download content methods
package download

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

// Content - Makes an HTTP request to a URL and gets the content in bytes format.
func Content(page model.Page) (model.File, error) {
	response, err := http.Get(page.URL) //nolint:gosec // received value needs to be a variable
	if err != nil {
		return model.File{}, fmt.Errorf("error making HTTP request to \"%s\": %s", contentURL, err.Error())
	}

	if response.StatusCode == http.StatusOK {
		// Read the response body into a byte slice
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return model.File{}, fmt.Errorf("error reading response body: %s", err.Error())
		}

		file := model.File{
			Title:   contentURL,
			Content: bodyBytes,
		}

		return file, nil
	}

	return model.File{}, fmt.Errorf("error response status: %s", response.Status)
}

// ContentBytes - Makes an HTTP request to a URL and gets the content in bytes format.
func ContentBytes(contentURL string) (model.File, error) {
	response, err := http.Get(contentURL) //nolint:gosec // received value needs to be a variable
	if err != nil {
		return model.File{}, fmt.Errorf("error making HTTP request to \"%s\": %s", contentURL, err.Error())
	}

	if response.StatusCode == http.StatusOK {
		// Read the response body into a byte slice
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return model.File{}, fmt.Errorf("error reading response body: %s", err.Error())
		}

		file := model.File{
			Title:   contentURL,
			Content: bodyBytes,
		}

		return file, nil
	}

	return model.File{}, fmt.Errorf("error response status: %s", response.Status)
}
