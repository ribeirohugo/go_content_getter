package patterns

import "github.com/ribeirohugo/go_content_getter/pkg/model"

// PatternMap is the pattern storage with description and regex pattern.
var PatternMap = []model.Pattern{
	{
		Description: "HTML title",
		Regex:       HTMLTitle,
	},
	{
		Description: "Image from src attribute",
		Regex:       ImageContentFromHrefURL,
	},
	{
		Description: "Image from Href URL",
		Regex:       ImageContentFromHrefURL,
	},
	{
		Description: "Image from Href URL without HTTP prefix",
		Regex:       ImageContentFromHrefURLWithoutHTTP,
	},
}
