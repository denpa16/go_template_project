package main

import (
	"context"
	"fmt"
	"go_template_project/internal/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dialect string = "pgx"
	dir     string = "./migrations"
	command string = "up"
)

func migrateUp(conf config.Config) error {
	log.Println("Start migrations...")

	ctx := context.Background()

	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Repository.Host,
		conf.Repository.Username,
		conf.Repository.Password,
		conf.Repository.Name,
		conf.Repository.Port,
	)

	db, err := goose.OpenDBWithDriver(dialect, dbString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	if err := goose.RunContext(ctx, command, db, dir); err != nil {
		log.Fatalf("migrate %v: %v", command, err)
	}

	log.Println("Migrations completed.")

	return nil
}
