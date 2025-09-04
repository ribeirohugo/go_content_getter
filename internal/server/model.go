package server

import "github.com/ribeirohugo/go_content_getter/pkg/model"

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ContentResponse represents the response for content endpoints
type ContentResponse struct {
	Files []model.File `json:"files"`
}

// DownloadRequest represents the request body for download content
type DownloadRequest struct {
	URLs           []string `json:"urls"`
	ContentPattern string   `json:"contentPattern"`
	TitlePattern   string   `json:"titlePattern"`
}

// DefaultPatternsResponse represents the response for default regex patterns
type DefaultPatternsResponse struct {
	ContentPattern string `json:"contentPattern"`
	TitlePattern   string `json:"titlePattern"`
}
