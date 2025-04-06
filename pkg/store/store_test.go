package store

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

func TestFile(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		tmpDir := t.TempDir()
		subFolder := "test_folder"

		testFile := model.File{
			Filename: "hello.txt",
			Content:  []byte("Hello, world!"),
		}

		err := File(tmpDir, subFolder, testFile)
		assert.NoError(t, err)

		expectedPath := filepath.Join(tmpDir, subFolder, testFile.Filename)

		data, err := os.ReadFile(expectedPath)
		assert.NoError(t, err)
		assert.Equal(t, testFile.Content, data)
	})

	t.Run("with errors", func(t *testing.T) {
		t.Run("of invalid path", func(t *testing.T) {
			// simulate error by using an invalid path (on most systems)
			invalidPath := string([]byte{0x00}) // null byte is invalid in file paths
			file := model.File{Filename: "fail.txt", Content: []byte("nope")}

			err := File(invalidPath, "fail", file)
			assert.Error(t, err)
		})
	})
}

func TestAll(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		tmpDir := t.TempDir()
		subFolder := "batch_folder"

		files := []model.File{
			{Filename: "file1.txt", Content: []byte("First file")},
			{Filename: "file2.txt", Content: []byte("Second file")},
		}

		err := All(tmpDir, subFolder, files)
		assert.NoError(t, err)

		for _, f := range files {
			expectedPath := filepath.Join(tmpDir, subFolder, f.Filename)
			data, err := os.ReadFile(expectedPath)
			assert.NoError(t, err)
			assert.Equal(t, f.Content, data)
		}
	})

	t.Run("with errors", func(t *testing.T) {
		t.Run("with invalid path", func(t *testing.T) {
			// simulate error by using an invalid path (on most systems)
			invalidPath := string([]byte{0x00}) // null byte is invalid in file paths

			files := []model.File{
				{Filename: "cant-write.txt", Content: []byte("will fail")},
			}

			err := All(invalidPath, "fail", files)
			assert.Error(t, err)
		})

		t.Run("of existing subfolder", func(t *testing.T) {
			tmpDir := t.TempDir()
			subFolder := "readonly"

			// create the subfolder manually and make it read-only
			fullPath := filepath.Join(tmpDir, subFolder)
			err := os.MkdirAll(fullPath, 0444)
			assert.NoError(t, err)

			files := []model.File{
				{Filename: "cant-write.txt", Content: []byte("will fail")},
			}

			err = All(tmpDir, subFolder, files)
			assert.Error(t, err)
		})
	})
}
