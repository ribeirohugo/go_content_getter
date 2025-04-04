package source

import (
	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/page"
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
	"github.com/ribeirohugo/go_content_getter/pkg/target"
)

const (
	defaultContentRegex = patterns.ImageContentFromHrefURL
	defaultTitleRegex   = patterns.HTMLTitle
)

// Source holds configurations data and methods
type Source struct {
	ContentRegex string `toml:"contentRegex"`
	TitleRegex   string `toml:"titleRegex"`
	Host         string `toml:"host"`
	Path         string `toml:"path"`
	URL          string `toml:"url"`
}

// New is a Getter constructor. It requires:
// A url string from a web page to look for content.
// A path string to define where to store fetched content. (Optional field)
// A contentRegex to select to download. (Optional field)
// A titleRegex to select folder title to fetched content. (Optional field)
func New(url string, path string, contentRegex string, titleRegex string) Source {
	contentRegexExpression := defaultContentRegex
	if contentRegex != "" {
		contentRegexExpression = contentRegex
	}

	titleRegexExpression := defaultTitleRegex
	if titleRegex != "" {
		titleRegexExpression = titleRegex
	}

	return Source{
		ContentRegex: contentRegexExpression,
		TitleRegex:   titleRegexExpression,
		Path:         path,
		URL:          url,
	}
}
