package models

import (
	"math/big"
	"time"
)

type PaymentModel struct {
	ID        int64
	Amount    *big.Float
	CreatedAt time.Time
}
