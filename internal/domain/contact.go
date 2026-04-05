package domain

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID        uuid.UUID  `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     *string    `json:"email,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
	CompanyID *uuid.UUID `json:"company_id,omitempty"`
	Position  *string    `json:"position,omitempty"`

	// Address
	Address    *string `json:"address,omitempty"`
	City       *string `json:"city,omitempty"`
	State      *string `json:"state,omitempty"`
	Country    *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`

	// Status and categorization
	Source       *string `json:"source,omitempty"`
	Status       string  `json:"status"` // lead, prospect, customer, inactive
	Tags         []string `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`

	// Ownership
	AssignedTo *uuid.UUID `json:"assigned_to,omitempty"`
	OwnerID    *uuid.UUID `json:"owner_id,omitempty"`

	// Metadata
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LastContactAt *time.Time `json:"last_contact_at,omitempty"`
}

// FullName returns the contact's full name
func (c *Contact) FullName() string {
	return c.FirstName + " " + c.LastName
}

// IsLead returns true if the contact is a lead
func (c *Contact) IsLead() bool {
	return c.Status == "lead"
}

// IsCustomer returns true if the contact is a customer
func (c *Contact) IsCustomer() bool {
	return c.Status == "customer"
}

// Validate checks if the contact has required fields
func (c *Contact) Validate() error {
	if c.FirstName == "" {
		return ErrInvalidInput{Field: "first_name", Message: "first name is required"}
	}
	if c.LastName == "" {
		return ErrInvalidInput{Field: "last_name", Message: "last name is required"}
	}
	if c.Email == nil && c.Phone == nil {
		return ErrInvalidInput{Field: "contact", Message: "email or phone is required"}
	}
	return nil
}
