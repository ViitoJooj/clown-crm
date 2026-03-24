package repository

import (
	"context"
	"errors"
	"log"

	"github.com/ViitoJooj/clown-crm/internal/domain"
)

func (r *PostgresUserRepository) UpdateUser(user *domain.User) error {
	_, err := r.db.Exec(
		context.Background(),
		`UPDATE users 
		 SET first_name = $1,
		     last_name = $2,
		     email = $3,
		     password = $4,
		     updated_at = $5
		 WHERE uuid = $6`,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Password,
		user.Updated_at,
		user.UUID,
	)

	if err != nil {
		log.Println(err)
		return errors.New("internal error")
	}

	return nil
}
