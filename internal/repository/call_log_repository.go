package repository

import (
	"context"
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *PostgresCallLogRepository) Create(ctx context.Context, callLog *domain.CallLog) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO call_logs (
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
		callLog.ID,
		callLog.CallType,
		callLog.ContactID,
		callLog.UserID,
		callLog.DurationSeconds,
		callLog.Status,
		callLog.TwilioCallSID,
		callLog.FromNumber,
		callLog.ToNumber,
		callLog.RecordingURL,
		callLog.RecordingDuration,
		callLog.Notes,
		callLog.Metadata,
		callLog.StartedAt,
		callLog.EndedAt,
		callLog.CreatedAt,
	)

	return err
}

func (r *PostgresCallLogRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.CallLog, error) {
	var callLog domain.CallLog

	row := r.db.QueryRow(ctx,
		`SELECT 
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		FROM call_logs WHERE id = $1`,
		id,
	)

	err := row.Scan(
		&callLog.ID,
		&callLog.CallType,
		&callLog.ContactID,
		&callLog.UserID,
		&callLog.DurationSeconds,
		&callLog.Status,
		&callLog.TwilioCallSID,
		&callLog.FromNumber,
		&callLog.ToNumber,
		&callLog.RecordingURL,
		&callLog.RecordingDuration,
		&callLog.Notes,
		&callLog.Metadata,
		&callLog.StartedAt,
		&callLog.EndedAt,
		&callLog.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound{Entity: "call_log", ID: id.String()}
		}
		return nil, err
	}

	return &callLog, nil
}

func (r *PostgresCallLogRepository) Update(ctx context.Context, callLog *domain.CallLog) error {
	result, err := r.db.Exec(ctx,
		`UPDATE call_logs 
		SET call_type = $1, contact_id = $2, user_id = $3,
			duration_seconds = $4, status = $5,
			twilio_call_sid = $6, from_number = $7, to_number = $8,
			recording_url = $9, recording_duration = $10,
			notes = $11, metadata = $12,
			started_at = $13, ended_at = $14
		WHERE id = $15`,
		callLog.CallType,
		callLog.ContactID,
		callLog.UserID,
		callLog.DurationSeconds,
		callLog.Status,
		callLog.TwilioCallSID,
		callLog.FromNumber,
		callLog.ToNumber,
		callLog.RecordingURL,
		callLog.RecordingDuration,
		callLog.Notes,
		callLog.Metadata,
		callLog.StartedAt,
		callLog.EndedAt,
		callLog.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "call_log", ID: callLog.ID.String()}
	}

	return nil
}

func (r *PostgresCallLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx, "DELETE FROM call_logs WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "call_log", ID: id.String()}
	}

	return nil
}

