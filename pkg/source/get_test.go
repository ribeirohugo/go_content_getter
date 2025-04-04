package source

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

func TestTarget_Success(t *testing.T) {
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

	file, err := download.Target(target)

	assert.NoError(t, err)
	assert.Equal(t, expectedFilename, file.Filename)
	assert.Equal(t, expectedContent, file.Content)
}
