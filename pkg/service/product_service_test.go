package service

import (
	SarkorTest "github.com/qsnake66/ProductWerehouse"
	"github.com/qsnake66/ProductWerehouse/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_CreateProduct(t *testing.T) {
	mockRepo := new(mocks.Product)
	mockService := NewProductService(mockRepo)

	product := SarkorTest.Product{
		Id:          1,
		Name:        "Test Product",
		Price:       100.0,
		Description: "Test Description",
		Quantity:    10,
	}

	mockRepo.On("CreateProduct", product).Return(1, nil)

	id, err := mockService.CreateProduct(product)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductById(t *testing.T) {
	mockRepo := new(mocks.Product)
	productService := NewProductService(mockRepo)

	product := SarkorTest.Product{
		Id:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100.0,
		Quantity:    10,
	}

	mockRepo.On("GetProductById", 1).Return(product, nil)

	result, err := productService.GetProductById(1)

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetAllProduct(t *testing.T) {
	mockRepo := new(mocks.Product)
	productService := NewProductService(mockRepo)

	products := []SarkorTest.Product{
		{
			Id:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       100.0,
			Quantity:    10,
		},
		{
			Id:          2,
			Name:        "Product 2",
			Description: "Description 2",
			Price:       200.0,
			Quantity:    20,
		},
	}

	mockRepo.On("GetAllProduct").Return(products, nil)

	result, err := productService.GetAllProduct()

	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct(t *testing.T) {
	mockRepo := new(mocks.Product)
	productService := NewProductService(mockRepo)

	product := SarkorTest.Product{
		Id:          1,
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       150.0,
		Quantity:    15,
	}

	mockRepo.On("UpdateProduct", 1, product).Return(nil)

	err := productService.UpdateProduct(1, product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProduct(t *testing.T) {
	mockRepo := new(mocks.Product)
	productService := NewProductService(mockRepo)

	mockRepo.On("DeleteProduct", 1).Return(nil)

	err := productService.DeleteProduct(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
