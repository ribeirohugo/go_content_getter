// Package download holds request download content methods
package download

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

// Target - Makes an HTTP request to a URL and gets the content in bytes format.
func Target(target model.Target) (model.File, error) {
	response, err := http.Get(target.URL)
	if err != nil {
		return model.File{}, fmt.Errorf("error making HTTP request to \"%s\": %s", target.URL, err.Error())
	}

	switch response.StatusCode {
	case http.StatusOK:
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return model.File{}, fmt.Errorf("error reading response body: %s", err.Error())
		}

		file := model.File{
			Filename: target.Filename,
			Content:  bodyBytes,
		}

		return file, nil
	case http.StatusNotFound:
		log.Printf("File not found for URL: %s\n", target.URL)

		return model.File{}, nil
	}

	return model.File{}, fmt.Errorf("error response status: %s", response.Status)
}

// ManyTargets allows to download many model.Target, and returns their file data.
func ManyTargets(targets []model.Target) ([]model.File, error) {
	var files []model.File

	for i := range targets {
		file, err := Target(targets[i])
		if err != nil {
			return files, err
		}

		if file.Filename != "" {
			files = append(files, file)
		}
	}

	return files, nil
}
