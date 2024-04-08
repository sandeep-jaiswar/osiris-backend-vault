package service

import (
	"context"

	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/model"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/repository"
)

type UserService struct {
    UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{UserRepository: userRepo}
}

func (us *UserService) CreateUser(ctx context.Context, user *models.User) error {
    return us.UserRepository.CreateUser(ctx, user)
}
