package repository

import (
	"context"
	"database/sql"

	"github.com/ViitoJooj/clown-crm/internal/domain"
)

func (r *PostgresUserRepository) ListUsers() ([]*domain.User, error) {
	rows, err := r.db.Query(context.Background(), `SELECT uuid, first_name, last_name, email, updated_at, created_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User

		err := rows.Scan(
			&user.UUID,
			&user.First_Name,
			&user.Last_Name,
			&user.Email,
			&user.Updated_at,
			&user.Created_at,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresUserRepository) FindUserByID(uuid string) (*domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(context.Background(), `SELECT uuid, first_name, last_name, email, updated_at, created_at FROM users WHERE uuid = $1`, uuid)

	err := row.Scan(
		&user.UUID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Updated_at,
		&user.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(context.Background(), `SELECT uuid, first_name, last_name, email, updated_at, created_at FROM users WHERE email = $1`, email)

	err := row.Scan(
		&user.UUID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Updated_at,
		&user.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
