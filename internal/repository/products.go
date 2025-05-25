package repository

import (
	"context"
	"fmt"
	"go_template_project/internal/domain"
)

func (r *Repository) GetProducts(
	ctx context.Context,
	limit, offset int64,
) ([]domain.Product, error) {
	// Products list
	sqProducts, err := r.queries.SqGetProducts(
		ctx,
		SqGetProductsParams{
			Limit:  limit,
			Offset: offset,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("sq get products error: %w", err)
	}

	products := make([]domain.Product, 0)

	for _, sqProduct := range sqProducts {
		products = append(products, domain.Product{
			ID:        sqProduct.ID,
			Name:      sqProduct.Name,
			CreatedAt: sqProduct.CreatedAt.Time,
			UpdatedAt: sqProduct.UpdatedAt.Time,
			DeletedAt: NConvertPgTimestamp(sqProduct.DeletedAt),
		})
	}

	return products, nil
}
