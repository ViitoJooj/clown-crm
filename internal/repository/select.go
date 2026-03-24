package repository

import (
	"context"
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (r *PostgresUserRepository) ListUsers() ([]*domain.User, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT uuid, first_name, last_name, email, updated_at, created_at FROM users`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)

	for rows.Next() {
		var user domain.User

		if err := rows.Scan(
			&user.UUID,
			&user.First_Name,
			&user.Last_Name,
			&user.Email,
			&user.Updated_at,
			&user.Created_at,
		); err != nil {
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

	row := r.db.QueryRow(context.Background(),
		`SELECT uuid, first_name, last_name, email, updated_at, created_at FROM users WHERE uuid = $1`,
		uuid,
	)

	err := row.Scan(
		&user.UUID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Updated_at,
		&user.Created_at,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(context.Background(),
		`SELECT uuid, first_name, last_name, email, updated_at, created_at FROM users WHERE email = $1`,
		email,
	)

	err := row.Scan(
		&user.UUID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Updated_at,
		&user.Created_at,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
