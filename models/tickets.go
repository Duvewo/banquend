package models

import "time"

type TicketModel struct {
	ID        int64
	OwnerID   int64
	Category  int8
	Body      string
	CreatedAt time.Time
}
