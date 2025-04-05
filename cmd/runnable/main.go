package main

import (
	"fmt"
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

	src := source.New(cfg.URL, cfg.Path, cfg.ContentRegex, cfg.TitleRegex)

	files, err := src.Get()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(files)
}
