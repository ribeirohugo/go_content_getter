package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Getter interface {
	Get() ([]string, string, error)
	GetFromUrl(url string) ([]string, string, error)
	Download(folder string, images []string) error
}

type httpServer struct {
	getter Getter
	host   string
	mux    *http.ServeMux
}

func New(getter Getter, host string) *httpServer {
	s := &httpServer{
		getter: getter,
		host:   host,
		mux:    http.NewServeMux(),
	}

	return s
}

func (h *httpServer) InitiateServer() error {
	router := gin.Default()

	router.Static("/assets", "./assets")

	router.LoadHTMLFiles("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home",
		})
	})

	router.POST("/", func(c *gin.Context) {
		c.Request.ParseForm()
		url := c.Request.PostForm["url_parse"][0]

		images, title, err := h.getter.GetFromUrl(url)
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
