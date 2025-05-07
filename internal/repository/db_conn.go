package repository

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxConn(ctx context.Context, config Config) (*pgxpool.Pool, error) {
	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Username,
		url.QueryEscape(config.Password),
		config.Host,
		config.Port,
		config.Name,
	)

	pool, err := pgxpool.New(ctx, connURL)

	if err != nil {
		return nil, errors.New("invalid credentials for db connection")
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
