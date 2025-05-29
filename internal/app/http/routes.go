package http

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go_template_project/api"
	productsRoutes "go_template_project/internal/app/http/products"
	"go_template_project/internal/config"
	dbRepo "go_template_project/internal/repository"
	"net/http"
)

func RegisterRoutes(
	config config.Config,
	mux *http.ServeMux,
	repo *dbRepo.Repository,
) {

	// Swagger (if enabled in config)
	if config.Server.SwaggerDocs {
		mux.Handle("GET /docs/", httpSwagger.WrapHandler)
	}

	// Prometheus exporter
	mux.Handle("GET /metrics/", promhttp.Handler())
	productsRoutes.RegisterRoutes(mux, repo)
}
