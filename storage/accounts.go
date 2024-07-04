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
	const q = "INSERT INTO accounts (id, owner_id, currency, amount, created_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(ctx, q, account.ID, account.OwnerID, account.Currency, account.Amount, account.CreatedAt, account.DeletedAt)
	return err
}

func (db *Accounts) ByOwnerID(ctx context.Context, ownerID int) ([]models.AccountModel, error) {
	return nil, nil
}
func (db *Accounts) ByID(ctx context.Context, ID int) (models.AccountModel, error) {
	return models.AccountModel{}, nil
}
