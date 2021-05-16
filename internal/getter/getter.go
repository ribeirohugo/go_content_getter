package getter

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/ribeirohugo/go_image_getter/internal/config"
)

const (
	generalRegex     = "href=[\"'](http[s]?://[a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"
	titleRegexString = "<title>(.+?)</title>"
)

type Getter struct {
	regex string
	url   string
}

func New(cfg config.Config) Getter {
	return Getter{
		regex: cfg.Regex,
		url:   cfg.Url,
	}
}

// Returns slice with all images URL, page title
// If any error occurs it returns empty
func (g *Getter) Get() ([]string, string, error) {
	response, err := http.Get(g.url)
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
		imgRegex = generalRegex
	}

	imageRegex := regexp.MustCompile(imgRegex)
	titleRegex := regexp.MustCompile(titleRegexString)

	imageMatch := imageRegex.FindAllStringSubmatch(bodyString, -1)
	titleMatch := titleRegex.FindStringSubmatch(bodyString)

	title := ""
	if len(titleMatch) > 0 {
		title = strings.TrimLeft(strings.TrimRight(titleMatch[0], "</title>"), "<title>")
	}

	var images []string
	for _, image := range imageMatch {
		images = append(images, image[1])
	}

	return images, title, nil
}

func (g *Getter) Download(folder string, images []string) error {
	_, err := os.Stat(folder)

	if os.IsNotExist(err) {
		err := os.MkdirAll(folder, 0755)
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
			file, err := os.Create(folder + "/" + name)
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
	splitUrl := strings.Split(url, "/")

	length := len(splitUrl)

	return splitUrl[length-1]
}
