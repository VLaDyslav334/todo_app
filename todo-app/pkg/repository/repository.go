package repository

import (
	"github.com/VLaDyslav334/todo-app/pkg/todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(param todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func (r Repository) CreateUser(param any) (int, error) {
	panic("unimplemented")
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
