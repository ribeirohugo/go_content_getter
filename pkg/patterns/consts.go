// Package patterns holds regex patterns to fetch specific content from a page
package patterns

const (
	// HTMLTitle - HTML title between title tags
	HTMLTitle = "(?:\\<title\\>)(.*)(?:<\\/title\\>)"

	// ImageSrc - is HTML image full URL source value
	ImageSrc = "src=[\"'](http[s]?://[a-zA-Z0-9/._-]+(?::[0-9]+)?/[a-zA-Z0-9/._-]*[.](?:jpg|gif|png))(?:[?&#].*)?[\"']"

	// ImageContentFromHrefURL - Href value with HTTP or HTTPS prefix and with jpg, gif or png format
	ImageContentFromHrefURL = "href=[\"'](http[s]?://[a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"

	// ImageContentFromHrefURLWithoutHTTP Href value without any HTTP prefix and with jpg, gif or png format
	ImageContentFromHrefURLWithoutHTTP = "href=[\"']([a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"

	// FileExtensionFromURL gets file extension from URL.
	// It ignores domain TLD, if no file exists.
	FileExtensionFromURL = `(?i)^[a-z][a-z0-9+\-.]*://[^/]+(?:/.*)?/[^/\\.][^/]*\.([A-Za-z0-9]+)(?:[?#].*)?$`

	// FileNameFromURL gets file name, without extension.
	FileNameFromURL = `(?i)^[a-z][a-z0-9+\-.]*://[^/]+(?:/.*)?/([^/.][^/?#]*?)(?:\.[^/?#]+)*(?:[?#].*)?$`

	// FullFilenameFromURL gets complete file name, with the name and extension.
	// Includes filenames starting with a dot (hidden files)
	FullFilenameFromURL = `(?i)^[a-z][a-z0-9+\-.]*://[^/]+(?:/.*)?/([^/?#]+)(?:[?#].*)?$`
)
