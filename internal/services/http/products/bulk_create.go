package products

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) BulkCreateProducts(
	ctx context.Context,
	data []productsDomain.Product,
) (int64, error) {
	count, err := h.repository.BulkCreateProducts(ctx, data)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
