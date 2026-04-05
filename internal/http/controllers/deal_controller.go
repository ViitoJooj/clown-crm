package controllers

import (
	"net/http"
	"strconv"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/http/dtos"
	"github.com/ViitoJooj/clown-crm/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DealController struct {
	service *services.DealService
}

func NewDealController(service *services.DealService) *DealController {
	return &DealController{
		service: service,
	}
}

// ListDeals godoc
// @Summary List deals
// @Description Get a list of deals with optional filters and pagination
// @Tags deals
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param stage_id query string false "Filter by stage ID"
// @Param assigned_to query string false "Filter by assigned user ID"
// @Param is_won query bool false "Filter by won status"
// @Param is_lost query bool false "Filter by lost status"
// @Success 200 {object} dtos.DealListOutput
// @Router /api/v1/deals [get]
func (c *DealController) ListDeals(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	filters := make(map[string]interface{})
	
	if stageID := ctx.Query("stage_id"); stageID != "" {
		if id, err := uuid.Parse(stageID); err == nil {
			filters["stage_id"] = id
		}
	}
	
	if assignedTo := ctx.Query("assigned_to"); assignedTo != "" {
		if id, err := uuid.Parse(assignedTo); err == nil {
			filters["assigned_to"] = id
		}
	}
	
	if isWon := ctx.Query("is_won"); isWon != "" {
		if val, err := strconv.ParseBool(isWon); err == nil {
			filters["is_won"] = val
		}
	}
	
	if isLost := ctx.Query("is_lost"); isLost != "" {
		if val, err := strconv.ParseBool(isLost); err == nil {
			filters["is_lost"] = val
		}
	}
	
	if contactID := ctx.Query("contact_id"); contactID != "" {
		if id, err := uuid.Parse(contactID); err == nil {
			filters["contact_id"] = id
		}
	}
	
	if companyID := ctx.Query("company_id"); companyID != "" {
		if id, err := uuid.Parse(companyID); err == nil {
			filters["company_id"] = id
		}
	}

	deals, total, err := c.service.ListDeals(ctx, filters, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	output := make([]dtos.DealOutput, len(deals))
	for i, deal := range deals {
		output[i] = mapDealToDealOutput(deal)
	}

	ctx.JSON(http.StatusOK, dtos.DealListOutput{
		Success: true,
		Message: "Deals retrieved successfully",
		Data:    output,
		Page:    page,
		Total:   total,
	})
}

// GetDeal godoc
// @Summary Get a deal
// @Description Get a single deal by ID
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Success 200 {object} dtos.DealDetailOutput
// @Router /api/v1/deals/{id} [get]
func (c *DealController) GetDeal(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	deal, err := c.service.GetDeal(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "deal not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal retrieved successfully",
		Data:    mapDealToDealOutput(deal),
	})
}

// CreateDeal godoc
// @Summary Create a deal
// @Description Create a new deal
// @Tags deals
// @Accept json
// @Produce json
// @Param deal body dtos.CreateDealInput true "Deal data"
// @Success 201 {object} dtos.DealDetailOutput
// @Router /api/v1/deals [post]
func (c *DealController) CreateDeal(ctx *gin.Context) {
	var input dtos.CreateDealInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Get user ID from JWT token/context
	createdBy := uuid.New() // Placeholder

	deal := &domain.Deal{
		Title:             input.Title,
		Value:             input.Value,
		Currency:          input.Currency,
		ExpectedCloseDate: input.ExpectedCloseDate,
		StageID:           input.StageID,
		ContactID:         input.ContactID,
		CompanyID:         input.CompanyID,
		AssignedTo:        input.AssignedTo,
		Source:            input.Source,
		Notes:             input.Notes,
		CustomFields:      input.CustomFields,
		Tags:              input.Tags,
	}

	if err := c.service.CreateDeal(ctx, deal, createdBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal created successfully",
		Data:    mapDealToDealOutput(deal),
	})
}

