package models

import "time"

type CardModel struct {
	ID           int64
	Status       int8 //0 - issued, 1 - blocked, 2 - frozen, 3 - deleted
	OwnerID      int64
	PublicNumber int16
	SecurityCode string // hashed
	ValidUntil   time.Time
	CreatedAt    time.Time
}
