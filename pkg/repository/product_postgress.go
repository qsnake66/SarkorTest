package repository

import (
	SarkorTest "github.com/qsnake66/ProductWerehouse"
	"gorm.io/gorm"
)

type ProductPostgresRepository struct {
	db *gorm.DB
}

func NewProductPostgresRepository(db *gorm.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{db: db}
}

func (p *ProductPostgresRepository) CreateProduct(product SarkorTest.Product) (int, error) {

	result := p.db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.Id, nil
}

func (p *ProductPostgresRepository) GetProductById(id int) (SarkorTest.Product, error) {

	var product SarkorTest.Product
	result := p.db.First(&product, id)
	if result.Error != nil {
		return SarkorTest.Product{}, result.Error
	}
	return product, nil
}
func (p *ProductPostgresRepository) GetAllProduct() ([]SarkorTest.Product, error) {
	var products []SarkorTest.Product
	result := p.db.Find(&products)
	if result.Error != nil {
		return []SarkorTest.Product{}, result.Error
	}
	return products, nil
}
func (p *ProductPostgresRepository) UpdateProduct(id int, product SarkorTest.Product) error {
	result := p.db.Model(&product).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *ProductPostgresRepository) DeleteProduct(id int) error {
	result := p.db.Delete(&SarkorTest.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
