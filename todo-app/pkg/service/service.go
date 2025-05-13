package service

import (
	"github.com/VLaDyslav334/pkg/repository"
)

type Authorization interface {
	CreateUser(user repository.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
