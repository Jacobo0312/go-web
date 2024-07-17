package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/Jacobo0312/go-web/internal/domain"
	"github.com/Jacobo0312/go-web/internal/user"
	"github.com/Jacobo0312/go-web/pkg/errors"
	"github.com/Jacobo0312/go-web/pkg/helpers"
)

// UserHandler interface
type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(r *http.ServeMux)
}

type userHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) UserHandler {
	return &userHandler{service: service}
}

// Register routes
func (h *userHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /users", h.CreateUser)
	r.HandleFunc("GET /users", h.GetUsers)

}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createUser, err := h.service.CreateUser(r.Context(), &user)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternalServerError("Error creating user", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, createUser)
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternalServerError("Error getting users", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, users)
}
