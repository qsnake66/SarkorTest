package service

import (
	SarkorTest "github.com/qsnake66/ProductWerehouse"
	"github.com/qsnake66/ProductWerehouse/pkg/repository"
)

type Product interface {
	CreateProduct(product SarkorTest.Product) (int, error)
	GetProductById(id int) (SarkorTest.Product, error)
	GetAllProduct() ([]SarkorTest.Product, error)
	UpdateProduct(id int, product SarkorTest.Product) error
	DeleteProduct(id int) error
}

type Service struct {
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repos),
	}
}
