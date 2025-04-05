// Package server has http server business logic
package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

// Source - content get methods interface
type Source interface {
	Get() ([]model.File, error)
	GetFromURL(url string) ([]model.File, error)
}

type HttpServer struct {
	source Source
	host   string
	mux    *http.ServeMux
}

// New - HTTP server constructor
func New(source Source, host string) *HttpServer {
	s := &HttpServer{
		source: source,
		host:   host,
		mux:    http.NewServeMux(),
	}

	return s
}

// InitiateServer - Initiates an HTTP server routers and configs and starts the server after that
func (h *HttpServer) InitiateServer() error {
	router := gin.Default()

	router.Static("/assets", "./assets")

	router.LoadHTMLFiles("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home",
		})
	})

	router.POST("/", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
		}

		url := c.Request.PostForm["url_parse"][0]

		files, err := h.source.GetFromURL(url)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"message": err.Error(),
			})

			return
		}

		err = h.getter.Download(title, images)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"message": err.Error(),
			})

			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Home",
			"message": "Success!",
		})
	})

	err := router.Run(h.host)

	return err
}
