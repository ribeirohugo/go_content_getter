package page

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

// GetHTTP - Makes an HTTP request to a URL and gets the content in bytes format.
func GetHTTP(contentURL string) (model.Page, error) {
	response, err := http.Get(contentURL) //nolint:gosec // received value needs to be a variable
	if err != nil {
		return model.Page{}, fmt.Errorf("error making HTTP request to \"%s\": %s", contentURL, err.Error())
	}

	if response.StatusCode == http.StatusOK {
		// Read the response body into a byte slice
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Page{}, fmt.Errorf("error reading response body: %s", err.Error())
		}

		pageTitle, err := titleFromHTTP(bodyBytes)
		if err != nil {
			pageTitle = ""
		}

		page := model.Page{
			Title:   pageTitle,
			Content: bodyBytes,
		}

		return page, nil
	}

	return model.Page{}, fmt.Errorf("error response status: %s", response.Status)
}

// titleFromHTTP gets title from HTTP page.
func titleFromHTTP(html []byte) (string, error) {
	re := regexp.MustCompile(patterns.HTMLTitle2)
	matches := re.FindSubmatch(html)
	if len(matches) >= 2 {
		return string(matches[1]), nil
	}
	return "", model.ErrTitleNotFound
}
