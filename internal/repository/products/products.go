package products

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	productsDomain "go_template_project/internal/domain/products"
)

func (r *Repository) GetProducts(
	ctx context.Context,
	data productsDomain.GetProductsDTO,
) ([]productsDomain.Product, error) {
	params := SqGetProductsParams{
		Limit:  uint64(data.Limit),
		Offset: uint64(data.Offset),
		Name:   data.Name,
		Title:  data.Title,
	}
	sqProducts, err := r.queries.SqGetProducts(ctx, params)
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

func (r *Repository) GetProduct(
	ctx context.Context,
	data productsDomain.GetProductDTO,
) (*productsDomain.Product, error) {
	params := SqGetProductParams{
		ID: pgtype.UUID{Bytes: data.ID, Valid: true},
	}
	sqProduct, err := r.queries.SqGetProduct(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, productsDomain.ErrProductNotFound
		}
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

func (r *Repository) CreateProduct(
	ctx context.Context,
	data productsDomain.CreateProductDTO,
) (*productsDomain.Product, error) {
	params := SqCreateProductParams{
		Name:  data.Name,
		Title: data.Title,
	}

	sqProduct, err := r.queries.SqCreateProduct(ctx, params)
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

func (r *Repository) PartialUpdateProduct(
	ctx context.Context,
	data productsDomain.PartialUpdateProductDTO,
) (*productsDomain.Product, error) {
	params := SqPartialUpdateProductParams{
		ID:    pgtype.UUID{Bytes: data.ID, Valid: true},
		Name:  data.Name,
		Title: data.Title,
	}

	sqProduct, err := r.queries.SqPartialUpdateProduct(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, productsDomain.ErrProductNotFound
		}
		return nil, fmt.Errorf("sq partial update product error: %w", err)
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

func (r *Repository) DeleteProduct(
	ctx context.Context,
	data productsDomain.DeleteProductDTO,
) (*productsDomain.Product, error) {
	params := SqDeleteProductParams{
		ID: pgtype.UUID{Bytes: data.ID, Valid: true},
	}

	sqProduct, err := r.queries.SqDeleteProduct(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, productsDomain.ErrProductNotFound
		}
		return nil, fmt.Errorf("sq delete product error: %w", err)
	}

	request := productsDomain.Product{
		ID: sqProduct.ID.Bytes,
	}

	return &request, nil
}

func (r *Repository) BulkCreateProducts(
	ctx context.Context,
	data []productsDomain.Product,
) ([]productsDomain.Product, error) {
	params := make([]SqProductRow, 0)
	for _, product := range data {
		params = append(params, SqProductRow{
			Name:  product.Name,
			Title: product.Title,
		})
	}
	sqProducts, err := r.queries.SqBulkCreateProducts(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("sq bulk create products error: %w", err)
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

func (r *Repository) BulkUpdateProducts(
	ctx context.Context,
	updateFields []string,
	data []productsDomain.Product,
) ([]productsDomain.Product, error) {
	params := SqBulkUpdateProductsParams{
		UpdateFields: updateFields,
	}
	for _, product := range data {
		params.Products = append(params.Products, SqProductRow{
			ID:    pgtype.UUID{Bytes: product.ID, Valid: true},
			Title: product.Title,
			Name:  product.Name,
		})
	}

	sqProducts, err := r.queries.SqBulkUpdateProducts(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to build bulk update query: %w", err)
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
