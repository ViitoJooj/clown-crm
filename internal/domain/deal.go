package domain

import (
	"time"

	"github.com/google/uuid"
)

// PipelineStage represents a stage in the sales pipeline
type PipelineStage struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	DisplayOrder int       `json:"display_order"`
	Probability  int       `json:"probability"` // 0-100
	Color        string    `json:"color"`
	IsActive     bool      `json:"is_active"`
	IsClosedWon  bool      `json:"is_closed_won"`
	IsClosedLost bool      `json:"is_closed_lost"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// IsClosed returns true if this stage represents a closed state (won or lost)
func (p *PipelineStage) IsClosed() bool {
	return p.IsClosedWon || p.IsClosedLost
}

// Validate checks if the stage has valid values
func (p *PipelineStage) Validate() error {
	if p.Name == "" {
		return ErrInvalidInput{Field: "name", Message: "stage name is required"}
	}
	if p.Probability < 0 || p.Probability > 100 {
		return ErrInvalidInput{Field: "probability", Message: "probability must be between 0 and 100"}
	}
	return nil
}

type Deal struct {
	ID                uuid.UUID  `json:"id"`
	Title             string     `json:"title"`
	Value             float64    `json:"value"`
	Currency          string     `json:"currency"`
	ExpectedCloseDate *time.Time `json:"expected_close_date,omitempty"`

	// Stage and probability
	StageID     uuid.UUID `json:"stage_id"`
	Probability int       `json:"probability"` // 0-100

	// Related entities
	ContactID *uuid.UUID `json:"contact_id,omitempty"`
	CompanyID *uuid.UUID `json:"company_id,omitempty"`

	// Assignment
	AssignedTo *uuid.UUID `json:"assigned_to,omitempty"`
	OwnerID    *uuid.UUID `json:"owner_id,omitempty"`

	// Source and metadata
	Source       *string                `json:"source,omitempty"`
	Notes        *string                `json:"notes,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags         []string               `json:"tags,omitempty"`

	// Status
	IsWon      bool    `json:"is_won"`
	IsLost     bool    `json:"is_lost"`
	LostReason *string `json:"lost_reason,omitempty"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
}

// Validate checks if the deal has required fields
func (d *Deal) Validate() error {
	if d.Title == "" {
		return ErrInvalidInput{Field: "title", Message: "deal title is required"}
	}
	if d.Value < 0 {
		return ErrInvalidInput{Field: "value", Message: "deal value cannot be negative"}
	}
	if d.Probability < 0 || d.Probability > 100 {
		return ErrInvalidInput{Field: "probability", Message: "probability must be between 0 and 100"}
	}
	if d.IsWon && d.IsLost {
		return ErrInvalidInput{Field: "status", Message: "deal cannot be both won and lost"}
	}
	return nil
}

// IsClosed returns true if the deal is won or lost
func (d *Deal) IsClosed() bool {
	return d.IsWon || d.IsLost
}

// IsActive returns true if the deal is not closed
func (d *Deal) IsActive() bool {
	return !d.IsClosed()
}

// MarkWon marks the deal as won
func (d *Deal) MarkWon() {
	d.IsWon = true
	d.IsLost = false
	now := time.Now()
	d.ClosedAt = &now
	d.Probability = 100
}

// MarkLost marks the deal as lost
func (d *Deal) MarkLost(reason string) {
	d.IsWon = false
	d.IsLost = true
	d.LostReason = &reason
	now := time.Now()
	d.ClosedAt = &now
	d.Probability = 0
}

// Reopen reopens a closed deal
func (d *Deal) Reopen() {
	d.IsWon = false
	d.IsLost = false
	d.ClosedAt = nil
	d.LostReason = nil
}
