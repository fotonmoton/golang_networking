package files

import (
	"fmt"
	"os"
)

type DirEntry struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

func ListFiles(path string) ([]DirEntry, error) {

	// state
	files, err := os.ReadDir(path)

	if err != nil {
		return nil, fmt.Errorf("can't open dir: %w", err)
	}

	entries := make([]DirEntry, 0, len(files))
	for _, entry := range files {
		entries = append(entries, DirEntry{entry.Name(), entry.IsDir()})
	}

	return entries, nil
}
