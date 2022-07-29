package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ribeirohugo/go_content_getter/getter"
	"github.com/ribeirohugo/go_content_getter/internal/config"
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

	downloader := getter.New(cfg.Regex, cfg.URL, cfg.Path)

	// Create a signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			getContent(downloader)
		}
	}()

	<-done
	fmt.Println("Console closed.")
}

func getContent(downloader getter.Getter) {
	var url string

	fmt.Println(insertMessage)

	_, err := fmt.Scan(&url)
	if err != nil {
		log.Println(err)
	}

	images, title, err := downloader.GetFromURL(url)
	if err != nil {
		log.Println(err)
	}

	err = downloader.Download(title, images)
	if err != nil {
		log.Println(err)
	}
}
