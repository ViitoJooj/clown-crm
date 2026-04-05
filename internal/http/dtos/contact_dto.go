package dtos

import "github.com/google/uuid"

// Contact input DTOs
type CreateContactInput struct {
	FirstName    string                 `json:"first_name" binding:"required"`
	LastName     string                 `json:"last_name" binding:"required"`
	Email        *string                `json:"email,omitempty"`
	Phone        *string                `json:"phone,omitempty"`
	CompanyID    *uuid.UUID             `json:"company_id,omitempty"`
	Position     *string                `json:"position,omitempty"`
	Address      *string                `json:"address,omitempty"`
	City         *string                `json:"city,omitempty"`
	State        *string                `json:"state,omitempty"`
	Country      *string                `json:"country,omitempty"`
	PostalCode   *string                `json:"postal_code,omitempty"`
	Source       *string                `json:"source,omitempty"`
	Status       *string                `json:"status,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

type UpdateContactInput struct {
	FirstName     string                 `json:"first_name" binding:"required"`
	LastName      string                 `json:"last_name" binding:"required"`
	Email         *string                `json:"email,omitempty"`
	Phone         *string                `json:"phone,omitempty"`
	CompanyID     *uuid.UUID             `json:"company_id,omitempty"`
	Position      *string                `json:"position,omitempty"`
	Address       *string                `json:"address,omitempty"`
	City          *string                `json:"city,omitempty"`
	State         *string                `json:"state,omitempty"`
	Country       *string                `json:"country,omitempty"`
	PostalCode    *string                `json:"postal_code,omitempty"`
	Source        *string                `json:"source,omitempty"`
	Status        string                 `json:"status" binding:"required"`
	Tags          []string               `json:"tags,omitempty"`
	CustomFields  map[string]interface{} `json:"custom_fields,omitempty"`
	AssignedTo    *uuid.UUID             `json:"assigned_to,omitempty"`
	LastContactAt *string                `json:"last_contact_at,omitempty"`
}

type AssignContactInput struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

type AddTagInput struct {
	Tag string `json:"tag" binding:"required"`
}

// Contact output DTOs
type ContactOutput struct {
	ID            uuid.UUID              `json:"id"`
	FirstName     string                 `json:"first_name"`
	LastName      string                 `json:"last_name"`
	FullName      string                 `json:"full_name"`
	Email         *string                `json:"email,omitempty"`
	Phone         *string                `json:"phone,omitempty"`
	CompanyID     *uuid.UUID             `json:"company_id,omitempty"`
	Position      *string                `json:"position,omitempty"`
	Address       *string                `json:"address,omitempty"`
	City          *string                `json:"city,omitempty"`
	State         *string                `json:"state,omitempty"`
	Country       *string                `json:"country,omitempty"`
	PostalCode    *string                `json:"postal_code,omitempty"`
	Source        *string                `json:"source,omitempty"`
	Status        string                 `json:"status"`
	Tags          []string               `json:"tags,omitempty"`
	CustomFields  map[string]interface{} `json:"custom_fields,omitempty"`
	AssignedTo    *uuid.UUID             `json:"assigned_to,omitempty"`
	OwnerID       *uuid.UUID             `json:"owner_id,omitempty"`
	LastContactAt *string                `json:"last_contact_at,omitempty"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

type ContactListResponse struct {
	Success  bool            `json:"success"`
	Data     []ContactOutput `json:"data"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int             `json:"total"`
}

type ContactResponse struct {
	Success bool          `json:"success"`
	Data    ContactOutput `json:"data"`
}

type ActivityOutput struct {
	ID              uuid.UUID              `json:"id"`
	ActivityType    string                 `json:"activity_type"`
	Title           string                 `json:"title"`
	Description     *string                `json:"description,omitempty"`
	DurationMinutes *int                   `json:"duration_minutes,omitempty"`
	ContactID       *uuid.UUID             `json:"contact_id,omitempty"`
	DealID          *uuid.UUID             `json:"deal_id,omitempty"`
	CompanyID       *uuid.UUID             `json:"company_id,omitempty"`
	UserID          *uuid.UUID             `json:"user_id,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt       string                 `json:"created_at"`
}

type ActivityListResponse struct {
	Success  bool             `json:"success"`
	Data     []ActivityOutput `json:"data"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
	Total    int              `json:"total"`
}

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
