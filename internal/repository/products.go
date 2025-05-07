package repository

import (
	"context"
	"fmt"
	"go_template_project/internal/domain"
)

func (r *Repository) GetProducts(
	ctx context.Context,
	limit, offset int,
) ([]domain.Product, error) {
	// Products list
	sqlcProducts, err := r.queries.SqlcGetProducts(
		ctx,
		SqlcGetProductsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("sqlc get products error: %w", err)
	}

	var products []domain.Product

	for _, sqlcProduct := range sqlcProducts {
		products = append(products, domain.Product{
			ID:        sqlcProduct.ID,
			Name:      sqlcProduct.Name,
			CreatedAt: sqlcProduct.CreatedAt.Time,
			UpdatedAt: sqlcProduct.UpdatedAt.Time,
			DeletedAt: NConvertPgTimestamp(sqlcProduct.DeletedAt),
		})
	}

	return products, nil
}
