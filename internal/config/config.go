// Package config holds configuration data struct and data loading
package config

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

const defaultHost = "localhost:8080"

// Config holds configurations data and methods
type Config struct {
	ContentRegex string `toml:"contentRegex"`
	TitleRegex   string `toml:"titleRegex"`
	Host         string `toml:"host"`
	Path         string `toml:"path"`
	URL          string `toml:"url"`
}

// Load - loads configurations from a given toml file path
func Load(filePath string) (Config, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Host:         defaultHost,
		ContentRegex: patterns.ImageContentFromHrefURL,
		TitleRegex:   patterns.HTMLTitle,
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
