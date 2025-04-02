// Package download holds request download content methods
package download

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// ContentBytes - Makes an HTTP request to a URL and
// gets the content in bytes format.
func ContentBytes(contentURL string) ([]byte, error) {
	response, err := http.Get(contentURL) //nolint:gosec // received value needs to be a variable
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

// MultipleContentBytes - Makes an HTTP request to a URL and downloads the page title and the URLs inside that page,
// to be downloaded.
// It returns the titlel, the slice of URLs and an error (or nil).
func TitleAndURLs(contentURL, titlePatterns, urlsPatterns string) (string, []string, error) {
	contentBytes, err := ContentBytes(contentURL)
	if err != nil {
		return "", []string{}, err
	}

	contentString := string(contentBytes)

	titleRegex := regexp.MustCompile(contentString)
	contentRegex := regexp.MustCompile(contentString)

	titleMatch := titleRegex.FindStringSubmatch(contentString)
	contentMatch := contentRegex.FindAllStringSubmatch(contentString, -1)

	title := ""
	if len(titleMatch) > 1 {
		title = titleMatch[1]
	}

	var urls []string
	for _, url := range contentMatch {
		urls = append(urls, url[1])
	}

	return title, urls, nil
}

// MultipleContentBytes - Makes an HTTP request to a URL download all the content from that page (for a given regex pattern)
// and returns a map with content bytes for a given content name.
// func MultipleContentBytes(contentURL, contentPattern string) (map[string][]byte, error) {
// 	contentBytes, err := ContentBytes(contentURL)
// 	if err != nil {
// 		return map[string][]byte{}, err
// 	}

// 	bodyString := string(contentBytes)

// 	contentRegex := regexp.MustCompile(contentPattern)

// 	contentMatch := contentRegex.FindAllStringSubmatch(bodyString, -1)
// 	// titleMatch := titleRegex.FindStringSubmatch(bodyString)

// 	var urls []string
// 	for _, url := range contentMatch {
// 		urls = append(urls, url[1])
// 	}

// }
