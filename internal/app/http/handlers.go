package http

import (
	dbRepo "go_template_project/internal/repository"
	httpCommandGetProducts "go_template_project/internal/services/http/products_get"
	"net/http"
)

func RegisterHandlers(
	mux *http.ServeMux,
	repo *dbRepo.Repository,
) {
	// Get products_get
	mux.Handle(
		"GET /api/products/",
		NewProductsGetHandler(
			httpCommandGetProducts.New(repo),
			"GET /api/products_get/",
		),
	)
}
