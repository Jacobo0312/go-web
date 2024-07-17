package repositories

import (
	"database/sql"

	"github.com/Jacobo0312/go-web/internal/models"
)

type ProductRepository interface {
	Create(p *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id int64) (*models.Product, error)
	Update(p *models.Product) error
	Delete(id int64) error
}

type productRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) Create(p *models.Product) error {
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

func (r *productRepository) GetAll() ([]models.Product, error) {
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

func (r *productRepository) GetByID(id int64) (*models.Product, error) {
	query := "SELECT id, name, price, description, category FROM products WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var p models.Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Category)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) Update(p *models.Product) error {
	query := "UPDATE products SET name = ?, price = ?, description = ?, category = ? WHERE id = ?"
	_, err := r.DB.Exec(query, p.Name, p.Price, p.Description, p.Category, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Delete(id int64) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
