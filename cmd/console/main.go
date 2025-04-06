package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ribeirohugo/go_content_getter/pkg/config"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
)

const (
	configFile    = "config.toml"
	insertMessage = "Insert a new URL to fetch content:"
)

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

	sourceGetter := source.New(cfg.Path, cfg.ContentRegex, cfg.TitleRegex)

	// Create a signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			getContent(sourceGetter)
		}
	}()

	<-done
	fmt.Println("Console closed.")
}

func getContent(src source.Getter) {
	var url string

	fmt.Println(insertMessage)

	_, err := fmt.Scan(&url)
	if err != nil {
		log.Println(err)
	}

	content, err := src.GetAndStore(url)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(content)
}
