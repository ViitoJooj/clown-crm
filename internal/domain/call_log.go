package domain

import (
	"time"

	"github.com/google/uuid"
)

// CallLog represents a phone or video call record
type CallLog struct {
	ID        uuid.UUID  `json:"id"`
	CallType  string     `json:"call_type"` // inbound, outbound, video
	ContactID *uuid.UUID `json:"contact_id,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`

	// Call metrics
	DurationSeconds *int    `json:"duration_seconds,omitempty"`
	Status          string  `json:"status"` // completed, missed, no_answer, busy, failed, cancelled

	// Twilio integration
	TwilioCallSID      *string `json:"twilio_call_sid,omitempty"`
	FromNumber         *string `json:"from_number,omitempty"`
	ToNumber           *string `json:"to_number,omitempty"`
	RecordingURL       *string `json:"recording_url,omitempty"`
	RecordingDuration  *int    `json:"recording_duration,omitempty"`

	// Additional info
	Notes    *string                `json:"notes,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Timestamps
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// Validate checks if the call log has required fields
func (c *CallLog) Validate() error {
	validTypes := map[string]bool{"inbound": true, "outbound": true, "video": true}
	if !validTypes[c.CallType] {
		return ErrInvalidInput{Field: "call_type", Message: "invalid call type"}
	}
	
	validStatuses := map[string]bool{
		"completed": true, "missed": true, "no_answer": true, 
		"busy": true, "failed": true, "cancelled": true,
	}
	if !validStatuses[c.Status] {
		return ErrInvalidInput{Field: "status", Message: "invalid call status"}
	}
	
	return nil
}

// IsCompleted returns true if the call was completed
func (c *CallLog) IsCompleted() bool {
	return c.Status == "completed"
}

// GetDuration returns the call duration in seconds
func (c *CallLog) GetDuration() int {
	if c.DurationSeconds != nil {
		return *c.DurationSeconds
	}
	if c.EndedAt != nil {
		return int(c.EndedAt.Sub(c.StartedAt).Seconds())
	}
	return 0
}
