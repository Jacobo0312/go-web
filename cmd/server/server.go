package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Jacobo0312/go-web/config"
	"github.com/Jacobo0312/go-web/internal/handlers"
	"github.com/Jacobo0312/go-web/internal/repositories"
	"github.com/Jacobo0312/go-web/internal/services"
	"github.com/Jacobo0312/go-web/pkg/middlewares"
)

type Server struct {
	config *config.Config
	router *http.ServeMux
	db     *sql.DB
}

func New(cfg *config.Config, db *sql.DB) *Server {
	return &Server{
		config: cfg,
		router: http.NewServeMux(),
		db:     db,
	}
}

func (s *Server) Start() error {

	//Health check
	s.router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	//Product
	productRepo := repositories.NewProductRepository(s.db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	productHandler.RegisterRoutes(s.router)

	//User
	userRepo := repositories.NewUserRepository(s.db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userHandler.RegisterRoutes(s.router)

	middleware := middlewares.MiddlewareChain()

	log.Printf("Starting server on %s", s.config.ServerAddr)
	server := &http.Server{
		Addr:    s.config.ServerAddr,
		Handler: middleware(s.router),
	}

	return server.ListenAndServe()
}
