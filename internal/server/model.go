package server

import "github.com/ribeirohugo/go_content_getter/pkg/model"

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ContentResponse represents the response for content endpoints
type ContentResponse struct {
	Title string       `json:"title"`
	Files []model.File `json:"files"`
}

// DownloadRequest represents the request body for download content
type DownloadRequest struct {
	URLs           []string `json:"urls"`
	ContentPattern string   `json:"contentPattern"`
	TitlePattern   string   `json:"titlePattern"`
}

// DownloadURLRequest represents the request body for download a URL content
type DownloadURLRequest struct {
	URL string
}
