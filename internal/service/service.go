package service

import (
	repo "forum/internal/repository"
	authService "forum/internal/service/auth"
)

type Service struct {
	authService.Authentication
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		Authentication: authService.NewAuthService(repo.Auth),
	}
}
