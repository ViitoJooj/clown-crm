package repository

import (
	"context"
	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
)

func (r *PostgresActivityRepository) Create(ctx context.Context, activity *domain.Activity) error {
	return nil
}

func (r *PostgresActivityRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Activity, error) {
	return nil, nil
}

func (r *PostgresActivityRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r *PostgresActivityRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Activity, int, error) {
	return nil, 0, nil
}

func (r *PostgresActivityRepository) GetByContactID(ctx context.Context, contactID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error) {
	return nil, 0, nil
}

func (r *PostgresActivityRepository) GetByDealID(ctx context.Context, dealID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error) {
	return nil, 0, nil
}

func (r *PostgresActivityRepository) GetByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error) {
	return nil, 0, nil
}

func (r *PostgresActivityRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error) {
	return nil, 0, nil
}

func (r *PostgresActivityRepository) GetByType(ctx context.Context, activityType string, limit, offset int) ([]*domain.Activity, int, error) {
	return nil, 0, nil
}

func (r *PostgresActivityRepository) Count(ctx context.Context) (int64, error) {
	return 0, nil
}
