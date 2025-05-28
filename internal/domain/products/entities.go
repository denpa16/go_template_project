package products

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type GetProductsDTO struct {
	Limit  int64  `json:"limit,omitempty"`
	Offset int64  `json:"offset,omitempty"`
	Name   string `json:"name,omitempty"`
	Title  string `json:"title,omitempty"`
}

type GetProductDTO struct {
	ID uuid.UUID `json:"id"`
}

type CreateProductDTO struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}
