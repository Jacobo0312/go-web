package services

import (
	"github.com/Jacobo0312/go-web/internal/models"
	"github.com/Jacobo0312/go-web/internal/repositories"
)

// ProductService struct
type ProductService struct {
	repo *repositories.ProductRepository
}

// NewProductService return a new ProductService
func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct create a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

// GetAllProducts return all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}