// UpdateDeal godoc
// @Summary Update a deal
// @Description Update an existing deal
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Param deal body dtos.UpdateDealInput true "Deal data"
// @Success 200 {object} dtos.DealDetailOutput
// @Router /api/v1/deals/{id} [put]
func (c *DealController) UpdateDeal(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	var input dtos.UpdateDealInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get existing deal
	deal, err := c.service.GetDeal(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "deal not found",
		})
		return
	}

	// Update fields if provided
	if input.Title != nil {
		deal.Title = *input.Title
	}
	if input.Value != nil {
		deal.Value = *input.Value
	}
	if input.Currency != nil {
		deal.Currency = *input.Currency
	}
	if input.ExpectedCloseDate != nil {
		deal.ExpectedCloseDate = input.ExpectedCloseDate
	}
	if input.StageID != nil {
		deal.StageID = *input.StageID
	}
	if input.ContactID != nil {
		deal.ContactID = input.ContactID
	}
	if input.CompanyID != nil {
		deal.CompanyID = input.CompanyID
	}
	if input.AssignedTo != nil {
		deal.AssignedTo = input.AssignedTo
	}
	if input.Source != nil {
		deal.Source = input.Source
	}
	if input.Notes != nil {
		deal.Notes = input.Notes
	}
	if input.CustomFields != nil {
		deal.CustomFields = input.CustomFields
	}
	if input.Tags != nil {
		deal.Tags = input.Tags
	}

	// TODO: Get user ID from JWT token/context
	updatedBy := uuid.New() // Placeholder

	if err := c.service.UpdateDeal(ctx, deal, updatedBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal updated successfully",
		Data:    mapDealToDealOutput(deal),
	})
}

// DeleteDeal godoc
// @Summary Delete a deal
// @Description Delete a deal by ID
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Success 200 {object} gin.H
// @Router /api/v1/deals/{id} [delete]
func (c *DealController) DeleteDeal(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	if err := c.service.DeleteDeal(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Deal deleted successfully",
	})
}

// GetPipeline godoc
// @Summary Get pipeline view
// @Description Get all deals organized by pipeline stages
// @Tags deals
// @Accept json
// @Produce json
// @Param assigned_to query string false "Filter by assigned user ID"
// @Success 200 {object} dtos.PipelineOutput
// @Router /api/v1/deals/pipeline [get]
func (c *DealController) GetPipeline(ctx *gin.Context) {
	filters := make(map[string]interface{})
	
	if assignedTo := ctx.Query("assigned_to"); assignedTo != "" {
		if id, err := uuid.Parse(assignedTo); err == nil {
			filters["assigned_to"] = id
		}
	}

	pipeline, err := c.service.GetPipeline(ctx, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Note: This is a simplified version. In production, you'd want to
	// fetch stage details and organize the response properly
	output := make([]dtos.PipelineStageOutput, 0)
	for stageName, deals := range pipeline {
		stageOutput := dtos.PipelineStageOutput{
			StageName: stageName,
			Deals:     make([]dtos.DealOutput, len(deals)),
			DealCount: len(deals),
		}
		
		var totalValue float64
		for i, deal := range deals {
			stageOutput.Deals[i] = mapDealToDealOutput(deal)
			totalValue += deal.Value
		}
		stageOutput.TotalValue = totalValue
		
		output = append(output, stageOutput)
	}

	ctx.JSON(http.StatusOK, dtos.PipelineOutput{
		Success: true,
		Message: "Pipeline retrieved successfully",
		Data:    output,
	})
}

// MoveDeal godoc
// @Summary Move deal to stage
// @Description Move a deal to a different pipeline stage
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Param input body dtos.MoveDealInput true "Stage ID"
// @Success 200 {object} dtos.DealDetailOutput
// @Router /api/v1/deals/{id}/move [post]
func (c *DealController) MoveDeal(ctx *gin.Context) {
	dealID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	var input dtos.MoveDealInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Get user ID from JWT token/context
	movedBy := uuid.New() // Placeholder

	if err := c.service.MoveDealToStage(ctx, dealID, input.StageID, movedBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get updated deal
	deal, _ := c.service.GetDeal(ctx, dealID)

	ctx.JSON(http.StatusOK, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal moved successfully",
		Data:    mapDealToDealOutput(deal),
	})
}

// WinDeal godoc
// @Summary Mark deal as won
// @Description Mark a deal as won
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Success 200 {object} dtos.DealDetailOutput
// @Router /api/v1/deals/{id}/win [post]
func (c *DealController) WinDeal(ctx *gin.Context) {
	dealID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	// TODO: Get user ID from JWT token/context
	userID := uuid.New() // Placeholder

	if err := c.service.WinDeal(ctx, dealID, userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get updated deal
	deal, _ := c.service.GetDeal(ctx, dealID)

	ctx.JSON(http.StatusOK, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal marked as won",
		Data:    mapDealToDealOutput(deal),
	})
}

// LoseDeal godoc
// @Summary Mark deal as lost
// @Description Mark a deal as lost with a reason
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Param input body dtos.LoseDealInput true "Lost reason"
// @Success 200 {object} dtos.DealDetailOutput
// @Router /api/v1/deals/{id}/lose [post]
func (c *DealController) LoseDeal(ctx *gin.Context) {
	dealID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	var input dtos.LoseDealInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Get user ID from JWT token/context
	userID := uuid.New() // Placeholder

	if err := c.service.LoseDeal(ctx, dealID, userID, input.Reason); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get updated deal
	deal, _ := c.service.GetDeal(ctx, dealID)

	ctx.JSON(http.StatusOK, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal marked as lost",
		Data:    mapDealToDealOutput(deal),
	})
}

