package getter

import "testing"

const (
	regexTest = "[abc]"
	urlTest   = "sub.domain"
)

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
