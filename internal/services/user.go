package services

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/Jacobo0312/go-web/internal/models"
	"github.com/Jacobo0312/go-web/internal/repositories"
	"github.com/Jacobo0312/go-web/pkg/firebase"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, userRequest *models.CreateUserRequest) (*models.User, error) {

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
			firebase.FirebaseAuth.DeleteUser(ctx, user.UID)
		}
	}()

	userModel := &models.User{
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
