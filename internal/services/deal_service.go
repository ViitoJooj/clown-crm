package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/repository"
	"github.com/google/uuid"
)

type DealService struct {
	dealRepo     repository.DealRepository
	stageRepo    repository.PipelineStageRepository
	activityRepo repository.ActivityRepository
}

func NewDealService(
	dealRepo repository.DealRepository,
	stageRepo repository.PipelineStageRepository,
	activityRepo repository.ActivityRepository,
) *DealService {
	return &DealService{
		dealRepo:     dealRepo,
		stageRepo:    stageRepo,
		activityRepo: activityRepo,
	}
}

func (s *DealService) CreateDeal(ctx context.Context, deal *domain.Deal, createdBy uuid.UUID) error {
	// Validate
	if err := deal.Validate(); err != nil {
		return err
	}
	
	// Get stage to set probability
	stage, err := s.stageRepo.GetByID(ctx, deal.StageID)
	if err != nil {
		return fmt.Errorf("invalid stage: %w", err)
	}
	
	// Set defaults
	deal.ID = uuid.New()
	deal.CreatedAt = time.Now()
	deal.UpdatedAt = time.Now()
	deal.Probability = stage.Probability
	
	if deal.OwnerID == nil {
		deal.OwnerID = &createdBy
	}
	
	if deal.Currency == "" {
		deal.Currency = "USD"
	}
	
	// Create deal
	if err := s.dealRepo.Create(ctx, deal); err != nil {
		return fmt.Errorf("failed to create deal: %w", err)
	}
	
	// Log activity
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "deal_stage_change",
		Title:        "Deal Created",
		Description:  strPtr(fmt.Sprintf("Deal '%s' created in stage %s", deal.Title, stage.Name)),
		DealID:       &deal.ID,
		UserID:       &createdBy,
		Metadata: map[string]interface{}{
			"stage":       stage.Name,
			"value":       deal.Value,
			"probability": deal.Probability,
		},
		CreatedAt: time.Now(),
	}
	s.activityRepo.Create(ctx, activity)
	
	return nil
}

// GetDeal retrieves a deal by ID
func (s *DealService) GetDeal(ctx context.Context, id uuid.UUID) (*domain.Deal, error) {
	return s.dealRepo.GetByID(ctx, id)
}

func (s *DealService) UpdateDeal(ctx context.Context, deal *domain.Deal, updatedBy uuid.UUID) error {
	// Validate
	if err := deal.Validate(); err != nil {
		return err
	}
	
	// Get existing to track changes
	existing, err := s.dealRepo.GetByID(ctx, deal.ID)
	if err != nil {
		return err
	}
	
	deal.UpdatedAt = time.Now()
	
	// Update deal
	if err := s.dealRepo.Update(ctx, deal); err != nil {
		return fmt.Errorf("failed to update deal: %w", err)
	}
	
	// Log stage change if changed
	if existing.StageID != deal.StageID {
		oldStage, _ := s.stageRepo.GetByID(ctx, existing.StageID)
		newStage, _ := s.stageRepo.GetByID(ctx, deal.StageID)
		
		activity := &domain.Activity{
			ID:           uuid.New(),
			ActivityType: "deal_stage_change",
			Title:        "Deal Stage Changed",
			Description:  strPtr(fmt.Sprintf("Deal moved from %s to %s", oldStage.Name, newStage.Name)),
			DealID:       &deal.ID,
			UserID:       &updatedBy,
			Metadata: map[string]interface{}{
				"old_stage": oldStage.Name,
				"new_stage": newStage.Name,
				"old_probability": existing.Probability,
				"new_probability": deal.Probability,
			},
			CreatedAt: time.Now(),
		}
		s.activityRepo.Create(ctx, activity)
	}
	
	return nil
}

