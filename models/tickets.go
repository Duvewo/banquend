package models

import "time"

type TicketModel struct {
	ID        int64
	OwnerID   int64
	Category  string
	Body      string
	CreatedAt time.Time
}
