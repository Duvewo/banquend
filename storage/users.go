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
func (db *Users) ByEmail(ctx context.Context, email string) (user models.UserModel, err error) {
	const q = "SELECT * FROM users WHERE email = $1 LIMIT 1"
	return user, db.QueryRow(ctx, q, email).Scan(&user)
}

func (db *Users) ByPhoneNumber(ctx context.Context, phoneNumber string) (user models.UserModel, err error) {
	const q = "SELECT * FROM users WHERE phone_number = $1 LIMIT 1"
	return user, db.QueryRow(ctx, q, phoneNumber).Scan(&user)
}

func (db *Users) ByPublicID(ctx context.Context, PID string) (user models.UserModel, err error) {
	const q = "SELECT * FROM users WHERE public_id = $1 LIMIT 1"
	return user, db.QueryRow(ctx, q, PID).Scan(&user)
}

func (db *Users) Delete(ctx context.Context, u models.UserModel) error {
	const q = "DELETE FROM"
	return nil
}
