package url

import (
	"regexp"

	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

// GetFileExtension extracts the file extension from a URL string
func GetFileExtension(url string) string {
	regex := regexp.MustCompile(patterns.FileExtensionFromURL)
	matches := regex.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// GetFileName extracts the file name from a URL string
func GetFileName(url string) string {
	regex := regexp.MustCompile(patterns.FileNameFromURL)
	matches := regex.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// GetFullFileName returns the filename including its extension from a URL string
func GetFullFileName(url string) string {
	regex := regexp.MustCompile(patterns.FullFilenameFromURL)
	matches := regex.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
