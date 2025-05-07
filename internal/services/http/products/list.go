package products

import (
	"context"
	"go_template_project/internal/domain"
	"log"
)

func (h Handler) GetProducts(
	ctx context.Context,
	limit, offset int,
) ([]domain.Product, error) {
	products, err := h.repository.GetProducts(ctx, limit, offset)
	if err != nil {
		log.Println(err)
		return products, err
	}
	return products, nil
}
