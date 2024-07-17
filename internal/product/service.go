package product

import (
	models "github.com/Jacobo0312/go-web/internal/domain"
)

// ProductService interface
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int64) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int64) error
}

// ProductService struct
type productService struct {
	repo ProductRepository
}

// NewProductService return a new ProductService
func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

// CreateProduct create a new product
func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

// GetAllProducts return all products
func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

// GetProductByID return a product by id
func (s *productService) GetProductByID(id int64) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// UpdateProduct update a product
func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

// DeleteProduct delete a product
func (s *productService) DeleteProduct(id int64) error {
	return s.repo.Delete(id)
}
