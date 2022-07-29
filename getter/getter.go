package getter

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	defaultRegex      = "href=[\"'](http[s]?://[a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"
	defaultTitleRegex = "(?:\\<title\\>)(.*)(?:<\\/title\\>)"
)

// Getter
type Getter struct {
	regex string
	url   string
	path  string
}

func New(regex string, url string, path string) Getter {
	regexExpression := defaultRegex
	if regex != "" {
		regexExpression = regex
	}

	return Getter{
		regex: regexExpression,
		url:   url,
		path:  path,
	}
}

// Returns slice with all images URL, page title
// If any error occurs it returns empty
func (g Getter) Get() ([]string, string, error) {
	return g.GetFromURL(g.url)
}

// Requires Url to get content from
// Returns slice with all images URL, page title
// If any error occurs it returns empty
func (g Getter) GetFromURL(url string) ([]string, string, error) {
	response, err := http.Get(url)
	if err != nil {
		return []string{}, "", nil
	}

	if response.StatusCode != http.StatusOK {
		return []string{}, "", nil
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []string{}, "", err
	}
	bodyString := string(bodyBytes)

	imgRegex := g.regex
	if imgRegex == "" {
		imgRegex = defaultRegex
	}

	imageRegex := regexp.MustCompile(imgRegex)
	titleRegex := regexp.MustCompile(defaultTitleRegex)

	imageMatch := imageRegex.FindAllStringSubmatch(bodyString, -1)
	titleMatch := titleRegex.FindStringSubmatch(bodyString)

	title := ""
	if len(titleMatch) > 1 {
		title = titleMatch[1]
	}

	var images []string
	for _, image := range imageMatch {
		images = append(images, image[1])
	}

	return images, title, nil
}

func (g Getter) Download(folder string, images []string) error {
	_, err := os.Stat(folder)

	fileDir := g.path + folder + string(os.PathSeparator)

	//Create Directory
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, 0755)
		if err != nil {
			return err
		}
	}

	for _, image := range images {
		response, err := http.Get(image)
		if err != nil {
			return err
		}

		if response.StatusCode == http.StatusOK {
			name := getImageName(image)

			//Create an empty file
			file, err := os.Create(fileDir + name)
			if err != nil {
				return err
			}
			defer file.Close()

			//Write file content
			_, err = io.Copy(file, response.Body)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getImageName(url string) string {
	splitURL := strings.Split(url, "/")

	length := len(splitURL)

	return splitURL[length-1]
}