// MoveDealToStage moves a deal to a new stage
func (s *DealService) MoveDealToStage(ctx context.Context, dealID, stageID, movedBy uuid.UUID) error {
	deal, err := s.dealRepo.GetByID(ctx, dealID)
	if err != nil {
		return err
	}
	
	// Get new stage
	newStage, err := s.stageRepo.GetByID(ctx, stageID)
	if err != nil {
		return fmt.Errorf("invalid stage: %w", err)
	}
	
	// Update stage and probability
	oldStageID := deal.StageID
	deal.StageID = stageID
	deal.Probability = newStage.Probability
	deal.UpdatedAt = time.Now()
	
	// Check if deal is being closed
	if newStage.IsClosedWon {
		deal.MarkWon()
	} else if newStage.IsClosedLost {
		deal.MarkLost("Moved to closed lost stage")
	} else if deal.IsClosed() {
		// Reopening a closed deal
		deal.Reopen()
	}
	
	if err := s.dealRepo.Update(ctx, deal); err != nil {
		return err
	}
	
	// Log activity
	oldStage, _ := s.stageRepo.GetByID(ctx, oldStageID)
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "deal_stage_change",
		Title:        fmt.Sprintf("Deal moved to %s", newStage.Name),
		Description:  strPtr(fmt.Sprintf("Deal '%s' moved from %s to %s", deal.Title, oldStage.Name, newStage.Name)),
		DealID:       &dealID,
		UserID:       &movedBy,
		Metadata: map[string]interface{}{
			"old_stage":       oldStage.Name,
			"new_stage":       newStage.Name,
			"old_probability": oldStage.Probability,
			"new_probability": newStage.Probability,
		},
		CreatedAt: time.Now(),
	}
	s.activityRepo.Create(ctx, activity)
	
	return nil
}

// WinDeal marks a deal as won
func (s *DealService) WinDeal(ctx context.Context, dealID, userID uuid.UUID) error {
	deal, err := s.dealRepo.GetByID(ctx, dealID)
	if err != nil {
		return err
	}
	
	// Find the closed won stage
	stages, _, err := s.stageRepo.List(ctx, map[string]interface{}{"is_closed_won": true}, 1, 1)
	if err != nil || len(stages) == 0 {
		return fmt.Errorf("no closed won stage found")
	}
	
	deal.MarkWon()
	deal.StageID = stages[0].ID
	deal.UpdatedAt = time.Now()
	
	if err := s.dealRepo.Update(ctx, deal); err != nil {
		return err
	}
	
	// Log activity
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "deal_stage_change",
		Title:        "Deal Won!",
		Description:  strPtr(fmt.Sprintf("Deal '%s' marked as won (Value: %.2f %s)", deal.Title, deal.Value, deal.Currency)),
		DealID:       &dealID,
		UserID:       &userID,
		Metadata: map[string]interface{}{
			"value":    deal.Value,
			"currency": deal.Currency,
		},
		CreatedAt: time.Now(),
	}
	s.activityRepo.Create(ctx, activity)
	
	return nil
}

// LoseDeal marks a deal as lost
func (s *DealService) LoseDeal(ctx context.Context, dealID, userID uuid.UUID, reason string) error {
	deal, err := s.dealRepo.GetByID(ctx, dealID)
	if err != nil {
		return err
	}
	
	// Find the closed lost stage
	stages, _, err := s.stageRepo.List(ctx, map[string]interface{}{"is_closed_lost": true}, 1, 1)
	if err != nil || len(stages) == 0 {
		return fmt.Errorf("no closed lost stage found")
	}
	
	deal.MarkLost(reason)
	deal.StageID = stages[0].ID
	deal.UpdatedAt = time.Now()
	
	if err := s.dealRepo.Update(ctx, deal); err != nil {
		return err
	}
	
	// Log activity
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "deal_stage_change",
		Title:        "Deal Lost",
		Description:  strPtr(fmt.Sprintf("Deal '%s' marked as lost. Reason: %s", deal.Title, reason)),
		DealID:       &dealID,
		UserID:       &userID,
		Metadata: map[string]interface{}{
			"reason": reason,
		},
		CreatedAt: time.Now(),
	}
	s.activityRepo.Create(ctx, activity)
	
	return nil
}

func (s *DealService) DeleteDeal(ctx context.Context, id uuid.UUID) error {
	return s.dealRepo.Delete(ctx, id)
}

