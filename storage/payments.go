package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Payments struct {
	*pgxpool.Pool
}
