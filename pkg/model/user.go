package models

import (
    "time"

    "github.com/go-playground/validator/v10"
    "gorm.io/gorm"
)

type User struct {
    ID        uint       `gorm:"primaryKey;autoIncrement"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`

    Email     string `gorm:"unique;not null" validate:"required,email"`
    FirstName string
    LastName  string
    UserType  UserType `gorm:"type:enum('Admin', 'Business_Owner', 'Customer');default:'Customer'"`
}

// Validate performs validation on the user struct.
func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}

// UserType represents the type of user.
type UserType string

// Enum values for UserType.
const (
    UserTypeAdmin        UserType = "Admin"
    UserTypeBusinessOwner UserType = "Business_Owner"
    UserTypeCustomer     UserType = "Customer"
)