// AssignDeal godoc
// @Summary Assign deal to user
// @Description Assign a deal to a specific user
// @Tags deals
// @Accept json
// @Produce json
// @Param id path string true "Deal ID"
// @Param input body dtos.AssignDealInput true "User ID"
// @Success 200 {object} dtos.DealDetailOutput
// @Router /api/v1/deals/{id}/assign [post]
func (c *DealController) AssignDeal(ctx *gin.Context) {
	dealID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid deal ID",
		})
		return
	}

	var input dtos.AssignDealInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Get user ID from JWT token/context
	assignedBy := uuid.New() // Placeholder

	if err := c.service.AssignDeal(ctx, dealID, input.UserID, assignedBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get updated deal
	deal, _ := c.service.GetDeal(ctx, dealID)

	ctx.JSON(http.StatusOK, dtos.DealDetailOutput{
		Success: true,
		Message: "Deal assigned successfully",
		Data:    mapDealToDealOutput(deal),
	})
}

// GetMetrics godoc
// @Summary Get pipeline metrics
// @Description Get various metrics for the pipeline
// @Tags deals
// @Accept json
// @Produce json
// @Success 200 {object} dtos.MetricsOutput
// @Router /api/v1/deals/metrics [get]
func (c *DealController) GetMetrics(ctx *gin.Context) {
	metrics, err := c.service.CalculatePipelineMetrics(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.MetricsOutput{
		Success: true,
		Message: "Metrics retrieved successfully",
		Data:    metrics,
	})
}

// Helper function to map domain.Deal to dtos.DealOutput
func mapDealToDealOutput(deal *domain.Deal) dtos.DealOutput {
	return dtos.DealOutput{
		ID:                deal.ID,
		Title:             deal.Title,
		Value:             deal.Value,
		Currency:          deal.Currency,
		ExpectedCloseDate: deal.ExpectedCloseDate,
		StageID:           deal.StageID,
		Probability:       deal.Probability,
		ContactID:         deal.ContactID,
		CompanyID:         deal.CompanyID,
		AssignedTo:        deal.AssignedTo,
		OwnerID:           deal.OwnerID,
		Source:            deal.Source,
		Notes:             deal.Notes,
		CustomFields:      deal.CustomFields,
		Tags:              deal.Tags,
		IsWon:             deal.IsWon,
		IsLost:            deal.IsLost,
		LostReason:        deal.LostReason,
		CreatedAt:         deal.CreatedAt,
		UpdatedAt:         deal.UpdatedAt,
		ClosedAt:          deal.ClosedAt,
	}
}
