package repository

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
)

type (
	ProductsRepository interface {
		GetProducts(
			ctx context.Context,
			data productsDomain.GetProductsDTO,
		) ([]productsDomain.Product, error)
		GetProduct(
			ctx context.Context,
			data productsDomain.GetProductDTO,
		) (*productsDomain.Product, error)
		CreateProduct(
			ctx context.Context,
			data productsDomain.CreateProductDTO,
		) (*productsDomain.Product, error)
		PartialUpdateProduct(
			ctx context.Context,
			data productsDomain.PartialUpdateProductDTO,
		) (*productsDomain.Product, error)
		DeleteProduct(
			ctx context.Context,
			data productsDomain.DeleteProductDTO,
		) (*productsDomain.Product, error)
		BulkCreateProducts(
			ctx context.Context,
			data []productsDomain.Product,
		) ([]productsDomain.Product, error)
		BulkUpdateProducts(
			ctx context.Context,
			updateFields []string,
			data []productsDomain.Product,
		) ([]productsDomain.Product, error)
	}
)
