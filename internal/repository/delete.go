package repository

import (
	"context"
	"errors"
	"log"
)

func (r *PostgresUserRepository) DeleteUserById(id string) error {
	res, err := r.db.Exec(
		context.Background(),
		"DELETE FROM users WHERE uuid = $1",
		id,
	)
	if err != nil {
		log.Println(err)
		return errors.New("internal error")
	}

	if res.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}
