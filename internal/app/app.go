package app

import (
	"context"
	"errors"
	"fmt"
	appHttp "go_template_project/internal/app/http"
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
		server     *http.Server
	}
)

func NewApp(ctx context.Context, config config.Config) (*App, error) {
	// DB connection
	conn, err := dbRepo.NewPgxConn(ctx, config.Repository)
	if err != nil {
		return nil, err
	}

	// Repository
	repo := dbRepo.NewRepo(conn)

	// HTTP router
	mux := appHttp.RegisterRoutes(config, repo)

	// Merge components into app
	return &App{
		config:     config,
		repository: repo,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
			Handler: mux,
		},
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
