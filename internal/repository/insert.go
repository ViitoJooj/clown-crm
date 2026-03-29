package repository

import (
	"context"

	"github.com/ViitoJooj/clown-crm/internal/domain"
)

func (r *PostgresUserRepository) CreateUser(user *domain.User) error {
	_, err := r.db.Exec(context.Background(),
		`INSERT INTO users (uuid, first_name, last_name, email, password, updated_at, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		user.UUID,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Password,
		user.Updated_at,
		user.Created_at,
	)

	return err
}

func (r *PostgresChatRepository) SaveMessage(msg domain.Chat) error {
	_, err := r.db.Exec(context.Background(),
		`INSERT INTO messages (sender, receiver, message, created_at)
		 VALUES ($1, $2, $3, $4)`,
		msg.From,
		msg.To,
		msg.Message,
		msg.Time,
	)

	return err
}
