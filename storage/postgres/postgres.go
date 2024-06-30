package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Open(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connString)

	if err != nil {
		return nil, fmt.Errorf("storage/open: to parse config: %w", err)
	}

	return OpenWithConfig(ctx, cfg)
}

func OpenWithConfig(ctx context.Context, cfg *pgxpool.Config) (*pgxpool.Pool, error) {
	return pgxpool.NewWithConfig(ctx, cfg)
}
