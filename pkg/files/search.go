package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(name string, search string) ([]string, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("can't get cwd: %w", err)
	}

	filename := cwd + "/" + name

	// state
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	number := 1

	// representation
	lines := []string{}

	for scanner.Scan() {

		text := scanner.Text()

		if strings.Contains(text, search) {
			lines = append(lines, fmt.Sprintf("%d\t %s", number, text))
		}

		number++
	}

	return lines, nil
}
