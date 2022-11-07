package postgressql

import (
	"context"
	"fmt"
	"test-postgres/config"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func New(ctx context.Context, maxAttempts int, cfg config.ConfigPostgres) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	doWithAttempts(func() error {
		ctx, cf := context.WithTimeout(ctx, 5*time.Second)
		defer cf()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, 5, 1*time.Second)

	return
}

func doWithAttempts(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}

	return
}
