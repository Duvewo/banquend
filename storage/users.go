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
func (db *Users) Search(ctx context.Context, u models.UserModel) (user models.UserModel, err error) {
	const q = "SELECT * FROM users WHERE email = $1 OR phone_number = $2 OR password= $3 LIMIT 1"
	return user, db.QueryRow(ctx, q, u.Email, u.PhoneNumber, u.Password).Scan(&user)
}
func (db *Users) Delete(ctx context.Context, u models.UserModel) error {
	const q = "DELETE FROM"
	return nil
}
