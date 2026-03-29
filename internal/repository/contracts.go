package repository

import (
	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindUserByID(id string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUserById(id string) error
}

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) UserRepository {
	return &PostgresUserRepository{db: db}
}

type ChatRepository interface {
	SaveMessage(msg domain.Chat) error
	ListMessages(userA, userB string) ([]domain.Chat, error)
}

type PostgresChatRepository struct {
	db *pgxpool.Pool
}

func NewPostgresChatRepository(db *pgxpool.Pool) ChatRepository {
	return &PostgresChatRepository{db: db}
}
