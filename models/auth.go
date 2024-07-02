package models

type AuthModel struct {
	ID          int64  `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password,omitempty"`
	Type        string `json:"type,omitempty"` // login, register, restore
}
