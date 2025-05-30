package http

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go_template_project/api"
	middlewaresHttp "go_template_project/internal/app/http/middlewares"
	productsRoutes "go_template_project/internal/app/http/products"
	"go_template_project/internal/config"
	dbRepo "go_template_project/internal/repository"
	"net/http"
)

func RegisterRoutes(
	config config.Config,
	repo *dbRepo.Repository,
) http.Handler {
	mux := http.NewServeMux()

	// Swagger (if enabled in config)
	if config.Server.SwaggerDocs {
		mux.Handle("GET /docs/", httpSwagger.WrapHandler)
	}

	// add logging middleware
	httpHandler := middlewaresHttp.LoggingMiddlewareHandler(mux)

	// add cors middleware
	if config.Server.AllowCors {
		httpHandler = middlewaresHttp.AllowCors(httpHandler)
	}

	// Prometheus exporter
	mux.Handle("GET /metrics/", promhttp.Handler())
	productsRoutes.RegisterRoutes(mux, repo)

	return mux
}
