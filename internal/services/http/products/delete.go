package products

import (
	"context"
	"errors"
	productsDomain "go_template_project/internal/domain/products"
	"log"
)

func (h Handler) DeleteProduct(
	ctx context.Context,
	data productsDomain.DeleteProductDTO,
) (*productsDomain.Product, error) {
	product, err := h.repository.DeleteProduct(ctx, data)
	if err != nil {
		if errors.Is(err, productsDomain.ErrProductNotFound) {
			return nil, err
		}
		log.Println(err)
		return nil, err
	}
	return product, nil
}
