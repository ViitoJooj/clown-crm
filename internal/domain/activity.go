package domain

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID              uuid.UUID              `json:"id"`
	ActivityType    string                 `json:"activity_type"` // call, email, meeting, note, status_change
	Title           string                 `json:"title"`
	Description     *string                `json:"description,omitempty"`
	DurationMinutes *int                   `json:"duration_minutes,omitempty"`

	// Related entities
	ContactID *uuid.UUID `json:"contact_id,omitempty"`
	DealID    *uuid.UUID `json:"deal_id,omitempty"`
	CompanyID *uuid.UUID `json:"company_id,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`

	// Metadata
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
}

// Validate checks if the activity has required fields
func (a *Activity) Validate() error {
	if a.Title == "" {
		return ErrInvalidInput{Field: "title", Message: "activity title is required"}
	}
	
	validTypes := map[string]bool{
		"call": true, "email": true, "meeting": true, 
		"note": true, "status_change": true, "deal_stage_change": true,
	}
	if !validTypes[a.ActivityType] {
		return ErrInvalidInput{Field: "activity_type", Message: "invalid activity type"}
	}
	
	return nil
}

type Note struct {
	ID        uuid.UUID  `json:"id"`
	Content   string     `json:"content"`
	ContactID *uuid.UUID `json:"contact_id,omitempty"`
	DealID    *uuid.UUID `json:"deal_id,omitempty"`
	CompanyID *uuid.UUID `json:"company_id,omitempty"`
	IsPinned  bool       `json:"is_pinned"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Validate checks if the note has required fields
func (n *Note) Validate() error {
	if n.Content == "" {
		return ErrInvalidInput{Field: "content", Message: "note content is required"}
	}
	if n.ContactID == nil && n.DealID == nil && n.CompanyID == nil {
		return ErrInvalidInput{Field: "entity", Message: "note must be attached to a contact, deal, or company"}
	}
	return nil
}
