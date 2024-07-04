package storage

import (
	"context"

	"github.com/Duvewo/banquend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Currencies struct {
	*pgxpool.Pool
}

func (db *Currencies) ByCode(ctx context.Context, code string) (models.CurrencyModel, error) {
	const q = "SELECT * FROM currencies WHERE code = $1 LIMIT 1"
	var currency models.CurrencyModel
	return currency, db.QueryRow(ctx, q, code).Scan(&currency)
}
