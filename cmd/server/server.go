package main

import (
	"log"

	"github.com/ribeirohugo/go_content_getter/internal/server"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
)

const configFile = "config.toml"

func main() {
	src, err := source.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	serverHTTP := server.New(src, src.Host)

	err = serverHTTP.InitiateServer()
	if err != nil {
		log.Fatal(err)
	}
}
