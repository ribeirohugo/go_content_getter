package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic filenames
		{"https://teste.com/teste.html?var=dsdds", "html"},
		{"https://teste.com/teste.html#section", "html"},
		{"https://teste.com/path/file.txt", "txt"},
		{"https://teste.com/file.php?x=1#y", "php"},

		// IPs
		{"https://172.16.0.1/file.htm?x=1#y", "htm"},
		{"https://192.168.1.1/style.css?x=1#y", "css"},

		// Edge cases
		{"https://teste.com/path/", ""},            // no file
		{"https://teste.com", ""},                  // should return
		{"https://teste.com/.hiddenfile", ""},      // hidden file without ext
		{"https://teste.com/image.jpeg", "jpeg"},   // normal case
		{"file://localhost/etc/hosts", ""},         // file without ext
		{"ftp://example.com/archive.tar.gz", "gz"}, // multiple dots
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := GetFileExtension(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

// TestGetFullFileName verifies extraction of the filename including the extension
func TestGetFullFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic filenames
		{"https://teste.com/teste.html?var=dsdds", "teste.html"},
		{"https://teste.com/teste.html#section", "teste.html"},
		{"https://teste.com/path/file.txt", "file.txt"},
		{"https://teste.com/file.php?x=1#y", "file.php"},

		// IPs
		{"https://172.16.0.1/file.htm?x=1#y", "file.htm"},
		{"https://192.168.1.1/style.css?x=1#y", "style.css"},

		// Edge cases
		{"https://teste.com/path/", ""},                        // no file
		{"https://teste.com", ""},                              // no path
		{"https://teste.com/.hiddenfile", ".hiddenfile"},       // hidden file without ext
		{"https://teste.com/image.jpeg", "image.jpeg"},         // normal case
		{"file://localhost/etc/hosts", "hosts"},                // file without ext -> hosts
		{"ftp://example.com/archive.tar.gz", "archive.tar.gz"}, // multiple dots
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := GetFullFileName(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}
