package database

import (
	"context"
	"fmt"
	"github.com/gookit/slog"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kstsm/wb-level-0/consumer/config"
)

func InitPostgres(ctx context.Context) *pgxpool.Pool {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	slog.Infof(
		"Connecting to the database... host=%s port=%s db=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		slog.Error("Failed to connect to the database", "error", err)
		panic(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		slog.Error("Database connection check failed", "error", err)
		panic(err)
	}

	return pool
}