// ListDeals lists deals with filters
func (s *DealService) ListDeals(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Deal, int, error) {
	return s.dealRepo.List(ctx, filters, page, pageSize)
}

// GetPipeline retrieves all deals organized by stage
func (s *DealService) GetPipeline(ctx context.Context, filters map[string]interface{}) (map[string][]*domain.Deal, error) {
	// Get all active stages
	stages, _, err := s.stageRepo.List(ctx, map[string]interface{}{"is_active": true}, 1, 100)
	if err != nil {
		return nil, err
	}
	
	pipeline := make(map[string][]*domain.Deal)
	
	// Get deals for each stage
	for _, stage := range stages {
		stageFilters := make(map[string]interface{})
		for k, v := range filters {
			stageFilters[k] = v
		}
		stageFilters["stage_id"] = stage.ID
		
		deals, _, err := s.dealRepo.List(ctx, stageFilters, 1, 1000)
		if err != nil {
			continue
		}
		
		pipeline[stage.Name] = deals
	}
	
	return pipeline, nil
}

// GetDealsByStage retrieves deals in a specific stage
func (s *DealService) GetDealsByStage(ctx context.Context, stageID uuid.UUID, page, pageSize int) ([]*domain.Deal, int, error) {
	return s.dealRepo.List(ctx, map[string]interface{}{"stage_id": stageID}, page, pageSize)
}

// AssignDeal assigns a deal to a user
func (s *DealService) AssignDeal(ctx context.Context, dealID, userID, assignedBy uuid.UUID) error {
	deal, err := s.dealRepo.GetByID(ctx, dealID)
	if err != nil {
		return err
	}
	
	oldAssignedTo := deal.AssignedTo
	deal.AssignedTo = &userID
	deal.UpdatedAt = time.Now()
	
	if err := s.dealRepo.Update(ctx, deal); err != nil {
		return err
	}
	
	// Log activity
	var desc string
	if oldAssignedTo == nil {
		desc = fmt.Sprintf("Deal assigned to user %s", userID)
	} else {
		desc = fmt.Sprintf("Deal reassigned from %s to %s", *oldAssignedTo, userID)
	}
	
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "assignment",
		Title:        "Deal Assigned",
		Description:  &desc,
		DealID:       &dealID,
		UserID:       &assignedBy,
		Metadata: map[string]interface{}{
			"assigned_to": userID,
			"old_assigned_to": oldAssignedTo,
		},
		CreatedAt: time.Now(),
	}
	s.activityRepo.Create(ctx, activity)
	
	return nil
}

// CalculatePipelineMetrics calculates metrics for the pipeline
func (s *DealService) CalculatePipelineMetrics(ctx context.Context) (map[string]interface{}, error) {
	metrics := make(map[string]interface{})
	
	// Total deal count
	allDeals, totalCount, err := s.dealRepo.List(ctx, map[string]interface{}{}, 1, 10000)
	if err != nil {
		return nil, err
	}
	
	metrics["total_deals"] = totalCount
	
	// Active deals (not closed)
	activeCount := 0
	var totalValue float64
	var weightedValue float64
	
	for _, deal := range allDeals {
		if deal.IsActive() {
			activeCount++
			totalValue += deal.Value
			weightedValue += deal.Value * (float64(deal.Probability) / 100.0)
		}
	}
	
	metrics["active_deals"] = activeCount
	metrics["total_pipeline_value"] = totalValue
	metrics["weighted_pipeline_value"] = weightedValue
	
	// Won/Lost deals
	wonCount := 0
	lostCount := 0
	var wonValue float64
	
	for _, deal := range allDeals {
		if deal.IsWon {
			wonCount++
			wonValue += deal.Value
		} else if deal.IsLost {
			lostCount++
		}
	}
	
	metrics["won_deals"] = wonCount
	metrics["lost_deals"] = lostCount
	metrics["won_value"] = wonValue
	
	// Win rate
	if wonCount+lostCount > 0 {
		metrics["win_rate"] = float64(wonCount) / float64(wonCount+lostCount) * 100
	} else {
		metrics["win_rate"] = 0
	}
	
	return metrics, nil
}
