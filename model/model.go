package model

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u User) Validate() error {
	if u.ID == 0 {
		return errors.New("User ID should not be empty")
	}
	return nil
}

type UserIDGetter struct {
	ID int `json:"id"`
}
