package models

import "time"

type MessageModel struct {
	ID           int64
	MID          int64
	SenderID     int64
	RecepientIDs []int64 // Possibility to create a group, []UserModel.ID
	Body         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
