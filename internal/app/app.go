package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	appHttp "go_template_project/internal/app/http"
	middlewaresHttp "go_template_project/internal/app/http/middlewares"
	"go_template_project/internal/config"
	dbRepo "go_template_project/internal/repository"
	"log"
	"net/http"
	"sync"

	_ "go_template_project/api"
)

type (
	App struct {
		config     config.Config
		repository *dbRepo.Repository
		mux        *http.ServeMux
		server     *http.Server
	}
)

func NewApp(ctx context.Context, config config.Config) (*App, error) {
	conn, err := dbRepo.NewPgxConn(ctx, config.Repository) // DB connection
	if err != nil {
		return nil, err
	}

	repo := dbRepo.NewRepo(conn) // Repository

	// HTTP router for integration layer
	mux := http.NewServeMux()

	// API ------------

	// Swagger (if enabled in config)
	if config.Server.SwaggerDocs {
		mux.Handle("GET /docs/", httpSwagger.WrapHandler)
	}

	// Prometheus exporter
	mux.Handle("GET /metrics/", promhttp.Handler())

	// Internal layer
	appHttp.RegisterRoutes(mux, repo)

	// END API ------------

	// add logging middleware
	httpHandler := middlewaresHttp.LoggingMiddlewareHandler(mux)

	// add cors middleware
	if config.Server.AllowCors {
		httpHandler = middlewaresHttp.AllowCors(httpHandler)
	}

	httpServerAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	// Merge components into app
	return &App{
		config:     config,
		repository: repo,
		mux:        mux,
		server:     &http.Server{Addr: httpServerAddr, Handler: httpHandler},
	}, nil
}

func (a *App) Run(ctx context.Context, wg *sync.WaitGroup) error {
	// Start webserver
	log.Println("Starting HTTP-server")
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Error starting server: %v", err)
		}
	}()

	log.Println("All components started")

	return nil
}

func (a *App) Close() error {
	err := a.server.Close()
	if err != nil {
		return err
	}
	return nil
}
