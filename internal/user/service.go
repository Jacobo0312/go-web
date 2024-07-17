package user

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/Jacobo0312/go-web/internal/domain"
	"github.com/Jacobo0312/go-web/pkg/firebase"
)

type UserService interface {
	CreateUser(ctx context.Context, userRequest *domain.CreateUserRequest) (*domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, userRequest *domain.CreateUserRequest) (*domain.User, error) {

	params := (&auth.UserToCreate{}).
		Email(userRequest.Email).
		EmailVerified(false).
		Password(userRequest.Password).
		DisplayName(userRequest.Name).
		Disabled(false)

	user, err := firebase.FirebaseAuth.CreateUser(ctx, params)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	defer func() {
		if err != nil {
			if err := firebase.FirebaseAuth.DeleteUser(ctx, user.UID); err != nil {
				log.Printf("Error deleting user from Firebase: %v", err)
			}
		}
	}()

	userModel := &domain.User{
		ID:    user.UID,
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Role:  userRequest.Role,
	}

	err = s.repo.Register(userModel)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return userModel, nil

}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}
