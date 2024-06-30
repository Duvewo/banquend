package models

import "time"

type MessageModel struct {
	ID           int64
	SenderID     int64
	RecepientIDs []int64 // Possibility to create a group
	Body         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
