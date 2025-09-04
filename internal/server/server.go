// Package server has http server business logic
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirohugo/go_middlewares/pkg/cors"

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

	allowedOrigins []string

	mux *http.ServeMux
}

// New - HTTP server constructor
func New(cfg config.Config) *HttpServer {
	s := &HttpServer{
		host:                cfg.Host,
		path:                cfg.Path,
		defaultRegexPattern: cfg.ContentRegex,
		defaultTitlePattern: cfg.TitleRegex,
		allowedOrigins:      cfg.AllowedOrigins,
		mux:                 http.NewServeMux(),
	}

	return s
}

// InitiateServer - Initiates an HTTP server routers and configs and starts the server after that
func (h *HttpServer) InitiateServer() error {
	router := gin.Default()

	// Middleware
	router.Use(corsMiddleware(h.allowedOrigins)) // Enables CORS using the custom middleware

	// Static
	// router.Static("/assets", "./assets")
	// router.LoadHTMLFiles("templates/index.html")

	// API group
	api := router.Group("/api")
	{
		// Download and Store
		api.POST("/download", h.DownloadManyHandler)
		api.POST("/download-and-store", h.DownloadAndStoreManyHandler)

		// Health endpoint
		api.GET("/health", h.HealthHandler)

		// Default patterns endpoint
		api.GET("/default-patterns", h.DefaultPatternsHandler)
	}

	err := router.Run(h.host)

	return err
}

func corsMiddleware(allowedOrigins []string) gin.HandlerFunc {
	corsMiddleware := cors.New(allowedOrigins)

	return func(c *gin.Context) {
		final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})
		corsMiddleware.Middleware(final).ServeHTTP(c.Writer, c.Request)
	}
}
