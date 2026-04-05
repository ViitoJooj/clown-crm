package repository

import (
	"context"
	"errors"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *PostgresNotificationRepository) Create(ctx context.Context, notification *domain.Notification) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO notifications (
			id, user_id, notification_type, title, message,
			related_to_type, related_to_id, action_url,
			is_read, read_at, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		notification.ID,
		notification.UserID,
		notification.NotificationType,
		notification.Title,
		notification.Message,
		notification.RelatedToType,
		notification.RelatedToID,
		notification.ActionURL,
		notification.IsRead,
		notification.ReadAt,
		notification.CreatedAt,
	)

	return err
}

func (r *PostgresNotificationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Notification, error) {
	var notification domain.Notification

	row := r.db.QueryRow(ctx,
		`SELECT 
			id, user_id, notification_type, title, message,
			related_to_type, related_to_id, action_url,
			is_read, read_at, created_at
		FROM notifications WHERE id = $1`,
		id,
	)

	err := row.Scan(
		&notification.ID,
		&notification.UserID,
		&notification.NotificationType,
		&notification.Title,
		&notification.Message,
		&notification.RelatedToType,
		&notification.RelatedToID,
		&notification.ActionURL,
		&notification.IsRead,
		&notification.ReadAt,
		&notification.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound{Entity: "notification", ID: id.String()}
		}
		return nil, err
	}

	return &notification, nil
}

func (r *PostgresNotificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	result, err := r.db.Exec(ctx,
		`UPDATE notifications 
		SET user_id = $1, notification_type = $2, title = $3, message = $4,
			related_to_type = $5, related_to_id = $6, action_url = $7,
			is_read = $8, read_at = $9
		WHERE id = $10`,
		notification.UserID,
		notification.NotificationType,
		notification.Title,
		notification.Message,
		notification.RelatedToType,
		notification.RelatedToID,
		notification.ActionURL,
		notification.IsRead,
		notification.ReadAt,
		notification.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "notification", ID: notification.ID.String()}
	}

	return nil
}

func (r *PostgresNotificationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx, "DELETE FROM notifications WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "notification", ID: id.String()}
	}

	return nil
}

func (r *PostgresNotificationRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Notification, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, user_id, notification_type, title, message,
			related_to_type, related_to_id, action_url,
			is_read, read_at, created_at
		FROM notifications
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := make([]*domain.Notification, 0)

	for rows.Next() {
		var notification domain.Notification

		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.NotificationType,
			&notification.Title,
			&notification.Message,
			&notification.RelatedToType,
			&notification.RelatedToID,
			&notification.ActionURL,
			&notification.IsRead,
			&notification.ReadAt,
			&notification.CreatedAt,
		); err != nil {
			return nil, err
		}

		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *PostgresNotificationRepository) GetUnreadByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Notification, error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, user_id, notification_type, title, message,
			related_to_type, related_to_id, action_url,
			is_read, read_at, created_at
		FROM notifications
		WHERE user_id = $1 AND is_read = false
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := make([]*domain.Notification, 0)

	for rows.Next() {
		var notification domain.Notification

		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.NotificationType,
			&notification.Title,
			&notification.Message,
			&notification.RelatedToType,
			&notification.RelatedToID,
			&notification.ActionURL,
			&notification.IsRead,
			&notification.ReadAt,
			&notification.CreatedAt,
		); err != nil {
			return nil, err
		}

		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *PostgresNotificationRepository) MarkAsRead(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx,
		`UPDATE notifications SET is_read = true, read_at = NOW() WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound{Entity: "notification", ID: id.String()}
	}

	return nil
}

func (r *PostgresNotificationRepository) MarkAllAsRead(ctx context.Context, userID uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`UPDATE notifications SET is_read = true, read_at = NOW() 
		 WHERE user_id = $1 AND is_read = false`,
		userID,
	)

	return err
}

func (r *PostgresNotificationRepository) CountUnreadByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND is_read = false`,
		userID,
	).Scan(&count)

	return count, err
}

func (r *PostgresNotificationRepository) DeleteOldNotifications(ctx context.Context, userID uuid.UUID, olderThanDays int) error {
	_, err := r.db.Exec(ctx,
		`DELETE FROM notifications 
		 WHERE user_id = $1 
		   AND is_read = true 
		   AND created_at < NOW() - INTERVAL '1 day' * $2`,
		userID, olderThanDays,
	)

	return err
}

func (r *PostgresNotificationRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM notifications").Scan(&count)
	return count, err
}
