// Package server has http server business logic
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirohugo/go_content_getter/pkg/config"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

// Source - content get methods interface
type Source interface {
	Get(string) ([]model.File, error)
	GetAndStore(string) ([]model.File, error)
}

type HttpServer struct {
	host                string
	path                string
	defaultRegexPattern string
	defaultTitlePattern string
	mux                 *http.ServeMux
}

// New - HTTP server constructor
func New(cfg config.Config) *HttpServer {
	s := &HttpServer{
		host:                cfg.Host,
		path:                cfg.Path,
		defaultRegexPattern: cfg.ContentRegex,
		defaultTitlePattern: cfg.TitleRegex,
		mux:                 http.NewServeMux(),
	}

	return s
}

// InitiateServer - Initiates an HTTP server routers and configs and starts the server after that
func (h *HttpServer) InitiateServer() error {
	router := gin.Default()

	router.Static("/assets", "./assets")

	router.LoadHTMLFiles("templates/index.html")

	// API group
	api := router.Group("/api")
	{
		// POST /api/download - download many
		api.POST("/download", h.DownloadManyHandler)
	}

	// Health endpoint
	api.GET("/health", h.HealthHandler)

	err := router.Run(h.host)

	return err
}
