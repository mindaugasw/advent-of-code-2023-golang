package internal

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadLines(fileName string) (lines []string, err error) {
	_, taskFile, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("could not get task file name")
	}

	taskDir := filepath.Dir(taskFile)
	fullPath := taskDir + "/data/" + fileName

	file, err := os.Open(fullPath)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line != "" {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	if err = file.Close(); err != nil {
		return
	}

	return lines, nil
}
