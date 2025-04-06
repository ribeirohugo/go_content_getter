package source

import (
	"fmt"
	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/page"
	"github.com/ribeirohugo/go_content_getter/pkg/store"
	"github.com/ribeirohugo/go_content_getter/pkg/target"
)

// Get returns slice with all files from a URL.
func (s Source) Get(url string) ([]model.File, error) {
	srcPage, err := page.GetHTTP(url)
	if err != nil {
		return []model.File{}, err
	}

	targets, err := target.GetAll(srcPage, s.ContentRegex)
	if err != nil {
		return []model.File{}, err
	}

	files, err := download.ManyTargets(targets)
	if err != nil {
		return []model.File{}, err
	}

	return files, nil
}

// GetAndStore returns slice with all files for a given URL string.
func (s Source) GetAndStore(url string) ([]model.File, error) {
	srcPage, err := page.GetHTTP(url)
	if err != nil {
		return []model.File{}, err
	}

	targets, err := target.GetAll(srcPage, s.ContentRegex)
	if err != nil {
		return []model.File{}, err
	}

	fmt.Println("targets: ", targets)
	fmt.Println(targets)

	files, err := download.ManyTargets(targets)
	if err != nil {
		return []model.File{}, err
	}

	err = store.All(s.Path, srcPage.Title, files)
	return files, err
}
