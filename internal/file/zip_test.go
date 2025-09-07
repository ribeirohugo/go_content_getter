package file

import (
	"archive/zip"
	"bytes"
	"io"
	"testing"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestZipFiles(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		b, err := ZipFiles([]model.File{})
		assert.NoError(t, err)
		// empty zip should still be a valid zip archive (may contain just the EOCD)
		r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		assert.NoError(t, err)
		assert.Len(t, r.File, 0)
	})

	t.Run("single file", func(t *testing.T) {
		files := []model.File{{Filename: "hello.txt", Content: []byte("hello world")}}
		b, err := ZipFiles(files)
		assert.NoError(t, err)

		r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		assert.NoError(t, err)
		assert.Len(t, r.File, 1)

		f := r.File[0]
		assert.Equal(t, "hello.txt", f.Name)

		rc, err := f.Open()
		assert.NoError(t, err)
		defer rc.Close()
		content, err := io.ReadAll(rc)
		assert.NoError(t, err)
		assert.Equal(t, []byte("hello world"), content)
	})

	t.Run("multiple files and sanitization", func(t *testing.T) {
		files := []model.File{
			{Filename: "path/to/image.png", Content: []byte{0x01, 0x02}},
			{Filename: "", Content: []byte("generated")},
		}
		b, err := ZipFiles(files)
		assert.NoError(t, err)

		r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		assert.NoError(t, err)
		assert.Len(t, r.File, 2)

		// first file should be sanitized to image.png
		f0 := r.File[0]
		assert.Equal(t, "image.png", f0.Name)
		rc0, err := f0.Open()
		assert.NoError(t, err)
		data0, err := io.ReadAll(rc0)
		rc0.Close()
		assert.NoError(t, err)
		assert.Equal(t, []byte{0x01, 0x02}, data0)

		// second file should be generated with file-1 name
		f1 := r.File[1]
		assert.Equal(t, "file-1", f1.Name)
		rc1, err := f1.Open()
		assert.NoError(t, err)
		data1, err := io.ReadAll(rc1)
		rc1.Close()
		assert.NoError(t, err)
		assert.Equal(t, []byte("generated"), data1)
	})

	t.Run("empty content file", func(t *testing.T) {
		files := []model.File{{Filename: "empty.txt", Content: []byte{}}}
		b, err := ZipFiles(files)
		assert.NoError(t, err)

		r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		assert.NoError(t, err)
		assert.Len(t, r.File, 1)

		f := r.File[0]
		assert.Equal(t, "empty.txt", f.Name)
		rc, err := f.Open()
		assert.NoError(t, err)
		defer rc.Close()
		content, err := io.ReadAll(rc)
		assert.NoError(t, err)
		assert.Len(t, content, 0)
	})
}
