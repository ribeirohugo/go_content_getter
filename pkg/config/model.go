package config

import (
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

const (
	defaultContentRegex = patterns.ImageSrc
	defaultTitleRegex   = patterns.HTMLTitle
)

// Config holds configurations data and methods
type Config struct {
	ContentRegex string   `toml:"content_regex"`
	TitleRegex   string   `toml:"title_regex"`
	Host         string   `toml:"host"`
	Path         string   `toml:"path"`
	URL          []string `toml:"url"`
}

// New is a Getter constructor. It requires:
// A url string from a web page to look for content.
// A path string to define where to store fetched content. (Optional field)
// A contentRegex to select to download. (Optional field)
// A titleRegex to select folder title to fetched content. (Optional field)
func New(url []string, path, contentRegex, titleRegex string) Config {
	contentRegexExpression := defaultContentRegex
	if contentRegex != "" {
		contentRegexExpression = contentRegex
	}

	titleRegexExpression := defaultTitleRegex
	if titleRegex != "" {
		titleRegexExpression = titleRegex
	}

	return Config{
		ContentRegex: contentRegexExpression,
		TitleRegex:   titleRegexExpression,
		Path:         path,
		URL:          url,
	}
}
