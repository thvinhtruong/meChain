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

// save file to disk
func Save(path string, filename string) error {
	// create file if file not exist
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	// write file
	_, err = file.Write([]byte(filename))
	if err != nil {
		return err
	}

	return nil
}
