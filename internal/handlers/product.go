package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jacobo0312/go-web/internal/models"
	"github.com/Jacobo0312/go-web/internal/services"

	//"github.com/Jacobo0312/go-web/pkg/middlewares"
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
	service services.ProductService
}

func NewProductHandler(service services.ProductService) ProductHandler {
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
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Get Product by ID
func (h *productHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := helpers.ReadIdParam(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}

// Update Product
func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Delete Product
func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIdParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
