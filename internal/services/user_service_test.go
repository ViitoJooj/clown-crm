package services_test

import (
	"errors"
	"testing"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/services"
)

type MockUserRepository struct {
	users map[string]*domain.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (m *MockUserRepository) CreateUser(user *domain.User) error {
	m.users[user.UUID] = user
	return nil
}

func (m *MockUserRepository) FindUserByID(id string) (*domain.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (m *MockUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) ListUsers() ([]*domain.User, error) {
	var users []*domain.User
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func (m *MockUserRepository) UpdateUser(user *domain.User) error {
	_, ok := m.users[user.UUID]
	if !ok {
		return errors.New("user not found")
	}
	m.users[user.UUID] = user
	return nil
}

func (m *MockUserRepository) DeleteUserById(id string) error {
	delete(m.users, id)
	return nil
}

func TestUser_FullFlow(t *testing.T) {
	mockRepo := NewMockUserRepository()
	service := services.NewUserService(mockRepo)

	user := &domain.User{
		First_Name: "João",
		Last_Name:  "Santana",
		Email:      "joao@email.com",
		Password:   "123456",
	}

	createdUser, err := service.CreateUser(user)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	found, err := service.ViewUser(&domain.User{
		UUID: createdUser.UUID,
	})
	if err != nil {
		t.Fatalf("view failed: %v", err)
	}

	if found == nil {
		t.Fatalf("expected user, got nil")
	}

	err = service.DeleteUser(&domain.User{
		UUID: createdUser.UUID,
	})
	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	deleted, err := service.ViewUser(&domain.User{
		UUID: createdUser.UUID,
	})
	if err != nil {
		t.Fatalf("view after delete failed: %v", err)
	}

	if deleted != nil {
		t.Fatalf("expected user to be deleted")
	}
}
