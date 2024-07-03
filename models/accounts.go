package models

import (
	"math/big"
	"time"
)

type AccountModel struct {
	ID        int64
	OwnerID   int64
	Currency  int8       // From table Currencies
	Amount    *big.Float // TODO: Find better type for money
	CreatedAt time.Time
	DeletedAt time.Time
}
