package repository

import (
	"codingChallenge/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func dataReader(pathFile string) (map[int]model.User, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
		return nil, err
	}

	var data []model.User // You can also use a struct instead of map
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return nil, err
	}

	var userRecords = make(map[int]model.User)
	for _, val := range data {
		userRecords[val.ID] = val
	}

	return userRecords, nil
}
