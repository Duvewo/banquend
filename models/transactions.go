package models

import "time"

type TransactionModel struct {
	ID          int64
	Total       int64
	Payments    []PaymentModel
	SenderID    int64
	RecepientID int64
	CreatedAt   time.Time
}
