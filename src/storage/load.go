package storage

import (
	"io"
	"os"
)

// load file from disk
func Load(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read file
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
