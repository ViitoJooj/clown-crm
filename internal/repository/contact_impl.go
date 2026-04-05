package repository

import (
	"context"
	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
)

func (r *PostgresContactRepository) Create(ctx context.Context, contact *domain.Contact) error {
	return nil
}

func (r *PostgresContactRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Contact, error) {
	return nil, nil
}

func (r *PostgresContactRepository) Update(ctx context.Context, contact *domain.Contact) error {
	return nil
}

func (r *PostgresContactRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r *PostgresContactRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Contact, int, error) {
	return nil, 0, nil
}

func (r *PostgresContactRepository) Search(ctx context.Context, term string, page, pageSize int) ([]*domain.Contact, int, error) {
	return nil, 0, nil
}

func (r *PostgresContactRepository) GetByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*domain.Contact, error) {
	return nil, nil
}

func (r *PostgresContactRepository) GetByEmail(ctx context.Context, email string) (*domain.Contact, error) {
	return nil, nil
}

func (r *PostgresContactRepository) Count(ctx context.Context) (int64, error) {
	return 0, nil
}
