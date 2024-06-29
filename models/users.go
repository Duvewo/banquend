package models

import "time"

type UserModel struct {
	ID          int64     `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Nationality string    `json:"nationality"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}
