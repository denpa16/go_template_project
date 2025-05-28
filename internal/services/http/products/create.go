package products

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) CreateProduct(
	ctx context.Context,
	data productsDomain.CreateProductDTO,
) (*productsDomain.Product, error) {
	product, err := h.repository.CreateProduct(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}
