package products

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) GetProducts(
	ctx context.Context,
	data productsDomain.GetProductsDTO,
) ([]productsDomain.Product, error) {
	products, err := h.repository.GetProducts(ctx, data)
	if err != nil {
		log.Println(err)
		return products, err
	}
	return products, nil
}
