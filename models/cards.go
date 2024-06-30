package models

import "time"

type CardModel struct {
	ID           int64
	Status       int8 //0 - issued, 1 - blocked, 2 - frozen, 3 - deleted
	AccountID    int64
	PublicNumber int
	SecurityCode string // hashed
	ValidUntil   time.Time
	CreatedAt    time.Time
}
