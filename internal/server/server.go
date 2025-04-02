// Package server has http server business logic
package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Getter - content get methods interface
type Getter interface {
	Get() ([]string, string, error)
	GetFromURL(url string) ([]string, string, error)
	Download(folder string, images []string) error
}

type HttpServer struct {
	getter Getter
	host   string
	mux    *http.ServeMux
}

// New - HTTP server constructor
func New(getter Getter, host string) *HttpServer {
	s := &HttpServer{
		getter: getter,
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

		images, title, err := h.getter.GetFromURL(url)
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
