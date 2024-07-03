package storage

import (
	"context"

	"github.com/Duvewo/banquend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Accounts struct {
	*pgxpool.Pool
}

func (db *Accounts) Create(ctx context.Context, account models.AccountModel) error {
	return nil
}
