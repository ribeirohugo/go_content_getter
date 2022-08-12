// Package patterns holds regex patterns to fetch specific content from a page
package patterns

const (
	// HtmlTitle
	HtmlTitle = "(?:\\<title\\>)(.*)(?:<\\/title\\>)"

	// ImageContentFromHrefURL
	ImageContentFromHrefURL = "href=[\"'](http[s]?://[a-zA-Z0-9/._-]+[.](?:jpg|gif|png))[\"']"
)
