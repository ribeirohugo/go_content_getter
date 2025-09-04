package main

import (
	"log"

	"github.com/ribeirohugo/go_content_getter/internal/server"
	"github.com/ribeirohugo/go_content_getter/pkg/config"
)

const configFile = "config.toml"

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	serverHTTP := server.New(cfg)

	err = serverHTTP.InitiateServer()
	if err != nil {
		log.Fatal(err)
	}
}
