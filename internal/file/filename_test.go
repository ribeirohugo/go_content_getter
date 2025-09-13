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
	t.Run("general characters", func(t *testing.T) {
		cases := []struct {
			name string
			base string
			ext  string
			want string
		}{
			{
				name: "sanitize invalid chars in base",
				base: "my:cool|video ",
				ext:  "mp4",
				want: "my_cool_video .mp4",
			},
			{
				name: "sanitize invalid chars in extension",
				base: "report",
				ext:  "v1:txt",
				want: "report.v1_txt",
			},
			{
				name: "remove high runes in base",
				base: "😀bad|name",
				ext:  "txt",
				want: "bad_name.txt",
			},
			{
				name: "trim surrounding spaces only (internal space kept)",
				base: "file",
				ext:  " mp4 ",
				want: "file. mp4",
			},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				got := CreateFilename(tc.base, tc.ext)
				assert.Equal(t, tc.want, got)
			})
		}
	})

	t.Run("Latin1 characters filtering", func(t *testing.T) {
		cases := []struct {
			name string
			in   string
			want string
		}{
			{
				name: "remove leading emoji",
				in:   "😀file.txt",
				want: "file.txt",
			},
			{
				name: "remove middle emoji, keep latin1",
				in:   "música🎵.mp3",
				want: "música.mp3",
			},
			{
				name: "only emojis become empty",
				in:   "😀🔥🚀",
				want: "",
			},
			{
				name: "keep accented latin1 chars",
				in:   "áéíóúçñÑÀÊÔÜ",
				want: "áéíóúçñÑÀÊÔÜ",
			},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				got := SanitizeFilename(tc.in)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
