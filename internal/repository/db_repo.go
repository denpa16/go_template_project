package repository

import (
	productsRepo "go_template_project/internal/repository/products"
)

type Repository struct {
	conn         Connect
	productsRepo ProductsRepository
}

func NewRepo(conn Connect) *Repository {
	queries := *New(conn)
	return &Repository{
		conn:         conn,
		productsRepo: productsRepo.NewProductsRepository(queries.db),
	}
}

type Queries struct {
	db DBTX
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}
