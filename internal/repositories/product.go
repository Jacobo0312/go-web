package repositories

import (
	"database/sql"

	"github.com/Jacobo0312/go-web/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(p *models.Product) error {
	query := "INSERT INTO products (name, price, description, category) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, p.Name, p.Price, p.Description, p.Category)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, description, category FROM products"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
