package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	TaskType    string     `json:"task_type"` // call, email, meeting, todo, follow_up
	Priority    string     `json:"priority"`  // low, medium, high, urgent
	Status      string     `json:"status"`    // pending, in_progress, completed, cancelled
	DueDate     *time.Time `json:"due_date,omitempty"`
	ReminderAt  *time.Time `json:"reminder_at,omitempty"`

	// Assignment
	AssignedTo *uuid.UUID `json:"assigned_to,omitempty"`
	CreatedBy  *uuid.UUID `json:"created_by,omitempty"`

	// Related entity (polymorphic)
	RelatedToType *string    `json:"related_to_type,omitempty"` // contact, deal, company
	RelatedToID   *uuid.UUID `json:"related_to_id,omitempty"`

	// Timestamps
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Validate checks if the task has required fields
func (t *Task) Validate() error {
	if t.Title == "" {
		return ErrInvalidInput{Field: "title", Message: "task title is required"}
	}
	
	validTypes := map[string]bool{"call": true, "email": true, "meeting": true, "todo": true, "follow_up": true}
	if !validTypes[t.TaskType] {
		return ErrInvalidInput{Field: "task_type", Message: "invalid task type"}
	}
	
	validPriorities := map[string]bool{"low": true, "medium": true, "high": true, "urgent": true}
	if !validPriorities[t.Priority] {
		return ErrInvalidInput{Field: "priority", Message: "invalid priority"}
	}
	
	validStatuses := map[string]bool{"pending": true, "in_progress": true, "completed": true, "cancelled": true}
	if !validStatuses[t.Status] {
		return ErrInvalidInput{Field: "status", Message: "invalid status"}
	}
	
	return nil
}

// IsCompleted returns true if the task is completed
func (t *Task) IsCompleted() bool {
	return t.Status == "completed"
}

// IsOverdue returns true if the task is past due
func (t *Task) IsOverdue() bool {
	if t.DueDate == nil || t.IsCompleted() {
		return false
	}
	return t.DueDate.Before(time.Now())
}

// Complete marks the task as completed
func (t *Task) Complete() {
	t.Status = "completed"
	now := time.Now()
	t.CompletedAt = &now
}

// Cancel marks the task as cancelled
func (t *Task) Cancel() {
	t.Status = "cancelled"
}
