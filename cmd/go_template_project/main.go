package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go_template_project/internal/app"
	"go_template_project/internal/config"
)

//	@title			GO TEMPLATE PROJECT
//	@version		1.0
//	@description	GO TEMPLATE PROJECT
// @BasePath /

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	var (
		wg   = &sync.WaitGroup{}
		conf = config.NewConfig(envVars)
		ctx  = runSignalHandler(context.Background(), wg)
	)

	// Применяем миграции БД
	err := migrateUp(conf)
	if err != nil {
		log.Fatal("{FATAL} ", err)
	}

	// Создаём новое приложение
	service, err := app.NewApp(ctx, conf)
	if err != nil {
		log.Fatal("{FATAL} ", err)
	}
	// Завершаем приложение gracefully
	defer service.Close()

	service.Run(ctx, wg)

	wg.Wait()
}

func runSignalHandler(ctx context.Context, wg *sync.WaitGroup) context.Context {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	sigCtx, cancel := context.WithCancel(ctx)

	wg.Add(1)
	go func() {
		defer fmt.Println("[signal] terminate")
		defer signal.Stop(sigterm)
		defer wg.Done()
		defer cancel()

		for {
			select {
			case sig, ok := <-sigterm:
				if !ok {
					fmt.Printf("[signal] signal chan closed: %s\n", sig.String())
					return
				}

				fmt.Printf("[signal] signal recv: %s\n", sig.String())
				return
			case _, ok := <-sigCtx.Done():
				if !ok {
					fmt.Println("[signal] context closed")
					return
				}

				fmt.Printf("[signal] ctx done: %s\n", ctx.Err().Error())
				return
			}
		}
	}()

	return sigCtx
}
