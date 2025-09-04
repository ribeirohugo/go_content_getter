package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
)

// DownloadManyHandler handles POST /api/download requests
func (h *HttpServer) DownloadManyHandler(c *gin.Context) {
	var req DownloadRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.URLs) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing urls in body"})
		return
	}

	downloadSource := source.New(h.path, req.ContentPattern, req.TitlePattern)

	var allFiles []model.File
	for _, url := range req.URLs {
		files, err := downloadSource.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		allFiles = append(allFiles, files...)
	}

	c.JSON(http.StatusOK, ContentResponse{Files: allFiles})
}

// DownloadAndStoreManyHandler handles download and store content
func (h *HttpServer) DownloadAndStoreManyHandler(c *gin.Context) {
	var req DownloadRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.URLs) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing urls in body"})
		return
	}

	downloadSource := source.New(h.path, req.ContentPattern, req.TitlePattern)

	var allFiles []model.File
	for _, url := range req.URLs {
		files, err := downloadSource.GetAndStore(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		allFiles = append(allFiles, files...)
	}

	c.JSON(http.StatusOK, ContentResponse{Files: allFiles})
}

// HealthHandler handles GET /api/health requests
func (h *HttpServer) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DefaultPatternsHandler handles GET /api/default-patterns requests
func (h *HttpServer) DefaultPatternsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, DefaultPatternsResponse{
		ContentPattern: h.defaultRegexPattern,
		TitlePattern:   h.defaultTitlePattern,
	})
}
