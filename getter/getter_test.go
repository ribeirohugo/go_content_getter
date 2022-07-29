package getter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirohugo/go_content_getter/internal/config"
)

const (
	regexTest = "[abc]"
	urlTest   = "sub.domain"
	pathTest  = "path/to/"
)

var pageRequest = `<title>Page Title</title>

<a href="https://sub.domain/image.png">Image</a>
`

func TestNewGetter(t *testing.T) {
	expected := Getter{
		path:  pathTest,
		regex: regexTest,
		url:   urlTest,
	}

	cfg := config.Config{
		Path:  pathTest,
		Regex: regexTest,
		URL:   urlTest,
	}

	result := New(cfg.Regex, cfg.URL, cfg.Path)

	if result != expected {
		t.Errorf("Wrong getter return,\n Got: %v,\n Want: %v.", result, expected)
	}
}

func TestGetImageName(t *testing.T) {
	expected := "image.png"
	result := getImageName("http://sub.domain/image.png")

	if result != expected {
		t.Errorf("Wrong image name return,\n Got: %s,\n Want: %s.", result, expected)
	}
}

func TestGetter_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(pageRequest))
		if err != nil {
			t.Errorf("Unexpected error while writing: %s", err)
		}
	}))
	defer server.Close()

	getter := Getter{
		url: server.URL,
	}

	images, title, err := getter.Get()

	length := len(images)
	expected := 1
	if length != 1 {
		t.Errorf("Wrong images slice length return,\n Got: %d,\n Want: %d.", length, expected)
	}

	expectedTitle := "Page Title"
	if expectedTitle != title {
		t.Errorf("Wrong page title return,\n Got: %s,\n Want: %s.", title, expectedTitle)
	}

	if err != nil {
		t.Errorf("Wrong image name return,\n Got: %v,\n Want: %v.", err, nil)
	}
}
