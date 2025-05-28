package http

import (
	productsRoutes "go_template_project/internal/app/http/products"
	dbRepo "go_template_project/internal/repository"
	"net/http"
)

func RegisterRoutes(
	mux *http.ServeMux,
	repo *dbRepo.Repository,
) {
	productsRoutes.RegisterRoutes(mux, repo)
}
