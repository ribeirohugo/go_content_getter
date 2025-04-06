package main

import (
	"log"

	"github.com/ribeirohugo/go_content_getter/pkg/config"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
)

const (
	configFile = "config.toml"
)

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	src := source.New(cfg.Path, cfg.ContentRegex, cfg.TitleRegex)

	for i := range cfg.URL {
		_, err := src.GetAndStore(cfg.URL[i])
		if err != nil {
			log.Println(err)
		}
	}
}
