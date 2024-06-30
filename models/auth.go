package models

type AuthModel struct {
	ID          int64
	PhoneNumber string
	Password    string
	Type        string // login, register, restore
}
