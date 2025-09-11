package server

import (
	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/video"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ContentResponse represents the response for content endpoints
type ContentResponse struct {
	Title string       `json:"title"`
	Files []model.File `json:"files"`
}

// ContentCompressedResponse represents the downloaded content in compressed mode.
type ContentCompressedResponse struct {
	Title   string `json:"title"`
	ZipFile []byte `json:"zipFile"`
}

// DownloadRequest represents the request body for download content
type DownloadRequest struct {
	URLs           []string `json:"urls"`
	ContentPattern string   `json:"contentPattern"`
	TitlePattern   string   `json:"titlePattern"`
	Store          bool     `json:"store"`
}

// DownloadURLRequest represents the request body for download a URL content
type DownloadURLRequest struct {
	URL   string
	Store bool `json:"store"`
}

// VideoInfoRequest represents the request body for getting youtube video info
type VideoInfoRequest struct {
	URL string `json:"url"`
}

// VideoInfoResponse represents the response containing youtube video metadata
type VideoInfoResponse struct {
	Video video.Video `json:"video"`
}

// VideoDownloadRequest represents the request body for downloading a youtube/video format
type VideoDownloadRequest struct {
	URL         string `json:"url"`
	VideoFormat string `json:"videoFormat"`
	AudioFormat string `json:"audioFormat"`
	Title       string `json:"title"`
	Store       bool   `json:"store"`
}
