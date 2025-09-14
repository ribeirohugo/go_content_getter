package target

import (
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

// GetAll allows to extract all targets for a given model.Page and a regex pattern string.
func GetAll(page model.Page, pattern string) ([]model.Target, error) {
	var targets []model.Target

	contentRegex := regexp.MustCompile(pattern)
	matches := contentRegex.FindAllStringSubmatch(string(page.Content), -1)

	for _, match := range matches {
		targets = append(targets, model.Target{
			URL:      match[1],
			Filename: extractFilename(match[1]),
		})
	}

	return targets, nil
}

func extractFilename(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		parts := strings.Split(rawURL, "/")
		return parts[len(parts)-1]
	}
	return path.Base(parsed.Path)
}
