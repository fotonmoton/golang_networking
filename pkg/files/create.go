package files

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func CreateFile(createIntent File) error {

	if createIntent.Name == "" || createIntent.Content == "" {
		return errors.New("Name and Content should be present")
	}

	cwd, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("can't get cwd: %w", err)
	}

	fileName := cwd + "/" + createIntent.Name

	if fileExists(fileName) {
		return errors.New("file already exists")
	}

	// state
	err = os.WriteFile(fileName, []byte(createIntent.Content), fs.ModePerm)

	if err != nil {
		return fmt.Errorf("can't write to file: %w", err)
	}

	return nil
}
