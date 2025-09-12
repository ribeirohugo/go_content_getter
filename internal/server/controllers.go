package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ribeirohugo/go_content_getter/internal/file"
	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
	"github.com/ribeirohugo/go_content_getter/pkg/source"
	"github.com/ribeirohugo/go_content_getter/pkg/store"
	urlUtils "github.com/ribeirohugo/go_content_getter/pkg/url"
	"github.com/ribeirohugo/go_content_getter/pkg/video"
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
		var (
			files []model.File
			err   error
		)

		if req.Store {
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

	if req.Store {
		c.JSON(http.StatusOK, ContentResponse{Files: allFiles})
		return
	}

	compressedFiles, err := file.ZipFiles(allFiles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=files.zip")
	c.Data(http.StatusOK, "application/zip", compressedFiles)
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

		downloadedFile, err := download.Target(target)
		if err != nil {
			log.Println(err.Error())

			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		// Only store locally if requested
		if req.Store {
			err = store.File(h.path, "", downloadedFile)
			if err != nil {
				log.Println(err.Error())

				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
				return
			}
		}
		allFiles = append(allFiles, downloadedFile)
	}

	if req.Store {
		c.JSON(http.StatusOK, ContentResponse{Files: allFiles})
		return
	}

	// Compress files into a zip and return as binary so frontend can download it
	compressedFiles, err := file.ZipFiles(allFiles)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=files.zip")
	c.Data(http.StatusOK, "application/zip", compressedFiles)
}

// HealthHandler handles GET /api/health requests
func (h *HttpServer) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// LoadPatternsHandler loads existing default patterns and returns it
func (h *HttpServer) LoadPatternsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, patterns.PatternMap)
}

// GetYoutubeInfoHandler returns video metadata for a given URL.
func (h *HttpServer) GetYoutubeInfoHandler(c *gin.Context) {
	var req VideoInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing url in body"})
		return
	}

	y := video.NewGetter()
	videoInfo, err := y.GetVideoInfo(req.URL)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, YoutubeInfoResponse{Video: videoInfo})
}

// DownloadYoutubeHandler allows to download video or audio.
func (h *HttpServer) DownloadYoutubeHandler(c *gin.Context) {
	var req YoutubeRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing fields in body"})
		return
	}

	videoGetter := video.NewGetter()
	data, err := videoGetter.DownloadYoutubeVideo(req.URL, req.VideoFormat, req.AudioFormat)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	filename := "video.mp4"
	if req.Title != "" {
		title := strings.ReplaceAll(req.Title, "/", "_")
		if title != "" {
			filename = fmt.Sprintf("%s.mp4", title)
		}
	} else if v, err := videoGetter.GetVideoInfo(req.URL); err == nil && v.Title != "" {
		title := strings.ReplaceAll(v.Title, "/", "_")
		filename = fmt.Sprintf("%s.mp4", title)
	}

	if req.Store {
		// store file locally
		f := model.File{Filename: filename, Content: data}
		if err := store.File(h.path, "", f); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, ContentResponse{Files: []model.File{f}})
		return
	}

	files := []model.File{{Filename: filename, Content: data}}

	zipData, err := file.ZipFiles(files)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	// return binary for direct download
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Data(http.StatusOK, "application/octet-stream", zipData)
}

// DownloadVideoHandler allows to download video or audio.
func (h *HttpServer) DownloadVideoHandler(c *gin.Context) {
	var req VideoDownloadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid or missing fields in body"})
		return
	}

	videoGetter := video.NewGetter()
	var files []model.File

	for _, url := range req.URLs {
		var (
			data []byte
			err  error
		)

		if req.Format == "mp3" {
			data, err = videoGetter.DownloadAudio(url, req.AudioQuality)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
				return
			}
		} else {
			data, err = videoGetter.DownloadVideo(url, req.VideoQuality, req.AudioQuality)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
				return
			}
		}

		filename, err := videoGetter.GetTitle(url)
		if err != nil {
			log.Println(err.Error())
		}

		filename = fmt.Sprintf("%s.%s", filename, req.Format)

		newFile := model.File{Filename: filename, Content: data}

		if req.Store {
			// store file locally
			if err := store.File(h.path, "", newFile); err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
				return
			}
		} else {
			files = append(files, newFile)
		}
	}

	if req.Store {
		c.JSON(http.StatusOK, ContentResponse{Files: files})
		return
	}

	zipData, err := file.ZipFiles(files)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	// return binary for direct download
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "files.zip"))
	c.Data(http.StatusOK, "application/octet-stream", zipData)
}
