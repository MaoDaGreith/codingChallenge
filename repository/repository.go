package repository

import (
	"codingChallenge/model"
	"encoding/json"
	"fmt"
)

func SearchID(request model.UserIDGetter) (model.User, error) {
	var users map[int]model.User
	var user model.User

	content, err := dataReader("./repository/data/users.json")
	if err != nil {
		return user, err
	}

	var data []model.User
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return user, err
	}

	var userRecords = make(map[int]model.User)
	for _, val := range data {
		userRecords[val.ID] = val
	}

	return users[request.ID], err
}

func ActionsCount(request model.ActionCountGetter) (model.Count, error) {
	var count model.Count

	content, err := dataReader("./repository/data/actions.json")
	if err != nil {
		return count, err
	}

	var data []model.Action
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return count, err
	}

	for _, val := range data {
		if val.UserID == request.ID {
			count.Count++
		}
	}
	return count, nil
}
