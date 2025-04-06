package main

import (
	"log"

	"github.com/ribeirohugo/go_content_getter/internal/server"
	"github.com/ribeirohugo/go_content_getter/pkg/config"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
)

const configFile = "config.toml"

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	src := source.New(cfg.Path, cfg.ContentRegex, cfg.TitleRegex)

	serverHTTP := server.New(src, src.Host)

	err = serverHTTP.InitiateServer()
	if err != nil {
		log.Fatal(err)
	}
}
