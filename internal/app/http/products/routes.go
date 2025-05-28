package products

import (
	dbRepo "go_template_project/internal/repository"
	command "go_template_project/internal/services/http/products"
	"net/http"
)

func RegisterRoutes(
	mux *http.ServeMux,
	repo *dbRepo.Repository,
) {
	// Get products
	mux.Handle(
		"GET /api/products/",
		NewProductsGetHandler(
			command.New(repo),
			"GET /api/products/",
		),
	)

	// Get product
	mux.Handle(
		"GET /api/products/{id}",
		NewProductGetHandler(
			command.New(repo),
			"GET /api/products/{id}",
		),
	)

	// Create product
	mux.Handle(
		"POST /api/products",
		NewProductCreateHandler(
			command.New(repo),
			"POST /api/products",
		),
	)

	// Partial update product
	mux.Handle(
		"PATCH /api/products/{id}",
		NewProductPartialUpdateHandler(
			command.New(repo),
			"PATCH /api/products/{id}",
		),
	)
	// Delete product
	mux.Handle(
		"DELETE /api/products/{id}",
		NewProductDeleteHandler(
			command.New(repo),
			"DELETE /api/products/{id}",
		),
	)
}
