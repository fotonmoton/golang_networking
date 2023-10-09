package files

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func UpdateFile(name string, content []byte) error {

	cwd, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("can't get cwd: %w", err)
	}

	fileName := cwd + "/" + name

	if !fileExists(fileName) {
		return errors.New("file doesn't exists")
	}

	err = os.WriteFile(fileName, content, fs.ModeAppend)

	if err != nil {
		return fmt.Errorf("can't write to file: %w", err)
	}

	return nil
}
