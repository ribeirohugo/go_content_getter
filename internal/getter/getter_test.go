package getter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	regexTest = "[abc]"
	urlTest   = "sub.domain"
)

var pageRequest = `<title>Page Title</title>

<a href="https://sub.domain/image.png">Image</a>
`

func TestNewGetter(t *testing.T) {
	expected := Getter{
		regex: regexTest,
		url:   urlTest,
	}

	result := New(urlTest, regexTest)

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
		// Send response to be tested
		rw.Write([]byte(pageRequest))
	}))
	// Close the server when test finishes
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
