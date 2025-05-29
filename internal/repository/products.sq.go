package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const CreateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const PartialUpdateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const DeleteProductSuffix = `RETURNING id`

type SqGetProductsRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqGetProductRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqCreateProductRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqPartialUpdateProductRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqDeleteProductRow struct {
	ID pgtype.UUID
}

type GetProductsParams struct {
	Limit  uint64
	Offset uint64
	Name   string `db:"name"`
	Title  string `db:"title"`
}

type GetProductParams struct {
	ID pgtype.UUID `db:"id"`
}

type CreateProductParams struct {
	Name  string `db:"name"`
	Title string `db:"title"`
}

type PartialUpdateProductParams struct {
	ID    pgtype.UUID
	Name  string `db:"name"`
	Title string `db:"title"`
}

type DeleteProductParams struct {
	ID pgtype.UUID
}

func (q *Queries) SqGetProducts(
	ctx context.Context,
	query string,
	args []interface{},
) ([]SqGetProductsRow, error) {
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SqGetProductsRow
	for rows.Next() {
		var i SqGetProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *Queries) SqGetProduct(
	ctx context.Context,
	query string,
	args []interface{},
) (SqGetProductRow, error) {
	row := q.db.QueryRow(ctx, query, args...)
	var i SqGetProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

func (q *Queries) SqCreateProduct(ctx context.Context, query string, args []interface{}) (*SqCreateProductRow, error) {
	row := q.db.QueryRow(ctx, query, args...)
	var i SqCreateProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

func (q *Queries) SqPartialUpdateProduct(ctx context.Context, query string, args []interface{}) (*SqPartialUpdateProductRow, error) {
	row := q.db.QueryRow(ctx, query, args...)
	var i SqPartialUpdateProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

func (q *Queries) SqDeleteProduct(ctx context.Context, query string, args []interface{}) (*SqDeleteProductRow, error) {
	row := q.db.QueryRow(ctx, query, args...)
	var i SqDeleteProductRow
	err := row.Scan(
		&i.ID,
	)
	return &i, err
}

func (q *Queries) SqBulkCreateProducts(
	ctx context.Context,
	columns []string,
	sqProducts [][]interface{},
) (int64, error) {
	count, err := q.db.CopyFrom(
		ctx,
		pgx.Identifier{"products"},
		columns,
		pgx.CopyFromRows(sqProducts),
	)
	if err != nil {
		return 0, fmt.Errorf("query failed: %w", err)
	}
	return count, nil
}
