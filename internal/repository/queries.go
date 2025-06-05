package repository

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
)

func (r *Repository) GetProducts(ctx context.Context, data productsDomain.GetProductsDTO) ([]productsDomain.Product, error) {
	return r.productsRepo.GetProducts(ctx, data)
}

func (r *Repository) GetProduct(ctx context.Context, data productsDomain.GetProductDTO) (*productsDomain.Product, error) {
	return r.productsRepo.GetProduct(ctx, data)
}

func (r *Repository) CreateProduct(ctx context.Context, data productsDomain.CreateProductDTO) (*productsDomain.Product, error) {
	return r.productsRepo.CreateProduct(ctx, data)
}

func (r *Repository) PartialUpdateProduct(
	ctx context.Context,
	data productsDomain.PartialUpdateProductDTO,
) (*productsDomain.Product, error) {
	return r.productsRepo.PartialUpdateProduct(ctx, data)
}

func (r *Repository) DeleteProduct(
	ctx context.Context,
	data productsDomain.DeleteProductDTO,
) (*productsDomain.Product, error) {
	return r.productsRepo.DeleteProduct(ctx, data)
}

func (r *Repository) BulkCreateProducts(
	ctx context.Context,
	data []productsDomain.Product,
) (int64, error) {
	return r.productsRepo.BulkCreateProducts(ctx, data)
}

func (r *Repository) BulkUpdateProducts(
	ctx context.Context,
	data []productsDomain.Product,
) ([]productsDomain.Product, error) {
	return r.productsRepo.BulkUpdateProducts(ctx, data)
}
