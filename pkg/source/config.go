// Package source holds configuration data struct and data loading
package source

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

const defaultHost = "localhost:8080"

// Source holds configurations data and methods
type Source struct {
	ContentRegex string `toml:"contentRegex"`
	TitleRegex   string `toml:"titleRegex"`
	Host         string `toml:"host"`
	Path         string `toml:"path"`
	URL          string `toml:"url"`
}

// Load - loads configurations from a given toml file path
func Load(filePath string) (Source, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return Source{}, err
	}

	config := Source{
		Host:         defaultHost,
		ContentRegex: patterns.ImageContentFromHrefURL,
		TitleRegex:   patterns.HTMLTitle,
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Source{}, err
	}

	return config, nil
}
