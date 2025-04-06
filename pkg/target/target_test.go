package target

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
)

func TestGetAll(t *testing.T) {
	html := []byte(`
		<html>
			<body>
				<img src="https://example.com/images/photo1.jpg?size=large" />
				<img src="https://cdn.site.org/assets/pic2.png" />
				<img src="/static/image3.gif" />
			</body>
		</html>
	`)

	page := model.Page{
		Content: html,
	}

	expected := []model.Target{
		{
			URL:      "https://example.com/images/photo1.jpg",
			Filename: "photo1.jpg",
		},
		{
			URL:      "https://cdn.site.org/assets/pic2.png",
			Filename: "pic2.png",
		},
	}

	targets, err := GetAll(page, patterns.ImageSrc)

	assert.NoError(t, err)
	assert.Equal(t, expected, targets)
}

func TestGetAll_NoMatch(t *testing.T) {
	page := model.Page{
		Content: []byte(`<html><body><p>No images here</p></body></html>`),
	}

	targets, err := GetAll(page, patterns.ImageSrc)

	assert.NoError(t, err)
	assert.Empty(t, targets)
}
