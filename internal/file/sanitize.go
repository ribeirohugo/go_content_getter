package file

import (
	"regexp"
	"strings"
)

func SanitizeFilename(name string) string {
	// Define invalid characters for most filesystems: \ / : * ? " < > |
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	safe := re.ReplaceAllString(name, "_")

	// Trim spaces at start/end
	safe = strings.TrimSpace(safe)

	// Keep only runes representable in ISO-8859-1 (Latin-1) while preserving them
	var b strings.Builder
	for _, r := range safe {
		if r <= 0xFF { // Latin-1 range
			b.WriteRune(r)
		}
	}
	return b.String()
}
