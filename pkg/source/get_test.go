package source

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(html))

		case filePath:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(fileContent))

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	src := Getter{
		ContentRegex: patterns.ImageSrc,
		TitleRegex:   defaultTitleRegex,
	}
	files, err := src.Get(server.URL)

	require.NoError(t, err)
	require.Len(t, files, 1)
	assert.Equal(t, filename, files[0].Filename)
	assert.Equal(t, fileContent, string(files[0].Content))
}

func TestSource_GetAndStore(t *testing.T) {
	const (
		filename    = "text.png"
		filePath    = "/text.png"
		fileContent = "content"

		subFolder = "test_folder"
	)
	var (
		server *httptest.Server

		tmpDir = t.TempDir()

		html = `<html><head><title>test_folder</title></head><body><img src="%s%s"></body></html>`
	)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/":
			html = fmt.Sprintf(html, server.URL, filePath)

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(html))

		case filePath:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(fileContent))

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	src := Getter{
		ContentRegex: patterns.ImageSrc,
		TitleRegex:   defaultTitleRegex,
		Path:         tmpDir,
	}
	files, err := src.GetAndStore(server.URL)

	require.NoError(t, err)
	require.Len(t, files, 1)
	assert.Equal(t, filename, files[0].Filename)
	assert.Equal(t, fileContent, string(files[0].Content))

	expectedPath := filepath.Join(tmpDir, subFolder, filename)

	data, err := os.ReadFile(expectedPath)
	require.NoError(t, err)
	assert.Equal(t, fileContent, string(data))
}
