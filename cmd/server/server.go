package main

import (
	"log"

	"github.com/ribeirohugo/go_image_getter/internal/config"
	"github.com/ribeirohugo/go_image_getter/internal/getter"
	"github.com/ribeirohugo/go_image_getter/internal/server"
)

const configFile = "config.toml"

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	imageGetter := getter.New(cfg)

	serverHttp := server.New(imageGetter, cfg.Host)

	err = serverHttp.InitiateServer()
	if err != nil {
		log.Fatal(err)
	}
}
