package main

import (
	"log"

	"github.com/ribeirohugo/go_content_getter/internal/config"
	"github.com/ribeirohugo/go_content_getter/internal/getter"
	"github.com/ribeirohugo/go_content_getter/internal/server"
)

const configFile = "config.toml"

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	imageGetter := getter.New(cfg)

	serverHTTP := server.New(imageGetter, cfg.Host)

	err = serverHTTP.InitiateServer()
	if err != nil {
		log.Fatal(err)
	}
}
