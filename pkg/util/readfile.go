package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) ([]string, error) {
	// readlines add <p> to the beginning of the line and </p> to the end of the line
	lines := []string{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = "<p>" + line + "</p>"
		fmt.Println(line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}
