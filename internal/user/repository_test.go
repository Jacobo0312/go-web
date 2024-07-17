package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	assert.NotNil(t, repo)
}

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	t.Run("successful registration", func(t *testing.T) {
		user := &domain.User{ID: "1", Name: "John Doe", Email: "john@example.com", Role: "user"}
		mock.ExpectExec("INSERT INTO users").WithArgs(user.ID, user.Name, user.Email, user.Role).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Register(user)
		assert.NoError(t, err)
	})

	t.Run("registration error", func(t *testing.T) {
		user := &domain.User{ID: "2", Name: "Jane Doe", Email: "jane@example.com", Role: "user"}
		mock.ExpectExec("INSERT INTO users").WithArgs(user.ID, user.Name, user.Email, user.Role).WillReturnError(errors.New("database error"))

		err := repo.Register(user)
		assert.Error(t, err)
	})
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	t.Run("user found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow("1", "John Doe", "john@example.com", "user")
		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").WithArgs("1").WillReturnRows(rows)

		user, err := repo.FindByID("1")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "1", user.ID)
		assert.Equal(t, "John Doe", user.Name)
	})

	t.Run("user not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").WithArgs("2").WillReturnError(sql.ErrNoRows)

		user, err := repo.FindByID("2")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	t.Run("get all users", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow("1", "John Doe", "john@example.com", "user").
			AddRow("2", "Jane Doe", "jane@example.com", "admin")
		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

		users, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, "John Doe", users[0].Name)
		assert.Equal(t, "Jane Doe", users[1].Name)
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnError(errors.New("database error"))

		users, err := repo.GetAll()
		assert.Error(t, err)
		assert.Nil(t, users)
	})
}