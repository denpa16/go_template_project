package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgtype"
	productsDomain "go_template_project/internal/domain/products"
)

func (r *Repository) GetProducts(
	ctx context.Context,
	data productsDomain.GetProductsDTO,
) ([]productsDomain.Product, error) {
	params := GetProductsParams{
		Limit:  uint64(data.Limit),
		Offset: uint64(data.Offset),
		Name:   data.Name,
		Title:  data.Title,
	}
	query, args, err := buildGetProductsQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get products build query error: %w", err)
	}
	sqProducts, err := r.queries.SqGetProducts(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("sq get products error: %w", err)
	}

	products := make([]productsDomain.Product, 0)

	for _, sqProduct := range sqProducts {
		products = append(products, productsDomain.Product{
			ID:        sqProduct.ID.Bytes,
			Name:      sqProduct.Name,
			Title:     sqProduct.Title,
			CreatedAt: sqProduct.CreatedAt.Time,
			UpdatedAt: sqProduct.UpdatedAt.Time,
			DeletedAt: NConvertPgTimestamp(sqProduct.DeletedAt),
		})
	}

	return products, nil
}

func buildGetProductsQuery(
	params GetProductsParams,
) (string, []interface{}, error) {
	dbFields := GetDbFieldsWithValues(params)
	query := sq.Select("id", "name", "title", "created_at", "updated_at", "deleted_at").
		From("products").
		Limit(params.Limit).
		Offset(params.Offset).
		PlaceholderFormat(sq.Dollar)
	query = AddWhereAnd([]string{"name", "title"}, query, dbFields)
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sq get products query to sql error: %w", err)
	}
	fmt.Println(sqlString)
	return sqlString, args, nil
}

func (r *Repository) GetProduct(
	ctx context.Context,
	data productsDomain.GetProductDTO,
) (*productsDomain.Product, error) {
	params := GetProductParams{
		ID: pgtype.UUID{Bytes: data.ID, Valid: true},
	}
	query, args, err := buildGetProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get product build query error: %w", err)
	}
	sqProduct, err := r.queries.SqGetProduct(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("sq get product error: %w", err)
	}
	product := &productsDomain.Product{
		ID:        sqProduct.ID.Bytes,
		Name:      sqProduct.Name,
		Title:     sqProduct.Title,
		CreatedAt: sqProduct.CreatedAt.Time,
		UpdatedAt: sqProduct.UpdatedAt.Time,
		DeletedAt: NConvertPgTimestamp(sqProduct.DeletedAt),
	}
	return product, nil
}

func buildGetProductQuery(
	params GetProductParams,
) (string, []interface{}, error) {
	dbFields := GetDbFieldsWithValues(params)
	query := sq.Select("id", "name", "title", "created_at", "updated_at", "deleted_at").
		From("products").
		PlaceholderFormat(sq.Dollar)
	query = AddWhereAnd([]string{"id"}, query, dbFields)
	sqlString, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sq get product query to sql error: %w", err)
	}
	fmt.Println(sqlString)
	return sqlString, args, nil
}

func (r *Repository) CreateProduct(
	ctx context.Context,
	data productsDomain.CreateProductDTO,
) (*productsDomain.Product, error) {
	params := CreateProductParams{
		Name:  data.Name,
		Title: data.Title,
	}

	query, args, err := buildCreateProductQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq create product build query error: %w", err)
	}
	sqProduct, err := r.queries.SqCreateRequest(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("sq create product error: %w", err)
	}

	request := productsDomain.Product{
		ID:        sqProduct.ID.Bytes,
		Name:      sqProduct.Name,
		Title:     sqProduct.Title,
		CreatedAt: sqProduct.CreatedAt.Time,
		UpdatedAt: sqProduct.UpdatedAt.Time,
		DeletedAt: NConvertPgTimestamp(sqProduct.DeletedAt),
	}

	return &request, nil
}

func buildCreateProductQuery(
	params CreateProductParams,
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
	fmt.Println(sqlString)
	return sqlString, args, nil
}
