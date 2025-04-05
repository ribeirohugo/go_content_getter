package source

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

func TestSource_Get(t *testing.T) {
	const (
		filename    = "text.png"
		filePath    = "/text.png"
		fileContent = "content"
	)
	var (
		server *httptest.Server

		html = `<html><head><title>Hello World</title></head><body><img src="%s%s"></body></html>`
	)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/":
			html = fmt.Sprintf(html, server.URL, filePath)
			fmt.Println(html)

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(html))

		case filePath:
			fmt.Println("chegou")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(fileContent))

		default:
			fmt.Println()
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	src := Source{
		URL:          server.URL,
		ContentRegex: patterns.ImageSrc,
		TitleRegex:   defaultTitleRegex,
	}
	files, err := src.Get()

	require.NoError(t, err)
	require.Len(t, files, 1)
	assert.Equal(t, filename, files[0].Filename)
	assert.Equal(t, fileContent, string(files[0].Content))
}

//
//func TestTarget_Success(t *testing.T) {
//
//	expectedContent := []byte("file content")
//	expectedFilename := "test.txt"
//
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		_, _ = w.Write(expectedContent)
//	}))
//	defer server.Close()
//
//	target := model.Target{
//		URL:      server.URL,
//		Filename: expectedFilename,
//	}
//
//	file, err := download.Target(target)
//
//	assert.NoError(t, err)
//	assert.Equal(t, expectedFilename, file.Filename)
//	assert.Equal(t, expectedContent, file.Content)
//}
//
//func TestSource_Get(t *testing.T) {
//	var (
//		html      = `<html><head><title>Hello World</title></head><body><p>Page Content</p></body></html>`
//		pageTitle = "Hello World"
//	)
//
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		_, _ = w.Write([]byte(html))
//	}))
//	defer server.Close()
//
//	page, err := GetHTTP(server.URL)
//
//	assert.NoError(t, err)
//	assert.Equal(t, html, string(page.Content))
//	assert.Equal(t, pageTitle, page.Title)
//}
