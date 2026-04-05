package repository

import (
	"context"
	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
)

func (r *PostgresDealRepository) Create(ctx context.Context, deal *domain.Deal) error {
	return nil
}

func (r *PostgresDealRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Deal, error) {
	return nil, nil
}

func (r *PostgresDealRepository) Update(ctx context.Context, deal *domain.Deal) error {
	return nil
}

func (r *PostgresDealRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r *PostgresDealRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Deal, int, error) {
	return nil, 0, nil
}

func (r *PostgresDealRepository) GetByStageID(ctx context.Context, stageID uuid.UUID, limit, offset int) ([]*domain.Deal, error) {
	return nil, nil
}

func (r *PostgresDealRepository) UpdateStage(ctx context.Context, dealID, stageID uuid.UUID) error {
	return nil
}

func (r *PostgresDealRepository) GetActiveDeals(ctx context.Context, limit, offset int) ([]*domain.Deal, error) {
	return nil, nil
}

func (r *PostgresDealRepository) GetWonDeals(ctx context.Context, limit, offset int) ([]*domain.Deal, error) {
	return nil, nil
}

func (r *PostgresDealRepository) GetLostDeals(ctx context.Context, limit, offset int) ([]*domain.Deal, error) {
	return nil, nil
}

func (r *PostgresDealRepository) GetByContactID(ctx context.Context, contactID uuid.UUID, limit, offset int) ([]*domain.Deal, error) {
	return nil, nil
}

func (r *PostgresDealRepository) Count(ctx context.Context) (int64, error) {
	return 0, nil
}
