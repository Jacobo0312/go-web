package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/Jacobo0312/go-web/pkg/errors"
	"github.com/Jacobo0312/go-web/pkg/test"
	"github.com/stretchr/testify/mock"
)

type mockProductService struct {
	mock.Mock
}

func (m *mockProductService) CreateProduct(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductService) GetAllProducts() ([]domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *mockProductService) GetProductByID(id int64) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *mockProductService) UpdateProduct(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductService) DeleteProduct(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHandlerCreateProduct(t *testing.T) {
	mockService := new(mockProductService)
	handler := NewProductHandler(mockService)

	testCases := []test.HandlerTestCase{
		{
			Name:             "successful creation",
			Method:           "POST",
			URL:              "/products",
			Body:             `{"id":0,"name":"Test Product","price":9.99,"category":"","description":""}`,
			ExpectedStatus:   http.StatusCreated,
			ExpectedResponse: `{"id":0,"name":"Test Product","price":9.99,"category":"","description":""}`},
		{
			Name:           "invalid payload",
			Method:         "POST",
			URL:            "/products",
			Body:           "invalid json",
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Name:           "service error",
			Method:         "POST",
			URL:            "/products",
			Body:           `{"name":"Error Product","price":19.99}`,
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		if tc.Name == "successful creation" {
			product := &domain.Product{Name: "Test Product", Price: 9.99}
			mockService.On("CreateProduct", product).Return(nil).Once()
		}else if tc.Name == "service error" {
			product := &domain.Product{Name: "Error Product", Price: 19.99}
			mockService.On("CreateProduct", product).Return(errors.NewBadRequest("Invalid request payload", nil)).Once()
		}

		test.ExecuteHandlerTestCase(t, handler.CreateProduct, tc)
	}

	mockService.AssertExpectations(t)
}

func TestHandlerGetAllProducts(t *testing.T) {
	mockService := new(mockProductService)
	handler := NewProductHandler(mockService)

	products := []domain.Product{{ID: 1, Name: "Product 1"}, {ID: 2, Name: "Product 2"}}
	productsJSON, _ := json.Marshal(products)

	testCases := []test.HandlerTestCase{
		{
			Name:             "successful retrieval",
			Method:           "GET",
			URL:              "/products",
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: string(productsJSON),
		},
		{
			Name:           "service error",
			Method:         "GET",
			URL:            "/products",
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		if tc.Name == "successful retrieval" {
			mockService.On("GetAllProducts").Return(products, nil).Once()
		} else if tc.Name == "service error" {
			mockService.On("GetAllProducts").Return([]domain.Product{}, errors.NewInternalServerError("Error getting products", nil)).Once()
		}

		test.ExecuteHandlerTestCase(t, handler.GetAllProducts, tc)
	}

	mockService.AssertExpectations(t)
}

// func TestHandlerGetProductByID(t *testing.T) {

// 	mockService := new(mockProductService)
// 	handler := NewProductHandler(mockService)
	
// 	product := &domain.Product{
// 		ID:          1,
// 		Name:        "Audifonos",
// 		Price:       19.99,
// 		Description: "Marca KZ",
// 		Category:    "Audio",
// 	}

// 	mockService.On("GetProductByID", int64(24)).Return(product, nil).Once()

// 	testCases := []test.HandlerTestCase{
// 		{
// 			Name:             "successful retrieval",
// 			Method:           "GET",
// 			URL:              "/products/1",
// 			ExpectedStatus:   http.StatusOK,
// 			ExpectedResponse: `{"id": 1,"name": "Audifonos","price": 19.99,"description": "Marca KZ","category": "Audio"}`,
// 		},
// 		// {
// 		//     Name:           "product not found",
// 		//     Method:         "GET",
// 		//     URL:            "/products/999",
// 		//     ExpectedStatus: http.StatusNotFound,
// 		// },
// 		// {
// 		//     Name:           "invalid id",
// 		//     Method:         "GET",
// 		//     URL:            "/products/invalid",
// 		//     ExpectedStatus: http.StatusBadRequest,
// 		// },
// 	}

// 	for _, tc := range testCases {
// 		test.ExecuteHandlerTestCase(t, handler.GetProductByID, tc)
// 	}

// 	mockService.AssertExpectations(t)
// }

// func TestHandlerDeleteProduct(t *testing.T) {
// 	mockService := new(mockProductService)
// 	handler := NewProductHandler(mockService)

// 	testCases := []test.HandlerTestCase{
// 		{
// 			Name:           "successful deletion",
// 			Method:         "DELETE",
// 			URL:            "/products/1",
// 			ExpectedStatus: http.StatusNoContent,
// 		},
// 		// {
// 		// 	Name:           "product not found",
// 		// 	Method:         "DELETE",
// 		// 	URL:            "/products/999",
// 		// 	ExpectedStatus: http.StatusNotFound,
// 		// },
// 		// {
// 		// 	Name:           "invalid id",
// 		// 	Method:         "DELETE",
// 		// 	URL:            "/products/invalid",
// 		// 	ExpectedStatus: http.StatusBadRequest,
// 		// },
// 	}

// 	for _, tc := range testCases {
// 		if tc.Name == "successful deletion" {
// 			mockService.On("DeleteProduct", int64(1)).Return(nil).Once()
// 		}

// 		// else if tc.Name == "product not found" {
// 		//     mockService.On("DeleteProduct", int64(999)).Return(errors.New("not found")).Once()
// 		// }

// 		test.ExecuteHandlerTestCase(t, handler.DeleteProduct, tc)
// 	}

// 	mockService.AssertExpectations(t)
// }
