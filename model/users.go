package model

import "time"

type UserModel struct {
	ID          int64 `json:"id" db:"id"`
	FirstName   string
	LastName    string
	Email       string
	Nationality string
	Password    string
	CreatedAt   time.Time
}
