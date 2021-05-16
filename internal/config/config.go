package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Regex string `toml:"regex"`
	Url   string `toml:"url"`
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

	var config Config

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
