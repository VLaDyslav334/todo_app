package service

import (
	"github.com/VLaDyslav334/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user repository.User) (int, error) {
	return s.repo.CreateUser(user)
}
