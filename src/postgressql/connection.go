package postgressql

import (
	"context"
	"fmt"
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

func New(ctx context.Context, maxAttempts int, username, password, host, port, database string) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgressql://%s:%s@%s:%s/%s", username, password, host, port, database)

	doWithAttempts(func() error {
		ctx, cf := context.WithTimeout(ctx, 5*time.Second)
		defer cf()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, 5, 5*time.Second)

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
