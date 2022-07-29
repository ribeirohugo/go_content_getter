package getter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var pageRequest = `<title>Page Title</title>

<a href="https://sub.domain/image.png">Image</a>
`

func TestNewGetter(t *testing.T) {
	const (
		contentRegexTest = "[abc]"
		titleRegexTest   = "title"
		pathTest         = "path/to/"
		urlTest          = "sub.domain"
	)

	t.Run("with content and title regex filled", func(t *testing.T) {
		expected := Getter{
			contentRegex: contentRegexTest,
			titleRegex:   titleRegexTest,
			path:         pathTest,
			url:          urlTest,
		}

		result := New(urlTest, pathTest, contentRegexTest, titleRegexTest)
		assert.Equal(t, expected, result)
	})

	t.Run("with content regex empty", func(t *testing.T) {
		expected := Getter{
			contentRegex: defaultContentRegex,
			titleRegex:   titleRegexTest,
			path:         pathTest,
			url:          urlTest,
		}

		result := New(urlTest, pathTest, "", titleRegexTest)
		assert.Equal(t, expected, result)
	})

	t.Run("with title regex empty", func(t *testing.T) {
		expected := Getter{
			contentRegex: contentRegexTest,
			titleRegex:   defaultTitleRegex,
			path:         pathTest,
			url:          urlTest,
		}

		result := New(urlTest, pathTest, contentRegexTest, "")
		assert.Equal(t, expected, result)
	})
}

func TestGetImageName(t *testing.T) {
	expected := "image.png"
	result := getImageName("http://sub.domain/image.png")

	assert.Equal(t, expected, result)
}

func TestGetter_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(pageRequest))
		require.NoError(t, err)
	}))
	defer server.Close()

	getter := Getter{
		url: server.URL,
	}

	images, title, err := getter.Get()
	assert.Len(t, images, 1)
	assert.Equal(t, "Page Title", title)
	assert.NoError(t, err)
}
