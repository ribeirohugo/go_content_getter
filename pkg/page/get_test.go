package page

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

func TestGetHTTP_Success_WithTitle(t *testing.T) {
	var (
		html      = `<html><head><title>Hello World</title></head><body><p>Page Content</p></body></html>`
		pageTitle = "Hello World"
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	page, err := GetHTTP(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, html, string(page.Content))
	assert.Equal(t, pageTitle, page.Title)
}

func TestGetHTTP_Success_NoTitle(t *testing.T) {
	html := `<html><head></head><body><p>No title here</p></body></html>`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	page, err := GetHTTP(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, html, string(page.Content))
	assert.Equal(t, "", page.Title)
}

func TestGetHTTP_HTTPError(t *testing.T) {
	// invalid domain to force http.Get to fail
	badURL := "http://invalid.localhost"

	page, err := GetHTTP(badURL)

	assert.Error(t, err)
	assert.Equal(t, model.Page{}, page)
	assert.Contains(t, err.Error(), "error making HTTP request")
}

func TestGetHTTP_BadStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}))
	defer server.Close()

	page, err := GetHTTP(server.URL)

	assert.Error(t, err)
	assert.Equal(t, model.Page{}, page)
	assert.Contains(t, err.Error(), "error response status: 403 Forbidden")
}
