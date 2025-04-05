package source

import (
	"fmt"
	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/page"
	"github.com/ribeirohugo/go_content_getter/pkg/target"
)

// Get returns slice with all images URL, page title
// If any error occurs it returns empty
func (s Source) Get() ([]model.File, error) {
	srcPage, err := page.GetHTTP(s.URL)
	if err != nil {
		return []model.File{}, err
	}

	fmt.Println(srcPage)

	targets, err := target.GetAll(srcPage, s.ContentRegex)
	if err != nil {
		return []model.File{}, err
	}

	fmt.Println(targets)

	files, err := download.ManyTargets(targets)
	if err != nil {
		return []model.File{}, err
	}

	return files, nil
}
