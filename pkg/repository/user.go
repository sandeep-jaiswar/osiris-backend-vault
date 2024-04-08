package repository

import (
	"context"

	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/model"
	"gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

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
