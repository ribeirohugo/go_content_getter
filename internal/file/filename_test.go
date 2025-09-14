package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeFilename(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no invalid chars stays same",
			input:    "simple_name.txt",
			expected: "simple_name.txt",
		},
		{
			name:     "replace invalid punctuation",
			input:    "My:Video*Title?.mp4",
			expected: "My_Video_Title_.mp4",
		},
		{
			name:     "trim surrounding spaces",
			input:    "  spaced name  ",
			expected: "spaced name",
		},
		{
			name:     "mixed invalid and angle brackets",
			input:    "in<va>lid|name",
			expected: "in_va_lid_name",
		},
		{
			name:     "slashes and backslashes",
			input:    "path\\to/file",
			expected: "path_to_file",
		},
		{
			name:     "latin1 accented letters preserved",
			input:    "música:ç?",
			expected: "música_ç_",
		},
		{
			name:     "multiple consecutive invalid",
			input:    "??bad",
			expected: "__bad",
		},
	}

	for _, tc := range cases {
		// capture range variable
		c := tc
		t.Run(c.name, func(t *testing.T) {
			got := SanitizeFilename(c.input)
			assert.Equal(t, c.expected, got)
		})
	}
}

func TestCreateFilename(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		ext      string
		expected string
	}{
		{
			name:     "basic join",
			filename: "file",
			ext:      "txt",
			expected: "file.txt",
		},
		{
			name:     "empty extension",
			filename: "file",
			ext:      "",
			expected: "file.",
		},
		{
			name:     "empty filename",
			filename: "",
			ext:      "txt",
			expected: ".txt",
		},
		{
			name:     "both empty",
			filename: "",
			ext:      "",
			expected: ".",
		},
		{
			name:     "duplicate extension not removed",
			filename: "video.mp4",
			ext:      "mp3",
			expected: "video.mp4.mp3",
		},
		{
			name:     "extension starting with dot causes double dot",
			filename: "video",
			ext:      ".mp4",
			expected: "video..mp4",
		},
		{
			name:     "no sanitization performed inside CreateFilename",
			filename: "my:file",
			ext:      "mp*4",
			expected: "my:file.mp*4",
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			got := CreateFilename(c.filename, c.ext)
			assert.Equal(t, c.expected, got)
		})
	}
}
