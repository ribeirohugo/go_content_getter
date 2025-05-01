package source

import (
	"github.com/ribeirohugo/go_content_getter/pkg/download"
	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/page"
	"github.com/ribeirohugo/go_content_getter/pkg/store"
	"github.com/ribeirohugo/go_content_getter/pkg/target"
)

// Get returns slice with all files from a URL.
func (s Getter) Get(url string) ([]model.File, error) {
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

// GetAndStore returns a slice with all files for a given URL string and stores it.
func (s Getter) GetAndStore(url string) ([]model.File, error) {
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

	err = store.All(s.Path, srcPage.Title, files)
	return files, err
}

// GetMany returns a map with slice of all files for a list of URL.
func (s Getter) GetMany(urls []string) (map[string][]model.File, error) {
	urlFiles := make(map[string][]model.File)

	for i := range urls {
		files, err := s.Get(urls[i])
		if err != nil {
			return nil, err
		}

		urlFiles[urls[i]] = files
	}

	return urlFiles, nil
}

// GetAndStoreMany returns a slice with all files for many URL and stores it.
func (s Getter) GetAndStoreMany(urls []string) (map[string][]model.File, error) {
	urlFiles := make(map[string][]model.File)

	for i := range urls {
		files, err := s.GetAndStore(urls[i])
		if err != nil {
			return nil, err
		}

		urlFiles[urls[i]] = files
	}

	return urlFiles, nil
}
