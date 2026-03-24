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

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	existing, err := s.repo.FindUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("internal error")
	}
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	newUser, err := domain.NewUser(
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) ViewUser(user *domain.User) (*domain.User, error) {
	return s.repo.FindUserByID(user.UUID)
}

func (s *UserService) DeleteUser(user *domain.User) error {
	return s.repo.DeleteUserById(user.UUID)
}
