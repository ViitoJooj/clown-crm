package repository

import "crm/internal/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindUserByID(id string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUserById(id string) error
}

type InMemoryUserRepository struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}
