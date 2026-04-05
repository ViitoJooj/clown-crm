package repository

import (
	"context"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
)

func (r *PostgresPipelineStageRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.PipelineStage, error) {
	query := `
		SELECT id, name, display_order, probability, color, is_active, 
			is_closed_won, is_closed_lost, created_at, updated_at
		FROM pipeline_stages WHERE id = $1
	`
	
	var stage domain.PipelineStage
	err := r.db.QueryRow(ctx, query, id).Scan(
		&stage.ID, &stage.Name, &stage.DisplayOrder, &stage.Probability,
		&stage.Color, &stage.IsActive, &stage.IsClosedWon, &stage.IsClosedLost,
		&stage.CreatedAt, &stage.UpdatedAt,
	)
	
	if err != nil {
		return nil, domain.ErrNotFound{Entity: "pipeline_stage", ID: id.String()}
	}
	return &stage, nil
}

func (r *PostgresPipelineStageRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.PipelineStage, int, error) {
	offset := (page - 1) * pageSize
	
	query := `
		SELECT id, name, display_order, probability, color, is_active,
			is_closed_won, is_closed_lost, created_at, updated_at
		FROM pipeline_stages
		ORDER BY display_order ASC
		LIMIT $1 OFFSET $2
	`
	
	rows, err := r.db.Query(ctx, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	
	var stages []*domain.PipelineStage
	for rows.Next() {
		var s domain.PipelineStage
		err := rows.Scan(
			&s.ID, &s.Name, &s.DisplayOrder, &s.Probability, &s.Color,
			&s.IsActive, &s.IsClosedWon, &s.IsClosedLost, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		stages = append(stages, &s)
	}
	
	var total int
	r.db.QueryRow(ctx, `SELECT COUNT(*) FROM pipeline_stages`).Scan(&total)
	
	return stages, total, nil
}

func (r *PostgresPipelineStageRepository) Create(ctx context.Context, stage *domain.PipelineStage) error {
	query := `
		INSERT INTO pipeline_stages (id, name, display_order, probability, color,
			is_active, is_closed_won, is_closed_lost, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	
	_, err := r.db.Exec(ctx, query,
		stage.ID, stage.Name, stage.DisplayOrder, stage.Probability, stage.Color,
		stage.IsActive, stage.IsClosedWon, stage.IsClosedLost,
		stage.CreatedAt, stage.UpdatedAt,
	)
	
	return err
}

func (r *PostgresPipelineStageRepository) Update(ctx context.Context, stage *domain.PipelineStage) error {
	query := `
		UPDATE pipeline_stages SET
			name = $2, display_order = $3, probability = $4, color = $5,
			is_active = $6, is_closed_won = $7, is_closed_lost = $8, updated_at = $9
		WHERE id = $1
	`
	
	_, err := r.db.Exec(ctx, query,
		stage.ID, stage.Name, stage.DisplayOrder, stage.Probability, stage.Color,
		stage.IsActive, stage.IsClosedWon, stage.IsClosedLost, stage.UpdatedAt,
	)
	
	return err
}

func (r *PostgresPipelineStageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM pipeline_stages WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
