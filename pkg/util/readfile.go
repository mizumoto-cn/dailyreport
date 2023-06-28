package util

import (
	"bufio"
	"log"
	"os"
	"time"
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
		log.Printf("%s read line: %s", time.Now().Format("2006-01-02 15:04:05"), line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}
