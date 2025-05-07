package products

import (
	"context"
	"go_template_project/internal/domain"
)

type repository interface {
	GetProducts(
		ctx context.Context,
		limit, offset int,
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
