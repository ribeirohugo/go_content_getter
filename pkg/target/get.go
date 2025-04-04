package target

import (
	"regexp"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

func GetAll(page model.Page, pattern string) ([]model.Target, error) {
	var (
		targets []model.Target
	)

	re := regexp.MustCompile(pattern)
	matches := re.FindSubmatch(page.Content)

	for i := range matches {
		targets = append(targets, model.Target{
			URL: string(matches[i]),
		})
	}

	return targets, nil
}
