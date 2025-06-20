package products

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) BulkCreateProducts(
	ctx context.Context,
	data []productsDomain.Product,
) ([]productsDomain.Product, error) {
	products, err := h.repository.BulkCreateProducts(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}
