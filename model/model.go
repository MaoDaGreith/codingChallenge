package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserIDGetter struct {
	ID int `json:"id"`
}

type Action struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	UserID     int       `json:"userId"`
	TargetUser int       `json:"targetUser"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ActionCountGetter struct {
	ID int `json:"id"`
}

type Count struct {
	Count int `json:"count"`
}

type ActionType struct {
	Type string `json:"type"`
}

type GroupedActions struct {
	UserID  int      `json:"userId"`
	Actions []Action `json:"actions"`
}
