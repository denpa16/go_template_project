package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"go_template_project/internal/domain"
)

func (r *Repository) GetProducts(
	ctx context.Context,
	data domain.GetProductsDTO,
) ([]domain.Product, error) {
	params := GetProductsParams{
		Limit:  uint64(data.Limit),
		Offset: uint64(data.Offset),
		Name:   data.Name,
		Title:  data.Title,
	}
	query, args, err := buildCreateRequestQuery(params)
	if err != nil {
		return nil, fmt.Errorf("sq get products build query error: %w", err)
	}
	sqProducts, err := r.queries.SqGetProducts(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("sq get products error: %w", err)
	}

	products := make([]domain.Product, 0)

	for _, sqProduct := range sqProducts {
		products = append(products, domain.Product{
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

func buildCreateRequestQuery(
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
