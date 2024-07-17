package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/Jacobo0312/go-web/internal/product"

	//"github.com/Jacobo0312/go-web/pkg/middlewares"
	"github.com/Jacobo0312/go-web/pkg/errors"
	"github.com/Jacobo0312/go-web/pkg/helpers"
)

// ProductHandler interface}
type ProductHandler interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetProductByID(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(r *http.ServeMux)
}

type productHandler struct {
	service product.ProductService
}

func NewProductHandler(service product.ProductService) ProductHandler {
	return &productHandler{service: service}
}

// Register routes
func (h *productHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /products", h.CreateProduct)
	//Protected route
	//r.HandleFunc("GET /products", middlewares.FirebaseAuthMiddleware(h.GetAllProducts))
	r.HandleFunc("GET /products", h.GetAllProducts)
	r.HandleFunc("GET /products/{id}", h.GetProductByID)
	r.HandleFunc("PUT /products", h.UpdateProduct)
	r.HandleFunc("DELETE /products/{id}", h.DeleteProduct)
}

func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.RespondWithError(w, errors.NewBadRequest("Invalid request payload", err))
		return
	}

	err = h.service.CreateProduct(&product)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternalServerError("Error creating product", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, product)
}

// Get All Products
func (h *productHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Get userID from context
	// userID, ok := r.Context().Value("userID").(string)
	// if !ok {
	// 	http.Error(w, "Invalid user", http.StatusUnauthorized)
	// 	return
	// }
	//------------------------
	products, err := h.service.GetAllProducts()
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternalServerError("Error getting products", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, products)
}

// Get Product by ID
func (h *productHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := helpers.ReadIdParam(r)

	if err != nil {
		helpers.RespondWithError(w, errors.NewBadRequest("Invalid product ID", err))
		return
	}

	product, err := h.service.GetProductByID(id)

	if err != nil {
		helpers.RespondWithError(w, errors.NewNotFound("Product not found", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, product)

}

// Update Product
func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.RespondWithError(w, errors.NewBadRequest("Invalid request payload", err))
		return
	}

	err = h.service.UpdateProduct(&product)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternalServerError("Error updating product", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, product)
}

// Delete Product
func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIdParam(r)
	if err != nil {
		helpers.RespondWithError(w, errors.NewBadRequest("Invalid product ID", err))
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternalServerError("Error deleting product", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusNoContent, nil)
}
