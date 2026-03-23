package services

import (
	"crm/internal/domain"
	"crm/internal/repository"
	"errors"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	existing, err := s.repo.FindUserByID(user.Email)
	if err != nil {
		return errors.New("Internal error")
	}
	if existing != nil {
		return errors.New("Invalid credentials")
	}

	return s.repo.CreateUser(user)
}
