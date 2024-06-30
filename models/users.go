package models

import "time"

type UserModel struct {
	ID          int64
	PublicID    string
	FirstName   string
	LastName    string
	Email       string
	AccountType int8 // 0 - user, 1 - support, 2 - admin
	PhoneNumber string
	Nationality string
	Password    string
	CreatedAt   time.Time
}

func (m UserModel) Validate() {

}