func (r *PostgresCallLogRepository) List(ctx context.Context, limit, offset int) ([]*domain.CallLog, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		FROM call_logs
		ORDER BY started_at DESC
		LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	callLogs := make([]*domain.CallLog, 0)

	for rows.Next() {
		var callLog domain.CallLog

		if err := rows.Scan(
			&callLog.ID,
			&callLog.CallType,
			&callLog.ContactID,
			&callLog.UserID,
			&callLog.DurationSeconds,
			&callLog.Status,
			&callLog.TwilioCallSID,
			&callLog.FromNumber,
			&callLog.ToNumber,
			&callLog.RecordingURL,
			&callLog.RecordingDuration,
			&callLog.Notes,
			&callLog.Metadata,
			&callLog.StartedAt,
			&callLog.EndedAt,
			&callLog.CreatedAt,
		); err != nil {
			return nil, err
		}

		callLogs = append(callLogs, &callLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return callLogs, nil
}

func (r *PostgresCallLogRepository) GetByContactID(ctx context.Context, contactID uuid.UUID, limit, offset int) ([]*domain.CallLog, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		FROM call_logs
		WHERE contact_id = $1
		ORDER BY started_at DESC
		LIMIT $2 OFFSET $3`,
		contactID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	callLogs := make([]*domain.CallLog, 0)

	for rows.Next() {
		var callLog domain.CallLog

		if err := rows.Scan(
			&callLog.ID,
			&callLog.CallType,
			&callLog.ContactID,
			&callLog.UserID,
			&callLog.DurationSeconds,
			&callLog.Status,
			&callLog.TwilioCallSID,
			&callLog.FromNumber,
			&callLog.ToNumber,
			&callLog.RecordingURL,
			&callLog.RecordingDuration,
			&callLog.Notes,
			&callLog.Metadata,
			&callLog.StartedAt,
			&callLog.EndedAt,
			&callLog.CreatedAt,
		); err != nil {
			return nil, err
		}

		callLogs = append(callLogs, &callLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return callLogs, nil
}

func (r *PostgresCallLogRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.CallLog, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		FROM call_logs
		WHERE user_id = $1
		ORDER BY started_at DESC
		LIMIT $2 OFFSET $3`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	callLogs := make([]*domain.CallLog, 0)

	for rows.Next() {
		var callLog domain.CallLog

		if err := rows.Scan(
			&callLog.ID,
			&callLog.CallType,
			&callLog.ContactID,
			&callLog.UserID,
			&callLog.DurationSeconds,
			&callLog.Status,
			&callLog.TwilioCallSID,
			&callLog.FromNumber,
			&callLog.ToNumber,
			&callLog.RecordingURL,
			&callLog.RecordingDuration,
			&callLog.Notes,
			&callLog.Metadata,
			&callLog.StartedAt,
			&callLog.EndedAt,
			&callLog.CreatedAt,
		); err != nil {
			return nil, err
		}

		callLogs = append(callLogs, &callLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return callLogs, nil
}

func (r *PostgresCallLogRepository) GetByStatus(ctx context.Context, status string, limit, offset int) ([]*domain.CallLog, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		FROM call_logs
		WHERE status = $1
		ORDER BY started_at DESC
		LIMIT $2 OFFSET $3`,
		status, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	callLogs := make([]*domain.CallLog, 0)

	for rows.Next() {
		var callLog domain.CallLog

		if err := rows.Scan(
			&callLog.ID,
			&callLog.CallType,
			&callLog.ContactID,
			&callLog.UserID,
			&callLog.DurationSeconds,
			&callLog.Status,
			&callLog.TwilioCallSID,
			&callLog.FromNumber,
			&callLog.ToNumber,
			&callLog.RecordingURL,
			&callLog.RecordingDuration,
			&callLog.Notes,
			&callLog.Metadata,
			&callLog.StartedAt,
			&callLog.EndedAt,
			&callLog.CreatedAt,
		); err != nil {
			return nil, err
		}

		callLogs = append(callLogs, &callLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return callLogs, nil
}

func (r *PostgresCallLogRepository) GetByTwilioCallSID(ctx context.Context, twilioCallSID string) (*domain.CallLog, error) {
	var callLog domain.CallLog

	row := r.db.QueryRow(ctx,
		`SELECT 
			id, call_type, contact_id, user_id,
			duration_seconds, status,
			twilio_call_sid, from_number, to_number,
			recording_url, recording_duration,
			notes, metadata,
			started_at, ended_at, created_at
		FROM call_logs WHERE twilio_call_sid = $1`,
		twilioCallSID,
	)

	err := row.Scan(
		&callLog.ID,
		&callLog.CallType,
		&callLog.ContactID,
		&callLog.UserID,
		&callLog.DurationSeconds,
		&callLog.Status,
		&callLog.TwilioCallSID,
		&callLog.FromNumber,
		&callLog.ToNumber,
		&callLog.RecordingURL,
		&callLog.RecordingDuration,
		&callLog.Notes,
		&callLog.Metadata,
		&callLog.StartedAt,
		&callLog.EndedAt,
		&callLog.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound{Entity: "call_log", ID: twilioCallSID}
		}
		return nil, err
	}

	return &callLog, nil
}

func (r *PostgresCallLogRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM call_logs").Scan(&count)
	return count, err
}
