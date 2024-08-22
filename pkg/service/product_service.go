package service

import (
	SarkorTest "github.com/qsnake66/ProductWerehouse"
	"github.com/qsnake66/ProductWerehouse/pkg/repository"
)

type ProductService struct {
	Repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{Repo: repo}
}

func (p *ProductService) CreateProduct(product SarkorTest.Product) (int, error) {
	return p.Repo.CreateProduct(product)
}

func (p *ProductService) GetProductById(id int) (SarkorTest.Product, error) {
	return p.Repo.GetProductById(id)
}
func (p *ProductService) GetAllProduct() ([]SarkorTest.Product, error) {
	return p.Repo.GetAllProduct()
}
func (p *ProductService) UpdateProduct(id int, product SarkorTest.Product) error {
	return p.Repo.UpdateProduct(id, product)
}
func (p *ProductService) DeleteProduct(id int) error {
	return p.Repo.DeleteProduct(id)
}
