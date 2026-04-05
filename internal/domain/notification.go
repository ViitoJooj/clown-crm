package domain

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID               uuid.UUID  `json:"id"`
	UserID           uuid.UUID  `json:"user_id"`
	NotificationType string     `json:"notification_type"` // task_due, deal_stage, mention, assignment, call_missed
	Title            string     `json:"title"`
	Message          *string    `json:"message,omitempty"`

	// Related entity (polymorphic)
	RelatedToType *string    `json:"related_to_type,omitempty"` // contact, deal, task, call
	RelatedToID   *uuid.UUID `json:"related_to_id,omitempty"`
	ActionURL     *string    `json:"action_url,omitempty"`

	// Read status
	IsRead   bool       `json:"is_read"`
	ReadAt   *time.Time `json:"read_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// Validate checks if the notification has required fields
func (n *Notification) Validate() error {
	if n.Title == "" {
		return ErrInvalidInput{Field: "title", Message: "notification title is required"}
	}
	
	validTypes := map[string]bool{
		"task_due": true, "deal_stage": true, "mention": true, 
		"assignment": true, "call_missed": true, "task_assigned": true,
	}
	if !validTypes[n.NotificationType] {
		return ErrInvalidInput{Field: "notification_type", Message: "invalid notification type"}
	}
	
	return nil
}

// MarkAsRead marks the notification as read
func (n *Notification) MarkAsRead() {
	n.IsRead = true
	now := time.Now()
	n.ReadAt = &now
}

type Tag struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	EntityType string    `json:"entity_type"` // contact, deal, company, all
	UsageCount int       `json:"usage_count"`
	CreatedAt  time.Time `json:"created_at"`
}

// Validate checks if the tag has required fields
func (t *Tag) Validate() error {
	if t.Name == "" {
		return ErrInvalidInput{Field: "name", Message: "tag name is required"}
	}
	
	validEntityTypes := map[string]bool{"contact": true, "deal": true, "company": true, "all": true}
	if !validEntityTypes[t.EntityType] {
		return ErrInvalidInput{Field: "entity_type", Message: "invalid entity type"}
	}
	
	return nil
}

// IncrementUsage increments the tag usage count
func (t *Tag) IncrementUsage() {
	t.UsageCount++
}

// DecrementUsage decrements the tag usage count
func (t *Tag) DecrementUsage() {
	if t.UsageCount > 0 {
		t.UsageCount--
	}
}
