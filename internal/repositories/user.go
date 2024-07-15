package repositories

import (
	"database/sql"

	"github.com/Jacobo0312/go-web/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Register(u *models.User) error {
	query := "INSERT INTO users (id, name, email, role) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, u.ID, u.Name, u.Email, u.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	query := "SELECT id, name, email, role FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
