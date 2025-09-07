package file

import (
	"archive/zip"
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

// ZipFiles receives multiple files (each with Filename and Content) and
// returns a zip archive as a byte slice containing all provided files.
func ZipFiles(files []model.File) ([]byte, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	for i, f := range files {
		// sanitize filename to avoid directory traversal
		name := filepath.Base(f.Filename)
		if name == "" {
			name = fmt.Sprintf("file-%d", i)
		}

		hdr := &zip.FileHeader{
			Name:   name,
			Method: zip.Deflate,
		}

		hdr.Modified = time.Now()

		w, err := zw.CreateHeader(hdr)
		if err != nil {
			zw.Close()
			return nil, err
		}

		if len(f.Content) > 0 {
			if _, err := w.Write(f.Content); err != nil {
				zw.Close()
				return nil, err
			}
		}
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
