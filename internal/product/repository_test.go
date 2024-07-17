package product

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewProductRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)
	assert.NotNil(t, repo)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	t.Run("successful creation", func(t *testing.T) {
		product := &domain.Product{Name: "Test Product", Price: 9.99, Description: "Test Description", Category: "Test Category"}
		mock.ExpectExec("INSERT INTO products").WithArgs(product.Name, product.Price, product.Description, product.Category).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(product)
		assert.NoError(t, err)
		assert.Equal(t, 1, product.ID)
	})

	t.Run("creation error", func(t *testing.T) {
		product := &domain.Product{Name: "Error Product", Price: 19.99, Description: "Error Description", Category: "Error Category"}
		mock.ExpectExec("INSERT INTO products").WithArgs(product.Name, product.Price, product.Description, product.Category).WillReturnError(errors.New("database error"))

		err := repo.Create(product)
		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	t.Run("get all products", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "category"}).
			AddRow(1, "Product 1", 9.99, "Description 1", "Category 1").
			AddRow(2, "Product 2", 19.99, "Description 2", "Category 2")
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnRows(rows)

		products, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, "Product 2", products[1].Name)
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnError(errors.New("database error"))

		products, err := repo.GetAll()
		assert.Error(t, err)
		assert.Nil(t, products)
	})
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	t.Run("product found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "category"}).
			AddRow(1, "Test Product", 9.99, "Test Description", "Test Category")
		mock.ExpectQuery("SELECT (.+) FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows)

		product, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, "Test Product", product.Name)
	})

	t.Run("product not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM products WHERE id = ?").WithArgs(2).WillReturnError(sql.ErrNoRows)

		product, err := repo.GetByID(2)
		assert.Error(t, err)
		assert.Nil(t, product)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	t.Run("successful update", func(t *testing.T) {
		product := &domain.Product{ID: 1, Name: "Updated Product", Price: 29.99, Description: "Updated Description", Category: "Updated Category"}
		mock.ExpectExec("UPDATE products SET").WithArgs(product.Name, product.Price, product.Description, product.Category, product.ID).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(product)
		assert.NoError(t, err)
	})

	t.Run("update error", func(t *testing.T) {
		product := &domain.Product{ID: 2, Name: "Error Product", Price: 39.99, Description: "Error Description", Category: "Error Category"}
		mock.ExpectExec("UPDATE products SET").WithArgs(product.Name, product.Price, product.Description, product.Category, product.ID).WillReturnError(errors.New("database error"))

		err := repo.Update(product)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	t.Run("successful delete", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM products WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Delete(1)
		assert.NoError(t, err)
	})

	t.Run("delete error", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM products WHERE id = ?").WithArgs(2).WillReturnError(errors.New("database error"))

		err := repo.Delete(2)
		assert.Error(t, err)
	})
}