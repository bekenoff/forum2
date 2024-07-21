package service

import (
	"forum/internal/models"
	"forum/pkg"
)

type Authentication interface {
	CreateClient(registration models.UserRegistration) error
}

type AuthService struct {
	repository Authentication
}

func NewAuthService(repository Authentication) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (u *AuthService) CreateClient(user models.UserRegistration) error {
	var err error
	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}
	return u.repository.CreateClient(user)
}
