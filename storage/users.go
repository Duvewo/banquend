package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Users struct {
	*pgxpool.Pool
}

func (db *Users) Create() {}
func (db *Users) Update() {}
func (db *Users) Search() {}
func (db *Users) Delete() {}
