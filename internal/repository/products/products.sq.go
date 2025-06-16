package products

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const CreateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const PartialUpdateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const DeleteProductSuffix = `RETURNING id`
const BulkUpdateProductsSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`

type SqProductRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

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

type SqBulkUpdateProductRow struct {
	ID    pgtype.UUID
	Name  string
	Title string
}

type SqGetProductsParams struct {
	Limit  uint64
	Offset uint64
	Name   string `db:"name"`
	Title  string `db:"title"`
}

type SqGetProductParams struct {
	ID pgtype.UUID `db:"id"`
}

type SqCreateProductParams struct {
	Name  string `db:"name"`
	Title string `db:"title"`
}

type SqPartialUpdateProductParams struct {
	ID    pgtype.UUID
	Name  string `db:"name"`
	Title string `db:"title"`
}

type SqDeleteProductParams struct {
	ID pgtype.UUID `db:"id"`
}

type SqBulkUpdateProductsParams struct {
	ID    pgtype.UUID
	Name  string `db:"name"`
	Title string `db:"title"`
}

func (q *RepoQueries) SqGetProducts(
	ctx context.Context,
	params SqGetProductsParams,
) ([]SqGetProductsRow, error) {
	query, args, err := buildGetProductsQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get products build query error: %w", err)
	}
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

func buildGetProductsQuery(
	params SqGetProductsParams,
) (string, []interface{}, error) {
	dbFields := GetDbFieldsWithValues(params)
	query := sq.Select("id", "name", "title", "created_at", "updated_at", "deleted_at").
		From("products").
		Limit(params.Limit).
		Offset(params.Offset).
		PlaceholderFormat(sq.Dollar)
	query = SelectBuilderAddWhereAnd([]string{"name", "title"}, query, dbFields)
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sq get products query to sql error: %w", err)
	}
	return sqlString, args, nil
}

func (q *RepoQueries) SqGetProduct(
	ctx context.Context,
	params SqGetProductParams,
) (*SqGetProductRow, error) {
	query, args, err := buildGetProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get product build query error: %w", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i SqGetProductRow
	err = row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

func buildGetProductQuery(
	params SqGetProductParams,
) (string, []interface{}, error) {
	dbFields := GetDbFieldsWithValues(params)
	query := sq.Select("id", "name", "title", "created_at", "updated_at", "deleted_at").
		From("products").
		PlaceholderFormat(sq.Dollar)
	query = SelectBuilderAddWhereAnd([]string{"id"}, query, dbFields)
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sq get product query to sql error: %w", err)
	}
	return sqlString, args, nil
}

func (q *RepoQueries) SqCreateProduct(
	ctx context.Context,
	params SqCreateProductParams,
) (*SqCreateProductRow, error) {
	query, args, err := buildCreateProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq create product build query error: %w", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i SqCreateProductRow
	err = row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

func buildCreateProductQuery(
	params SqCreateProductParams,
) (string, []interface{}, error) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	dbFields := GetDbFieldsWithValues(params)
	for k, v := range dbFields {
		columns = append(columns, k)
		values = append(values, v)
	}

	query := sq.Insert("products").
		Columns(columns...).
		Values(values...).
		Suffix(CreateProductSuffix).
		PlaceholderFormat(sq.Dollar)
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	return sqlString, args, nil
}

func (q *RepoQueries) SqPartialUpdateProduct(
	ctx context.Context,
	params SqPartialUpdateProductParams,
) (*SqPartialUpdateProductRow, error) {
	query, args, err := buildPartialUpdateProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq partial update product build query error: %w", err)
	}

	row := q.db.QueryRow(ctx, query, args...)
	var i SqPartialUpdateProductRow
	err = row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

func buildPartialUpdateProductQuery(
	params SqPartialUpdateProductParams,
) (string, []interface{}, error) {
	dbFields := GetDbFieldsWithValues(params)
	query := sq.Update("products").
		SetMap(dbFields).
		Suffix(PartialUpdateProductSuffix).
		PlaceholderFormat(sq.Dollar)
	query = UpdateBuilderAddWhereAnd([]string{"id"}, query, map[string]interface{}{"id": params.ID})
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	fmt.Println(sqlString)
	return sqlString, args, nil
}

func (q *RepoQueries) SqDeleteProduct(
	ctx context.Context,
	params SqDeleteProductParams,
) (*SqDeleteProductRow, error) {
	query, args, err := buildDeleteProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq delete product build query error: %w", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i SqDeleteProductRow
	err = row.Scan(
		&i.ID,
	)
	return &i, err
}

func buildDeleteProductQuery(
	params SqDeleteProductParams,
) (string, []interface{}, error) {
	query := sq.Delete("products").
		Suffix(DeleteProductSuffix).
		PlaceholderFormat(sq.Dollar)
	query = DeleteBuilderAddWhereAnd([]string{"id"}, query, map[string]interface{}{"id": params.ID})
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	return sqlString, args, nil
}

func (q *RepoQueries) SqBulkCreateProducts(
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

func (q *RepoQueries) SqBulkUpdateProducts(
	ctx context.Context,
	query string,
	args []interface{},
) ([]SqProductRow, error) {
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SqProductRow
	for rows.Next() {
		var i SqProductRow
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
