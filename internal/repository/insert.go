package repository

import (
	"context"
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
)

func (r *PostgresUserRepository) CreateUser(user *domain.User) error {
	err := r.db.QueryRow(context.Background(), "INSERT INTO users (uuid, first_name, last_name, email, password, updated_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.UUID, user.First_Name, user.Last_Name, user.Email, user.Password, user.Updated_at, user.Created_at,
	)
	if err != nil {
		return errors.New("internal error")
	}
	return nil
}
