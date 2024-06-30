package models

import (
	"math/big"
	"time"
)

type AccountModel struct {
	ID        int64
	OwnerID   int64
	Currency  string     // USD, EUR, BTC, ETH
	Amount    *big.Float //
	CreatedAt time.Time
	DeletedAt time.Time
}
