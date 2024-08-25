// Package getter holds content structs and methods of getting content
package getter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/ribeirohugo/go_content_getter/download"
	"github.com/ribeirohugo/go_content_getter/patterns"
)

const (
	defaultContentRegex = patterns.ImageContentFromHrefURL
	defaultTitleRegex   = patterns.HTMLTitle
	defaultFolderName   = "content"

	filePermissions = 0750
)

// Getter holds content content Getter struct
type Getter struct {
	contentRegex string
	titleRegex   string
	path         string
	url          string
}

// New is a Getter constructor. It requires:
// A url string from a web page to look for content.
// A path string to define where to store fetched content. (Optional field)
// A contentRegex to select to download. (Optional field)
// A titleRegex to select folder title to fetched content. (Optional field)
func New(url string, path string, contentRegex string, titleRegex string) Getter {
	contentRegexExpression := defaultContentRegex
	if contentRegex != "" {
		contentRegexExpression = contentRegex
	}

	titleRegexExpression := defaultTitleRegex
	if titleRegex != "" {
		titleRegexExpression = titleRegex
	}

	return Getter{
		contentRegex: contentRegexExpression,
		titleRegex:   titleRegexExpression,
		path:         path,
		url:          url,
	}
}

// Get returns slice with all images URL, page title
// If any error occurs it returns empty
func (g Getter) Get() ([]string, string, error) {
	return g.GetFromURL(g.url)
}

// GetFromURL returns slice with all images URL, page title
// Requires Url to get content from
// If any error occurs it returns empty
func (g Getter) GetFromURL(url string) ([]string, string, error) {
	response, err := http.Get(url) //nolint:gosec // received value needs to be a variable
	if err != nil {
		return []string{}, "", nil
	}

	if response.StatusCode != http.StatusOK {
		return []string{}, "", nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return []string{}, "", err
	}

	bodyString := string(bodyBytes)

	contentRegexString := g.contentRegex
	if contentRegexString == "" {
		contentRegexString = defaultContentRegex
	}

	titleRegexString := g.titleRegex
	if titleRegexString == "" {
		titleRegexString = defaultTitleRegex
	}

	contentRegex := regexp.MustCompile(contentRegexString)
	titleRegex := regexp.MustCompile(titleRegexString)

	contentMatch := contentRegex.FindAllStringSubmatch(bodyString, -1)
	titleMatch := titleRegex.FindStringSubmatch(bodyString)

	title := ""
	if len(titleMatch) > 1 {
		title = titleMatch[1]
	}

	var images []string
	for _, image := range contentMatch {
		images = append(images, image[1])
	}

	return images, title, nil
}

// Download - Requests data url and stores it as content data
// requires the folder path where content will be stored
// requires a slice with content URLs to be downloaded
func (g Getter) Download(folder string, contentURL []string) error {
	_, err := os.Stat(folder)

	folderName := folder
	if folder == "" {
		folderName = defaultFolderName
	}

	fileDir := g.path + folderName + string(os.PathSeparator)

	//Create Directory
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, filePermissions)
		if err != nil {
			return err
		}
	}

	for i := range contentURL {
		response, err := download.ContentBytes(contentURL[i])
		if err != nil {
			return err
		}

		name := getImageName(contentURL[i])

		// Create an empty file
		file, err := os.Create(fileDir + name) //nolint:gosec // received value needs to be a variable
		if err != nil {
			return fmt.Errorf("error creating file: %s", err.Error())
		}

		_, err = file.Write(response)
		if err != nil {
			return fmt.Errorf("error writing file: %s", err.Error())
		}

		// Close stream
		err = file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func getImageName(url string) string {
	splitURL := strings.Split(url, "/")

	length := len(splitURL)

	return splitURL[length-1]
}
