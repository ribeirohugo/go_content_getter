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
