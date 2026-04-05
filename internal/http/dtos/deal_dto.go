package dtos

import (
	"time"

	"github.com/google/uuid"
)

// CreateDealInput represents the input for creating a deal
type CreateDealInput struct {
	Title             string                 `json:"title" binding:"required"`
	Value             float64                `json:"value" binding:"required,min=0"`
	Currency          string                 `json:"currency"`
	ExpectedCloseDate *time.Time             `json:"expected_close_date"`
	StageID           uuid.UUID              `json:"stage_id" binding:"required"`
	ContactID         *uuid.UUID             `json:"contact_id"`
	CompanyID         *uuid.UUID             `json:"company_id"`
	AssignedTo        *uuid.UUID             `json:"assigned_to"`
	Source            *string                `json:"source"`
	Notes             *string                `json:"notes"`
	CustomFields      map[string]interface{} `json:"custom_fields"`
	Tags              []string               `json:"tags"`
}

// UpdateDealInput represents the input for updating a deal
type UpdateDealInput struct {
	Title             *string                `json:"title"`
	Value             *float64               `json:"value" binding:"omitempty,min=0"`
	Currency          *string                `json:"currency"`
	ExpectedCloseDate *time.Time             `json:"expected_close_date"`
	StageID           *uuid.UUID             `json:"stage_id"`
	ContactID         *uuid.UUID             `json:"contact_id"`
	CompanyID         *uuid.UUID             `json:"company_id"`
	AssignedTo        *uuid.UUID             `json:"assigned_to"`
	Source            *string                `json:"source"`
	Notes             *string                `json:"notes"`
	CustomFields      map[string]interface{} `json:"custom_fields"`
	Tags              []string               `json:"tags"`
}

// MoveDealInput represents the input for moving a deal to a different stage
type MoveDealInput struct {
	StageID uuid.UUID `json:"stage_id" binding:"required"`
}

// LoseDealInput represents the input for marking a deal as lost
type LoseDealInput struct {
	Reason string `json:"reason" binding:"required"`
}

// AssignDealInput represents the input for assigning a deal to a user
type AssignDealInput struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

// DealOutput represents the output for a deal
type DealOutput struct {
	ID                uuid.UUID              `json:"id"`
	Title             string                 `json:"title"`
	Value             float64                `json:"value"`
	Currency          string                 `json:"currency"`
	ExpectedCloseDate *time.Time             `json:"expected_close_date,omitempty"`
	StageID           uuid.UUID              `json:"stage_id"`
	Probability       int                    `json:"probability"`
	ContactID         *uuid.UUID             `json:"contact_id,omitempty"`
	CompanyID         *uuid.UUID             `json:"company_id,omitempty"`
	AssignedTo        *uuid.UUID             `json:"assigned_to,omitempty"`
	OwnerID           *uuid.UUID             `json:"owner_id,omitempty"`
	Source            *string                `json:"source,omitempty"`
	Notes             *string                `json:"notes,omitempty"`
	CustomFields      map[string]interface{} `json:"custom_fields,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	IsWon             bool                   `json:"is_won"`
	IsLost            bool                   `json:"is_lost"`
	LostReason        *string                `json:"lost_reason,omitempty"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	ClosedAt          *time.Time             `json:"closed_at,omitempty"`
}

// DealListOutput represents a paginated list of deals
type DealListOutput struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []DealOutput `json:"data"`
	Page    int          `json:"page"`
	Total   int          `json:"total"`
}

// DealDetailOutput represents a single deal response
type DealDetailOutput struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    DealOutput `json:"data"`
}

// PipelineStageOutput represents a pipeline stage with deals
type PipelineStageOutput struct {
	StageID      uuid.UUID    `json:"stage_id"`
	StageName    string       `json:"stage_name"`
	Probability  int          `json:"probability"`
	Color        string       `json:"color"`
	DisplayOrder int          `json:"display_order"`
	Deals        []DealOutput `json:"deals"`
	TotalValue   float64      `json:"total_value"`
	DealCount    int          `json:"deal_count"`
}

// PipelineOutput represents the entire pipeline view
type PipelineOutput struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    []PipelineStageOutput `json:"data"`
}

// MetricsOutput represents pipeline metrics
type MetricsOutput struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// CreatePipelineStageInput represents the input for creating a pipeline stage
type CreatePipelineStageInput struct {
	Name         string `json:"name" binding:"required"`
	DisplayOrder int    `json:"display_order" binding:"required"`
	Probability  int    `json:"probability" binding:"required,min=0,max=100"`
	Color        string `json:"color"`
	IsClosedWon  bool   `json:"is_closed_won"`
	IsClosedLost bool   `json:"is_closed_lost"`
}

// UpdatePipelineStageInput represents the input for updating a pipeline stage
type UpdatePipelineStageInput struct {
	Name         *string `json:"name"`
	DisplayOrder *int    `json:"display_order"`
	Probability  *int    `json:"probability" binding:"omitempty,min=0,max=100"`
	Color        *string `json:"color"`
	IsActive     *bool   `json:"is_active"`
	IsClosedWon  *bool   `json:"is_closed_won"`
	IsClosedLost *bool   `json:"is_closed_lost"`
}

// PipelineStageDetailOutput represents a single pipeline stage
type PipelineStageDetailOutput struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	DisplayOrder int       `json:"display_order"`
	Probability  int       `json:"probability"`
	Color        string    `json:"color"`
	IsActive     bool      `json:"is_active"`
	IsClosedWon  bool      `json:"is_closed_won"`
	IsClosedLost bool      `json:"is_closed_lost"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PipelineStageListOutput represents a list of pipeline stages
type PipelineStageListOutput struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Data    []PipelineStageDetailOutput `json:"data"`
}

// PipelineStageOutput represents a single stage response
type PipelineStageOutputSingle struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Data    PipelineStageDetailOutput `json:"data"`
}
