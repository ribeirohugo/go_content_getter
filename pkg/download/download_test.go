package download

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

func TestTarget(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedContent := []byte("file content")
		expectedFilename := "test.txt"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(expectedContent)
		}))
		defer server.Close()

		target := model.Target{
			URL:      server.URL,
			Filename: expectedFilename,
		}

		file, err := Target(target)

		assert.NoError(t, err)
		assert.Equal(t, expectedFilename, file.Filename)
		assert.Equal(t, expectedContent, file.Content)
	})

	t.Run("HTTP error", func(t *testing.T) {
		badURL := "http://invalid.localhost"

		target := model.Target{
			URL:      badURL,
			Filename: "bad.txt",
		}

		_, err := Target(target)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error making HTTP request")
	})

	t.Run("status not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "not found", http.StatusNotFound)
		}))
		defer server.Close()

		target := model.Target{
			URL:      server.URL,
			Filename: "404.txt",
		}

		file, err := Target(target)

		require.NoError(t, err)
		assert.Equal(t, model.File{}, file)
	})

	t.Run("status not ok", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "not found", http.StatusInternalServerError)
		}))
		defer server.Close()

		target := model.Target{
			URL:      server.URL,
			Filename: "404.txt",
		}

		_, err := Target(target)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error response status: 500 Internal Server Error")
	})
}

func TestManyTargets(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		contents := []string{"file1", "file2"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/file1" {
				_, _ = io.WriteString(w, contents[0])
			} else if r.URL.Path == "/file2" {
				_, _ = io.WriteString(w, contents[1])
			} else {
				http.NotFound(w, r)
			}
		}))
		defer server.Close()

		targets := []model.Target{
			{URL: server.URL + "/file1", Filename: "file1.txt"},
			{URL: server.URL + "/file2", Filename: "file2.txt"},
		}

		files, err := ManyTargets(targets)

		assert.NoError(t, err)
		assert.Len(t, files, 2)
		assert.Equal(t, "file1.txt", files[0].Filename)
		assert.Equal(t, []byte("file1"), files[0].Content)
		assert.Equal(t, "file2.txt", files[1].Filename)
		assert.Equal(t, []byte("file2"), files[1].Content)
	})

	t.Run("fails", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "fail", http.StatusInternalServerError)
		}))
		defer server.Close()

		targets := []model.Target{
			{URL: server.URL, Filename: "bad.txt"},
			{URL: server.URL, Filename: "also-bad.txt"},
		}

		files, err := ManyTargets(targets)

		assert.Error(t, err)
		assert.Len(t, files, 0)
		assert.Contains(t, err.Error(), "error response status")
	})
}
