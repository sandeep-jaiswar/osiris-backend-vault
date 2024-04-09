package repository

import (
	"context"
	"errors"

	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	result := ur.DB.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := ur.DB.WithContext(ctx).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    }
    if result.Error != nil {
        return nil, result.Error
    }
	return &user, nil
}
