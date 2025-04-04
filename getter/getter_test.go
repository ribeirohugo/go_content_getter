package getter

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

const pageRequest = `<title>Page Title</title>

<a href="image.png">Image</a>
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
	t.Run("with no error return", func(t *testing.T) {
		t.Run("with content return", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				_, err := rw.Write([]byte(pageRequest))
				require.NoError(t, err)
			}))
			defer server.Close()

			getter := Getter{
				url:          server.URL,
				contentRegex: patterns.ImageContentFromHrefURLWithoutHTTP,
			}

			images, title, err := getter.Get()
			assert.Len(t, images, 1)
			assert.Equal(t, "Page Title", title)
			assert.NoError(t, err)
		})

		t.Run("without no content return", func(t *testing.T) {
			t.Run("no content found in response body", func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					_, err := rw.Write([]byte(pageRequest))
					require.NoError(t, err)
				}))
				defer server.Close()

				getter := Getter{
					url: server.URL,
				}

				images, title, err := getter.Get()
				assert.Len(t, images, 0)
				assert.Equal(t, "Page Title", title)
				assert.NoError(t, err)
			})

			t.Run("page status code is not ok", func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					rw.WriteHeader(http.StatusBadRequest)
				}))
				defer server.Close()

				getter := Getter{
					url:          server.URL + "/not_found",
					contentRegex: patterns.ImageContentFromHrefURLWithoutHTTP,
				}

				images, title, err := getter.Get()
				assert.Len(t, images, 0)
				assert.Equal(t, "", title)
				assert.NoError(t, err)
			})

			t.Run("invalid url", func(t *testing.T) {
				getter := Getter{
					url: "",
				}

				images, title, err := getter.Get()
				assert.Len(t, images, 0)
				assert.Equal(t, "", title)
				assert.NoError(t, err)
			})
		})
	})

	t.Run("with error return", func(t *testing.T) {
		t.Run("reading body response", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.Header().Set("Content-Length", "1")
			}))
			defer server.Close()

			getter := Getter{
				url: server.URL,
			}

			images, title, err := getter.Get()
			assert.Len(t, images, 0)
			assert.Equal(t, "", title)
			assert.Error(t, err)
		})
	})
}

func TestGetter_Download(t *testing.T) {
	const folderName = "example"

	var getter = Getter{}

	t.Run("with no error return", func(t *testing.T) {
		t.Run("with folder defined", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				_, err := rw.Write([]byte("pageBody"))
				require.NoError(t, err)
			}))
			defer server.Close()

			err := getter.Download(folderName, []string{server.URL})
			assert.NoError(t, err)

			err = os.RemoveAll(folderName)
			require.NoError(t, err)
		})

		t.Run("with empty folder name", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				_, err := rw.Write([]byte("pageBody"))
				require.NoError(t, err)
			}))
			defer server.Close()

			err := getter.Download("", []string{server.URL})
			assert.NoError(t, err)

			err = os.RemoveAll(defaultFolderName)
			require.NoError(t, err)
		})
	})

	t.Run("with error return", func(t *testing.T) {
		t.Run("creating file", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				_, err := rw.Write([]byte("pageBody"))
				require.NoError(t, err)
			}))
			defer server.Close()

			err := getter.Download(string(os.PathSeparator), []string{server.URL})
			assert.Error(t, err)
		})
	})
}
