package file

import (
	"fmt"
	"regexp"
	"strings"
)

// SanitizeFilename allows to exclude invalid OS path characters and trims filename spaces.
func SanitizeFilename(name string) string {
	// Define invalid characters for most filesystems: \ / : * ? " < > |
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	safe := re.ReplaceAllString(name, "_")

	// Trim spaces at start/end
	safe = strings.TrimSpace(safe)

	return safe
}

// CreateFilename uses filename and extension and finally sanitizes, returning final filename.
func CreateFilename(filename, extension string) string {
	filename = fmt.Sprintf("%s.%s", filename, extension)

	return filename
}
