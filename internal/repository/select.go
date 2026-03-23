package repository

import "github.com/ViitoJooj/clown-crm/internal/domain"

func (r *InMemoryUserRepository) ListUsers() ([]*domain.User, error) {

	return nil, nil
}

func (r *InMemoryUserRepository) FindUserByID(id string) (*domain.User, error) {

	return nil, nil
}

func (r *InMemoryUserRepository) FindUserByEmail(email string) (*domain.User, error) {

	return nil, nil
}
