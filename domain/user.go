package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email            string `gorm:"unique"`
	OrganizationName string
	Password         string
}

type UserRepository interface {
	Create(user User) error
	GetByEmail(email string) (User, error)
}

type UserUseCase interface {
	SignUp(input *UserSignUp) error
	SignIn(input *UserSignIn) (string, error)
}

type UserSignUp struct {
	Email            string `binding:"required"`
	OrganizationName string `binding:"required" json:"organization_name"`
	Password         string `binding:"required"`
}

type UserSignIn struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
