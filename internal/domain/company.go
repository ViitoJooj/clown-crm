package domain

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Industry    *string    `json:"industry,omitempty"`
	CompanySize *string    `json:"company_size,omitempty"`
	Website     *string    `json:"website,omitempty"`
	Phone       *string    `json:"phone,omitempty"`

	// Address
	Address    *string `json:"address,omitempty"`
	City       *string `json:"city,omitempty"`
	State      *string `json:"state,omitempty"`
	Country    *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`

	// Status and metadata
	Status       string                 `json:"status"` // active, inactive, prospect
	Tags         []string               `json:"tags,omitempty"`
	Notes        *string                `json:"notes,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`

	// Ownership
	OwnerID *uuid.UUID `json:"owner_id,omitempty"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate checks if the company has required fields
func (c *Company) Validate() error {
	if c.Name == "" {
		return ErrInvalidInput{Field: "name", Message: "company name is required"}
	}
	return nil
}

// IsActive returns true if the company is active
func (c *Company) IsActive() bool {
	return c.Status == "active"
}
