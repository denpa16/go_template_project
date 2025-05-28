package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

const CreateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`

type SqGetProductsRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqGetProduct struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqCreateProduct struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
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
) (SqGetProduct, error) {
	row := q.db.QueryRow(ctx, query, args...)
	var i SqGetProduct
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

func (q *Queries) SqCreateRequest(ctx context.Context, query string, args []interface{}) (*SqCreateProduct, error) {
	row := q.db.QueryRow(ctx, query, args...)
	var i SqCreateProduct
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
