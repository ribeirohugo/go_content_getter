package source

import (
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

const (
	defaultContentRegex = patterns.ImageContentFromHrefURL
	defaultTitleRegex   = patterns.HTMLTitle
)

// Getter holds configurations data and methods
type Getter struct {
	ContentRegex string `toml:"contentRegex"`
	TitleRegex   string `toml:"titleRegex"`
	Host         string `toml:"host"`
	Path         string `toml:"path"`
}

// New is a Getter constructor. It requires:
// A url string from a web page to look for content.
// A path string to define where to store fetched content. (Optional field)
// A contentRegex to select to download. (Optional field)
// A titleRegex to select folder title to fetched content. (Optional field)
func New(path, contentRegex, titleRegex string) Getter {
	contentRegexExpression := defaultContentRegex
	if contentRegex != "" {
		contentRegexExpression = contentRegex
	}

	titleRegexExpression := defaultTitleRegex
	if titleRegex != "" {
		titleRegexExpression = titleRegex
	}

	return Getter{
		ContentRegex: contentRegexExpression,
		TitleRegex:   titleRegexExpression,
		Path:         path,
	}
}
