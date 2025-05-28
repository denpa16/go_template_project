package products

import (
	"context"
	"errors"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) GetProduct(
	ctx context.Context,
	data productsDomain.GetProductDTO,
) (*productsDomain.Product, error) {
	product, err := h.repository.GetProduct(ctx, data)
	if err != nil {
		if errors.Is(err, productsDomain.ErrProductNotFound) {
			return nil, err
		}
		log.Println(err)
		return nil, err
	}
	return product, nil
}
