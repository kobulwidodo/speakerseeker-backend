package usecase

import (
	"errors"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/middleware"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{UserRepository: ur}
}

func (u *UserUseCase) SignUp(input *domain.UserSignUp) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user := domain.User{
		OrganizationName: input.OrganizationName,
		Email:            input.Email,
		Password:         string(hash),
	}
	if err := u.UserRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) SignIn(input *domain.UserSignIn) (string, error) {
	user, err := u.UserRepository.GetByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("kredential not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
