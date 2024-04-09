package service

import (
	"context"
	"errors"

	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/model"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

var ErrUserNotFound = errors.New("user not found")

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (us *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return us.UserRepository.CreateUser(ctx, user)
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    user, err := us.UserRepository.GetUserByEmail(ctx, email)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, ErrUserNotFound
    }
    return user, nil
}

