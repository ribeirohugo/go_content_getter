package store

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ribeirohugo/go_content_getter/pkg/model"
)

const (
	filePermissions = 0o750
)

func File(path, folder string, file model.File) error {
	fileDir := filepath.Join(path, folder)

	_, err := os.Stat(fileDir)

	// Create Directory
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, filePermissions)
		if err != nil {
			return err
		}
	}

	err = storeFiles(fileDir, file)
	return err
}

func All(path, folder string, files []model.File) error {
	fileDir := filepath.Join(path, folder)

	_, err := os.Stat(fileDir)

	// Create Directory
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, filePermissions)
		if err != nil {
			return err
		}
	}

	for i := range files {
		err := storeFiles(fileDir, files[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func storeFiles(directory string, modelFile model.File) error {
	fileDir := filepath.Join(directory, modelFile.Filename)

	// Create an empty file
	file, err := os.Create(fileDir)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err.Error())
	}

	_, err = file.Write(modelFile.Content)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}

	// Close stream
	err = file.Close()
	if err != nil {
		return fmt.Errorf("error closing stream: %s", err.Error())
	}

	return nil
}
