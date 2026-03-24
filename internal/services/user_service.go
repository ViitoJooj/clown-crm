package services

import (
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/repository"
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

	user, err = domain.NewUser(user.First_Name, user.Last_Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) ViewUser(user *domain.User) (*domain.User, error) {
	return s.repo.FindUserByID(user.UUID)
}

func (s *UserService) DeleteUser(user *domain.User) error {
	return s.repo.DeleteUserById(user.UUID)
}
