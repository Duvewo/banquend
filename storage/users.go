package storage

import (
	"context"

	"github.com/Duvewo/banquend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	*pgxpool.Pool
}

func (db *Users) Create(ctx context.Context, u models.UserModel) error {
	const q = "INSERT INTO users (id, first_name, last_name) VALUES ($1, $2, $3)"
	db.Exec(ctx, q, u.ID, u.FirstName, u.LastName)
	return nil
}

func (db *Users) Update(ctx context.Context, u models.UserModel) error {
	return nil
}
func (db *Users) Search(ctx context.Context, u models.UserModel) (models.UserModel, error) {
	return models.UserModel{}, nil
}
func (db *Users) Delete(ctx context.Context, u models.UserModel) error {
	return nil
}
