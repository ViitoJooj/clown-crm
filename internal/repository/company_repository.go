package repository

import (
	"context"
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *PostgresCompanyRepository) Create(ctx context.Context, company *domain.Company) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO companies (
			id, name, industry, company_size, website, phone,
			address, city, state, country, postal_code,
			status, tags, notes, custom_fields, owner_id,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`,
		company.ID,
		company.Name,
		company.Industry,
		company.CompanySize,
		company.Website,
		company.Phone,
		company.Address,
		company.City,
		company.State,
		company.Country,
		company.PostalCode,
		company.Status,
		company.Tags,
		company.Notes,
		company.CustomFields,
		company.OwnerID,
		company.CreatedAt,
		company.UpdatedAt,
	)

	return err
}

func (r *PostgresCompanyRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Company, error) {
	var company domain.Company

	row := r.db.QueryRow(ctx,
		`SELECT 
			id, name, industry, company_size, website, phone,
			address, city, state, country, postal_code,
			status, tags, notes, custom_fields, owner_id,
			created_at, updated_at
		FROM companies WHERE id = $1`,
		id,
	)

	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Industry,
		&company.CompanySize,
		&company.Website,
		&company.Phone,
		&company.Address,
		&company.City,
		&company.State,
		&company.Country,
		&company.PostalCode,
		&company.Status,
		&company.Tags,
		&company.Notes,
		&company.CustomFields,
		&company.OwnerID,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound{Entity: "company", ID: id.String()}
		}
		return nil, err
	}

	return &company, nil
}

func (r *PostgresCompanyRepository) Update(ctx context.Context, company *domain.Company) error {
	result, err := r.db.Exec(ctx,
		`UPDATE companies 
		SET name = $1, industry = $2, company_size = $3, website = $4, phone = $5,
			address = $6, city = $7, state = $8, country = $9, postal_code = $10,
			status = $11, tags = $12, notes = $13, custom_fields = $14, 
			owner_id = $15, updated_at = $16
		WHERE id = $17`,
		company.Name,
		company.Industry,
		company.CompanySize,
		company.Website,
		company.Phone,
		company.Address,
		company.City,
		company.State,
		company.Country,
		company.PostalCode,
		company.Status,
		company.Tags,
		company.Notes,
		company.CustomFields,
		company.OwnerID,
		company.UpdatedAt,
		company.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "company", ID: company.ID.String()}
	}

	return nil
}

func (r *PostgresCompanyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx, "DELETE FROM companies WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "company", ID: id.String()}
	}

	return nil
}

func (r *PostgresCompanyRepository) List(ctx context.Context, limit, offset int) ([]*domain.Company, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, name, industry, company_size, website, phone,
			address, city, state, country, postal_code,
			status, tags, notes, custom_fields, owner_id,
			created_at, updated_at
		FROM companies
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]*domain.Company, 0)

	for rows.Next() {
		var company domain.Company

		if err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Industry,
			&company.CompanySize,
			&company.Website,
			&company.Phone,
			&company.Address,
			&company.City,
			&company.State,
			&company.Country,
			&company.PostalCode,
			&company.Status,
			&company.Tags,
			&company.Notes,
			&company.CustomFields,
			&company.OwnerID,
			&company.CreatedAt,
			&company.UpdatedAt,
		); err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *PostgresCompanyRepository) SearchByName(ctx context.Context, query string, limit, offset int) ([]*domain.Company, error) {
	searchPattern := "%" + query + "%"

	rows, err := r.db.Query(ctx,
		`SELECT 
			id, name, industry, company_size, website, phone,
			address, city, state, country, postal_code,
			status, tags, notes, custom_fields, owner_id,
			created_at, updated_at
		FROM companies
		WHERE name ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
		searchPattern, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]*domain.Company, 0)

	for rows.Next() {
		var company domain.Company

		if err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Industry,
			&company.CompanySize,
			&company.Website,
			&company.Phone,
			&company.Address,
			&company.City,
			&company.State,
			&company.Country,
			&company.PostalCode,
			&company.Status,
			&company.Tags,
			&company.Notes,
			&company.CustomFields,
			&company.OwnerID,
			&company.CreatedAt,
			&company.UpdatedAt,
		); err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *PostgresCompanyRepository) FilterByStatus(ctx context.Context, status string, limit, offset int) ([]*domain.Company, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, name, industry, company_size, website, phone,
			address, city, state, country, postal_code,
			status, tags, notes, custom_fields, owner_id,
			created_at, updated_at
		FROM companies
		WHERE status = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
		status, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]*domain.Company, 0)

	for rows.Next() {
		var company domain.Company

		if err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Industry,
			&company.CompanySize,
			&company.Website,
			&company.Phone,
			&company.Address,
			&company.City,
			&company.State,
			&company.Country,
			&company.PostalCode,
			&company.Status,
			&company.Tags,
			&company.Notes,
			&company.CustomFields,
			&company.OwnerID,
			&company.CreatedAt,
			&company.UpdatedAt,
		); err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *PostgresCompanyRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM companies").Scan(&count)
	return count, err
}
