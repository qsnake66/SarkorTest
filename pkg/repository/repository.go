package repository

import (
	"database/sql"

	SarkorTest "github.com/qsnake66/ProductWerehouse"
)

type Product interface {
	CreateProduct(product SarkorTest.Product) (int, error)
	GetProductById(id int) (SarkorTest.Product, error)
	GetAllProduct() ([]SarkorTest.Product, error)
	UpdateProduct(id int, product SarkorTest.Product) error
	DeleteProduct(id int) error
}

type Repository struct {
	Product
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Product: NewProductPostgresRepository(db)}
}
