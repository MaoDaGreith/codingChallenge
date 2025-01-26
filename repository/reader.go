package repository

import (
	"fmt"
	"io"
	"os"
)

func dataReader(pathFile string) ([]byte, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
		return nil, err
	}

	return content, nil
}
