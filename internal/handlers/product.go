package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jacobo0312/go-web/internal/models"
	"github.com/Jacobo0312/go-web/internal/services"
	"github.com/Jacobo0312/go-web/pkg/middlewares"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// Register routes
func (h *ProductHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /products", h.CreateProduct)
	//Protected route
	r.HandleFunc("GET /products", middlewares.FirebaseAuthMiddleware(h.GetAllProducts))
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Get userID from context
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}
	//------------------------
	fmt.Println("UserID:", userID)
	products, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
