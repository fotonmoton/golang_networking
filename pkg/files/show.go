package files

import (
	"fmt"
	"os"
)

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func ShowFile(name string) (*File, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("can't get cwd: %w", err)
	}

	filename := cwd + "/" + name

	// state
	file, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("can't read file: %w", err)
	}

	return &File{
		Name:    name,
		Content: string(file),
	}, nil
}
