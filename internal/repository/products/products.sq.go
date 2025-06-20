package products

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	ProductsTable = "products"
)

const CreateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const PartialUpdateProductSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const DeleteProductSuffix = `RETURNING id`
const BulkCreateProductsSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`
const BulkUpdateProductsSuffix = `RETURNING id, name, title, created_at, updated_at, deleted_at`

type SqProductRow struct {
	ID        pgtype.UUID
	Name      string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqCreateProductParams struct {
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

type SqPartialUpdateProductParams struct {
	ID    pgtype.UUID
	Name  string `db:"name"`
	Title string `db:"title"`
}

type SqDeleteProductParams struct {
	ID pgtype.UUID `db:"id"`
}

type SqBulkUpdateProductsParams struct {
	UpdateFields []string
	Products     []SqProductRow
}

func (q *RepoQueries) SqGetProducts(
	ctx context.Context,
	params SqGetProductsParams,
) ([]SqProductRow, error) {
	query, args, err := buildGetProductsQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get products build query error: %w", err)
	}
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

func buildGetProductsQuery(
	params SqGetProductsParams,
) (string, []interface{}, error) {
	dbFields := GetDbFieldsWithValues(params)
	query := sq.Select("id", "name", "title", "created_at", "updated_at", "deleted_at").
		From(ProductsTable).
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
) (*SqProductRow, error) {
	query, args, err := buildGetProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get product build query error: %w", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i SqProductRow
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
		From(ProductsTable).
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
) (*SqProductRow, error) {
	query, args, err := buildCreateProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq create product build query error: %w", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i SqProductRow
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

	query := sq.Insert(ProductsTable).
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
) (*SqProductRow, error) {
	query, args, err := buildPartialUpdateProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq partial update product build query error: %w", err)
	}

	row := q.db.QueryRow(ctx, query, args...)
	var i SqProductRow
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
	query := sq.Update(ProductsTable).
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
) (*SqProductRow, error) {
	query, args, err := buildDeleteProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq delete product build query error: %w", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i SqProductRow
	err = row.Scan(
		&i.ID,
	)
	return &i, err
}

func buildDeleteProductQuery(
	params SqDeleteProductParams,
) (string, []interface{}, error) {
	query := sq.Delete(ProductsTable).
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
	params []SqProductRow,
) ([]SqProductRow, error) {
	query, args, err := buildBulkCreateProductsQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq bulk create proudcts build query error: %w", err)
	}
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

func buildBulkCreateProductsQuery(params []SqProductRow) (string, []interface{}, error) {
	columns := []string{"name", "title"}
	query := sq.Insert(ProductsTable).
		Columns(columns...).
		Suffix(BulkCreateProductsSuffix).
		PlaceholderFormat(sq.Dollar)

	for _, product := range params {
		query = query.Values(
			product.ID,
			product.Name,
			product.Title,
		)
	}

	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}

	return sqlString, args, nil
}

func (q *RepoQueries) SqBulkUpdateProducts(
	ctx context.Context,
	params SqBulkUpdateProductsParams,
) ([]SqProductRow, error) {
	query, args, err := buildBulkUpdateProductsQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq bulk update products build query error: %w", err)
	}
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

func buildBulkUpdateProductsQuery(
	params SqBulkUpdateProductsParams,
) (string, []interface{}, error) {
	updateFieldsMap := make(map[string]bool)
	for _, field := range params.UpdateFields {
		updateFieldsMap[field] = true
	}

	query := sq.Update(ProductsTable).
		Suffix(BulkUpdateProductsSuffix).
		PlaceholderFormat(sq.Dollar)

	nameCase := sq.Case()
	titleCase := sq.Case()
	for _, product := range params.Products {
		if _, ok := updateFieldsMap["name"]; ok {
			nameCase = nameCase.When(
				sq.Eq{"id": product.ID.Bytes}, sq.Expr("?", product.Name),
			).Else("name")
		}
		if _, ok := updateFieldsMap["title"]; ok {
			nameCase = nameCase.When(
				sq.Eq{"id": product.ID.Bytes}, sq.Expr("?", product.Title),
			).Else("title")
		}
	}

	query = query.
		Set("name", nameCase).
		Set("title", titleCase)

	sqlString, args, err := query.ToSql()

	if err != nil {
		return "", nil, err
	}
	return sqlString, args, nil
}
