package repository

import (
	"context"
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *PostgresTagRepository) Create(ctx context.Context, tag *domain.Tag) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO tags (
			id, name, color, entity_type, usage_count, created_at
		) VALUES ($1, $2, $3, $4, $5, $6)`,
		tag.ID,
		tag.Name,
		tag.Color,
		tag.EntityType,
		tag.UsageCount,
		tag.CreatedAt,
	)

	return err
}

func (r *PostgresTagRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error) {
	var tag domain.Tag

	row := r.db.QueryRow(ctx,
		`SELECT id, name, color, entity_type, usage_count, created_at
		FROM tags WHERE id = $1`,
		id,
	)

	err := row.Scan(
		&tag.ID,
		&tag.Name,
		&tag.Color,
		&tag.EntityType,
		&tag.UsageCount,
		&tag.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound{Entity: "tag", ID: id.String()}
		}
		return nil, err
	}

	return &tag, nil
}

func (r *PostgresTagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	result, err := r.db.Exec(ctx,
		`UPDATE tags 
		SET name = $1, color = $2, entity_type = $3, usage_count = $4
		WHERE id = $5`,
		tag.Name,
		tag.Color,
		tag.EntityType,
		tag.UsageCount,
		tag.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "tag", ID: tag.ID.String()}
	}

	return nil
}

func (r *PostgresTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx, "DELETE FROM tags WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "tag", ID: id.String()}
	}

	return nil
}

func (r *PostgresTagRepository) List(ctx context.Context, limit, offset int) ([]*domain.Tag, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, color, entity_type, usage_count, created_at
		FROM tags
		ORDER BY name ASC
		LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]*domain.Tag, 0)

	for rows.Next() {
		var tag domain.Tag

		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Color,
			&tag.EntityType,
			&tag.UsageCount,
			&tag.CreatedAt,
		); err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *PostgresTagRepository) GetByEntityType(ctx context.Context, entityType string, limit, offset int) ([]*domain.Tag, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, color, entity_type, usage_count, created_at
		FROM tags
		WHERE entity_type = $1 OR entity_type = 'all'
		ORDER BY usage_count DESC, name ASC
		LIMIT $2 OFFSET $3`,
		entityType, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]*domain.Tag, 0)

	for rows.Next() {
		var tag domain.Tag

		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Color,
			&tag.EntityType,
			&tag.UsageCount,
			&tag.CreatedAt,
		); err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *PostgresTagRepository) GetByName(ctx context.Context, name, entityType string) (*domain.Tag, error) {
	var tag domain.Tag

	row := r.db.QueryRow(ctx,
		`SELECT id, name, color, entity_type, usage_count, created_at
		FROM tags WHERE name = $1 AND entity_type = $2`,
		name, entityType,
	)

	err := row.Scan(
		&tag.ID,
		&tag.Name,
		&tag.Color,
		&tag.EntityType,
		&tag.UsageCount,
		&tag.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound{Entity: "tag", ID: name}
		}
		return nil, err
	}

	return &tag, nil
}

func (r *PostgresTagRepository) SearchByName(ctx context.Context, query string, limit, offset int) ([]*domain.Tag, error) {
	searchPattern := "%" + query + "%"

	rows, err := r.db.Query(ctx,
		`SELECT id, name, color, entity_type, usage_count, created_at
		FROM tags
		WHERE name ILIKE $1
		ORDER BY usage_count DESC, name ASC
		LIMIT $2 OFFSET $3`,
		searchPattern, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]*domain.Tag, 0)

	for rows.Next() {
		var tag domain.Tag

		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Color,
			&tag.EntityType,
			&tag.UsageCount,
			&tag.CreatedAt,
		); err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *PostgresTagRepository) IncrementUsage(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx,
		`UPDATE tags SET usage_count = usage_count + 1 WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "tag", ID: id.String()}
	}

	return nil
}

func (r *PostgresTagRepository) DecrementUsage(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx,
		`UPDATE tags SET usage_count = GREATEST(usage_count - 1, 0) WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "tag", ID: id.String()}
	}

	return nil
}

func (r *PostgresTagRepository) GetMostUsed(ctx context.Context, limit int) ([]*domain.Tag, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, color, entity_type, usage_count, created_at
		FROM tags
		WHERE usage_count > 0
		ORDER BY usage_count DESC
		LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]*domain.Tag, 0)

	for rows.Next() {
		var tag domain.Tag

		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Color,
			&tag.EntityType,
			&tag.UsageCount,
			&tag.CreatedAt,
		); err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *PostgresTagRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM tags").Scan(&count)
	return count, err
}
