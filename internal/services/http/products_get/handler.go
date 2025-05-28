package products_get

import (
	"context"
	"go_template_project/internal/domain"
	"log"
)

type repository interface {
	GetProducts(
		ctx context.Context,
		params domain.GetProductsDTO,
	) ([]domain.Product, error)
}

func New(repo repository) Handler {
	return Handler{
		repository: repo,
	}
}

type Handler struct {
	repository repository
}

func (h Handler) GetProducts(
	ctx context.Context,
	params domain.GetProductsDTO,
) ([]domain.Product, error) {
	products, err := h.repository.GetProducts(ctx, params)
	if err != nil {
		log.Println(err)
		return products, err
	}
	return products, nil
}
