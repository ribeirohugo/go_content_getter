package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/store"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
	urlUtils "github.com/ribeirohugo/go_content_getter/pkg/url"
)

// DownloadManyHandler handles POST /api/download requests
func (h *HttpServer) DownloadManyHandler(c *gin.Context) {
	h.handleDownload(c, false)
}

// DownloadAndStoreManyHandler handles download and store content
func (h *HttpServer) DownloadAndStoreManyHandler(c *gin.Context) {
	h.handleDownload(c, true)
}

// DownloadURLsHandler downloads content from one or many URLs.
func (h *HttpServer) DownloadURLsHandler(c *gin.Context) {
	var req DownloadRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.URLs) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing urls in body"})
		return
	}

	var allFiles []model.File
	for i, url := range req.URLs {
		filename := urlUtils.GetFullFileName(url)
		if filename == "" {
			filename = fmt.Sprintf("Untitled(%d)", i)
		}

		target := model.Target{
			URL:      url,
			Filename: filename,
		}

		file, err := download.Target(target)
		if err != nil {
			log.Println(err.Error())

			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		err = store.File(h.path, "", file)
		if err != nil {
			log.Println(err.Error())

			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		allFiles = append(allFiles, file)
	}

	c.JSON(http.StatusOK, ContentResponse{Files: allFiles})
}

// HealthHandler handles GET /api/health requests
func (h *HttpServer) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// LoadPatternsHandler loads existing default patterns and returns it
func (h *HttpServer) LoadPatternsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, patterns.PatternMap)
}

// handleDownload is responsible for download process used by controllers
func (h *HttpServer) handleDownload(c *gin.Context, useStore bool) {
	var req DownloadRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.URLs) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing urls in body"})
		return
	}

	downloadSource := source.New(h.path, req.ContentPattern, req.TitlePattern)

	var allFiles []model.File
	for _, url := range req.URLs {
		var (
			files []model.File
			err   error
		)

		if useStore {
			files, err = downloadSource.GetAndStore(url)
		} else {
			files, err = downloadSource.Get(url)
		}

		if err != nil {
			log.Println(err.Error())

			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		allFiles = append(allFiles, files...)
	}

	c.JSON(http.StatusOK, ContentResponse{Files: allFiles})
}
