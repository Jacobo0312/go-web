package user

import (
	"testing"

	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Register(u *domain.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *mockUserRepository) FindByID(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *mockUserRepository) GetAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func TestServiceGetUsers(t *testing.T) {
	mockRepo := new(mockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("successful get users", func(t *testing.T) {
		expectedUsers := []domain.User{
			{ID: "1", Name: "User 1", Email: "user1@example.com", Role: "user"},
			{ID: "2", Name: "User 2", Email: "user2@example.com", Role: "admin"},
		}

		mockRepo.On("GetAll").Return(expectedUsers, nil)

		users, err := service.GetUsers()

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)

		mockRepo.AssertExpectations(t)
	})
}
