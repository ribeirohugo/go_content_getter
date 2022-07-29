package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

const defaultHost = "localhost:8080"

type Config struct {
	ContentRegex string `toml:"contentRegex"`
	TitleRegex   string `toml:"titleRegex"`
	Host         string `toml:"host"`
	Path         string `toml:"path"`
	URL          string `toml:"url"`
}

func Load(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	_ = file.Close()

	config := Config{
		Host: "localhost:8080",
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
