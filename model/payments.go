package model

import (
	"math/big"
	"time"
)

type PaymentModel struct {
	ID          int64
	SenderID    int64
	RecepientID int64
	Sum         *big.Int
	CreatedAt   time.Time
}
