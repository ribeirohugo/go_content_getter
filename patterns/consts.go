// Package patterns holds regex patterns to fetch specific content from a page
package patterns

const (
	// HtmlTitle - HTML title between title tags
	HtmlTitle = "(?:\\<title\\>)(.*)(?:<\\/title\\>)"

	// ImageContentFromHrefURL - Href value with HTTP or HTTPS prefix and with jpg, gif or png format
	ImageContentFromHrefURL = "href=[\"'](http[s]?://[a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"

	// ImageContentFromHrefURLWithoutHTTP Href value without any HTTP prefix and with jpg, gif or png format
	ImageContentFromHrefURLWithoutHTTP = "href=[\"']([a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"
)
