package main

import (
	"log"

	"github.com/ribeirohugo/go_content_getter/getter"
	"github.com/ribeirohugo/go_content_getter/internal/config"
)

const (
	configFile = "config.toml"
)

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	downloader := getter.New(cfg.URL, cfg.Path, cfg.Regex, "")

	images, title, err := downloader.Get()
	if err != nil {
		log.Println(err)
	}

	err = downloader.Download(title, images)
	if err != nil {
		log.Println(err)
	}
}
