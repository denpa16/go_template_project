package products

import (
	"context"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) BulkUpdateProducts(
	ctx context.Context,
	data []productsDomain.Product,
) ([]productsDomain.Product, error) {
	updateFields := []string{"name", "title"}
	products, err := h.repository.BulkUpdateProducts(ctx, updateFields, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}
