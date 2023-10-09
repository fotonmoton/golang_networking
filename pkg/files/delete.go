package files

import (
	"errors"
	"fmt"
	"os"
)

func DeleteFile(name string) error {

	cwd, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("can't get cwd: %w", err)
	}

	fileName := cwd + "/" + name

	if !fileExists(fileName) {
		return errors.New("file doesn't exists")
	}

	err = os.Remove(fileName)

	if err != nil {
		return fmt.Errorf("can't delete file: %w", err)
	}

	return nil
}
