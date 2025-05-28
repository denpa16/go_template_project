package products

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
)

type repository interface {
	CreateProduct(
		ctx context.Context,
		data productsDomain.CreateProductDTO,
	) (*productsDomain.Product, error)
	GetProducts(
		ctx context.Context,
		data productsDomain.GetProductsDTO,
	) ([]productsDomain.Product, error)
	GetProduct(
		ctx context.Context,
		data productsDomain.GetProductDTO,
	) (*productsDomain.Product, error)
}
