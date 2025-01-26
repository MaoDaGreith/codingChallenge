package repository

import "codingChallenge/model"

func SearchID(request model.UserIDGetter) (model.User, error) {
	var users map[int]model.User
	var user model.User
	users, err := dataReader("./repository/data/users.json")
	if err != nil {
		return user, err
	}

	return users[request.ID], err
}
