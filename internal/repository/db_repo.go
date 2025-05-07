package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go_template_project/internal/domain/ports"
)

type Repository struct {
	conn    Connect
	queries Queries
}

type Connect interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewRepo(conn Connect) *Repository {
	return &Repository{
		conn:    conn,
		queries: *New(conn),
	}
}

func (r *Repository) WithTx(tx pgx.Tx) *Repository {
	return &Repository{
		conn:    tx,
		queries: *New(tx),
	}
}

func (r *Repository) InTx(ctx context.Context, f func(repo ports.Transaction) error) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}

	qtx := r.WithTx(tx)

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	err = f(qtx)
	if err != nil {
		return err
	}

	return nil
}
