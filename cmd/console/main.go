package main

import (
	"log"

	"github.com/ribeirohugo/go_image_getter/internal/config"
	"github.com/ribeirohugo/go_image_getter/internal/getter"
)

const (
	configFile = "config.toml"
)

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	downloader := getter.New(cfg.Url, cfg.Regex)

	images, title, err := downloader.Get()
	if err != nil {
		log.Println(err)
	}

	err = downloader.Download(title, images)
	if err != nil {
		log.Println(err)
	}
}
