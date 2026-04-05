package dtos

import (
	"time"

	"github.com/google/uuid"
)

// InputTaskDTO represents the input for creating a task
type InputTaskDTO struct {
	Title         string     `json:"title" validate:"required,min=1,max=255"`
	Description   *string    `json:"description,omitempty"`
	TaskType      string     `json:"task_type" validate:"required,oneof=call email meeting todo follow_up"`
	Priority      string     `json:"priority" validate:"required,oneof=low medium high urgent"`
	Status        string     `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress completed cancelled"`
	DueDate       *time.Time `json:"due_date,omitempty"`
	ReminderAt    *time.Time `json:"reminder_at,omitempty"`
	AssignedTo    *uuid.UUID `json:"assigned_to,omitempty"`
	RelatedToType *string    `json:"related_to_type,omitempty" validate:"omitempty,oneof=contact deal company"`
	RelatedToID   *uuid.UUID `json:"related_to_id,omitempty"`
}

// UpdateTaskDTO represents the input for updating a task
type UpdateTaskDTO struct {
	Title         *string    `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description   *string    `json:"description,omitempty"`
	TaskType      *string    `json:"task_type,omitempty" validate:"omitempty,oneof=call email meeting todo follow_up"`
	Priority      *string    `json:"priority,omitempty" validate:"omitempty,oneof=low medium high urgent"`
	Status        *string    `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress completed cancelled"`
	DueDate       *time.Time `json:"due_date,omitempty"`
	ReminderAt    *time.Time `json:"reminder_at,omitempty"`
	AssignedTo    *uuid.UUID `json:"assigned_to,omitempty"`
	RelatedToType *string    `json:"related_to_type,omitempty" validate:"omitempty,oneof=contact deal company"`
	RelatedToID   *uuid.UUID `json:"related_to_id,omitempty"`
}

// OutputTaskDTO represents the output for a task
type OutputTaskDTO struct {
	ID            uuid.UUID  `json:"id"`
	Title         string     `json:"title"`
	Description   *string    `json:"description,omitempty"`
	TaskType      string     `json:"task_type"`
	Priority      string     `json:"priority"`
	Status        string     `json:"status"`
	DueDate       *time.Time `json:"due_date,omitempty"`
	ReminderAt    *time.Time `json:"reminder_at,omitempty"`
	AssignedTo    *uuid.UUID `json:"assigned_to,omitempty"`
	CreatedBy     *uuid.UUID `json:"created_by,omitempty"`
	RelatedToType *string    `json:"related_to_type,omitempty"`
	RelatedToID   *uuid.UUID `json:"related_to_id,omitempty"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// TaskListResponse represents the response for listing tasks
type TaskListResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Tasks   []OutputTaskDTO `json:"tasks"`
	Page    int             `json:"page"`
	PerPage int             `json:"per_page"`
	Total   int             `json:"total"`
}

// TaskResponse represents the response for a single task operation
type TaskResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Task    OutputTaskDTO `json:"task"`
}

// AssignTaskDTO represents the input for assigning a task
type AssignTaskDTO struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// SetReminderDTO represents the input for setting a reminder
type SetReminderDTO struct {
	ReminderAt time.Time `json:"reminder_at" validate:"required"`
}

// TaskStatsResponse represents task statistics
type TaskStatsResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Stats   map[string]int `json:"stats"`
}
