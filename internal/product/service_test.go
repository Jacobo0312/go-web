package product

import (
	"errors"
	"testing"

	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockProductRepository struct {
	mock.Mock
}

func (m *mockProductRepository) Create(p *domain.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *mockProductRepository) GetAll() ([]domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *mockProductRepository) GetByID(id int64) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *mockProductRepository) Update(p *domain.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *mockProductRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceCreateProduct(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := NewProductService(mockRepo)

	t.Run("successful product creation", func(t *testing.T) {
		product := &domain.Product{Name: "Test Product", Price: 9.99}
		mockRepo.On("Create", product).Return(nil)

		err := service.CreateProduct(product)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		product := &domain.Product{Name: "Error Product", Price: 19.99}
		mockRepo.On("Create", product).Return(errors.New("database error"))

		err := service.CreateProduct(product)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestServiceGetAllProducts(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := NewProductService(mockRepo)

	t.Run("successful get all products", func(t *testing.T) {
		expectedProducts := []domain.Product{
			{ID: 1, Name: "Product 1", Price: 9.99},
			{ID: 2, Name: "Product 2", Price: 19.99},
		}
		mockRepo.On("GetAll").Return(expectedProducts, nil)

		products, err := service.GetAllProducts()

		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, products)
		mockRepo.AssertExpectations(t)
	})

}

func TestServiceGetProductByID(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := NewProductService(mockRepo)

	t.Run("product found", func(t *testing.T) {
		expectedProduct := &domain.Product{ID: 1, Name: "Test Product", Price: 9.99}
		mockRepo.On("GetByID", int64(1)).Return(expectedProduct, nil)

		product, err := service.GetProductByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("product not found", func(t *testing.T) {
		mockRepo.On("GetByID", int64(2)).Return((*domain.Product)(nil), errors.New("product not found"))

		product, err := service.GetProductByID(2)

		assert.Error(t, err)
		assert.Nil(t, product)
		mockRepo.AssertExpectations(t)
	})
}

func TestServiceUpdateProduct(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := NewProductService(mockRepo)

	t.Run("successful update", func(t *testing.T) {
		product := &domain.Product{ID: 1, Name: "Updated Product", Price: 29.99}
		mockRepo.On("Update", product).Return(nil)

		err := service.UpdateProduct(product)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("update error", func(t *testing.T) {
		product := &domain.Product{ID: 2, Name: "Error Product", Price: 39.99}
		mockRepo.On("Update", product).Return(errors.New("database error"))

		err := service.UpdateProduct(product)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestServiceDeleteProduct(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := NewProductService(mockRepo)

	t.Run("successful delete", func(t *testing.T) {
		mockRepo.On("Delete", int64(1)).Return(nil)

		err := service.DeleteProduct(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("delete error", func(t *testing.T) {
		mockRepo.On("Delete", int64(2)).Return(errors.New("database error"))

		err := service.DeleteProduct(2)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}