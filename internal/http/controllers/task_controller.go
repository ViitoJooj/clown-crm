package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/http/dtos"
	"github.com/ViitoJooj/clown-crm/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	service *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

// ListTasks handles GET /api/v1/tasks - list tasks with filters
func (c *TaskController) ListTasks(ctx *gin.Context) {
	// Pagination
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "20"))

	// Build filters
	filters := make(map[string]interface{})

	if status := ctx.Query("status"); status != "" {
		filters["status"] = status
	}
	if priority := ctx.Query("priority"); priority != "" {
		filters["priority"] = priority
	}
	if assignedTo := ctx.Query("assigned_to"); assignedTo != "" {
		if userID, err := uuid.Parse(assignedTo); err == nil {
			filters["assigned_to"] = userID
		}
	}
	if dueDate := ctx.Query("due_date"); dueDate != "" {
		if date, err := time.Parse("2006-01-02", dueDate); err == nil {
			filters["due_date"] = date
		}
	}
	if taskType := ctx.Query("task_type"); taskType != "" {
		filters["task_type"] = taskType
	}

	tasks, total, err := c.service.ListTasks(ctx.Request.Context(), filters, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Convert to DTOs
	taskDTOs := make([]dtos.OutputTaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = c.taskToDTO(task)
	}

	ctx.JSON(http.StatusOK, dtos.TaskListResponse{
		Success: true,
		Message: "Tasks retrieved successfully",
		Tasks:   taskDTOs,
		Page:    page,
		PerPage: perPage,
		Total:   total,
	})
}

// GetTask handles GET /api/v1/tasks/:id - get a single task
func (c *TaskController) GetTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	task, err := c.service.GetTask(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.TaskResponse{
		Success: true,
		Message: "Task retrieved successfully",
		Task:    c.taskToDTO(task),
	})
}

// CreateTask handles POST /api/v1/tasks - create a new task
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var input dtos.InputTaskDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get user ID from context (should be set by auth middleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	createdBy, ok := userID.(uuid.UUID)
	if !ok {
		// Try to parse if it's a string
		userIDStr, ok := userID.(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid user ID format",
			})
			return
		}
		var err error
		createdBy, err = uuid.Parse(userIDStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid user ID format",
			})
			return
		}
	}

	task := &domain.Task{
		Title:         input.Title,
		Description:   input.Description,
		TaskType:      input.TaskType,
		Priority:      input.Priority,
		Status:        input.Status,
		DueDate:       input.DueDate,
		ReminderAt:    input.ReminderAt,
		AssignedTo:    input.AssignedTo,
		RelatedToType: input.RelatedToType,
		RelatedToID:   input.RelatedToID,
	}

	if err := c.service.CreateTask(ctx.Request.Context(), task, createdBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.TaskResponse{
		Success: true,
		Message: "Task created successfully",
		Task:    c.taskToDTO(task),
	})
}

// UpdateTask handles PUT /api/v1/tasks/:id - update a task
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var input dtos.UpdateTaskDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get existing task
	task, err := c.service.GetTask(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Update fields if provided
	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = input.Description
	}
	if input.TaskType != nil {
		task.TaskType = *input.TaskType
	}
	if input.Priority != nil {
		task.Priority = *input.Priority
	}
	if input.Status != nil {
		task.Status = *input.Status
	}
	if input.DueDate != nil {
		task.DueDate = input.DueDate
	}
	if input.ReminderAt != nil {
		task.ReminderAt = input.ReminderAt
	}
	if input.AssignedTo != nil {
		task.AssignedTo = input.AssignedTo
	}
	if input.RelatedToType != nil {
		task.RelatedToType = input.RelatedToType
	}
	if input.RelatedToID != nil {
		task.RelatedToID = input.RelatedToID
	}

	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	updatedBy := c.parseUserID(userID)

	if err := c.service.UpdateTask(ctx.Request.Context(), task, updatedBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.TaskResponse{
		Success: true,
		Message: "Task updated successfully",
		Task:    c.taskToDTO(task),
	})
}

// DeleteTask handles DELETE /api/v1/tasks/:id - delete a task
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	if err := c.service.DeleteTask(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Task deleted successfully",
	})
}

// CompleteTask handles POST /api/v1/tasks/:id/complete - mark task as complete
func (c *TaskController) CompleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	completedBy := c.parseUserID(userID)

	if err := c.service.CompleteTask(ctx.Request.Context(), id, completedBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	task, _ := c.service.GetTask(ctx.Request.Context(), id)

	ctx.JSON(http.StatusOK, dtos.TaskResponse{
		Success: true,
		Message: "Task marked as complete",
		Task:    c.taskToDTO(task),
	})
}

// CancelTask handles POST /api/v1/tasks/:id/cancel - mark task as cancelled
func (c *TaskController) CancelTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	cancelledBy := c.parseUserID(userID)

	if err := c.service.CancelTask(ctx.Request.Context(), id, cancelledBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	task, _ := c.service.GetTask(ctx.Request.Context(), id)

	ctx.JSON(http.StatusOK, dtos.TaskResponse{
		Success: true,
		Message: "Task marked as cancelled",
		Task:    c.taskToDTO(task),
	})
}

// GetMyTasks handles GET /api/v1/tasks/my - get current user's tasks
func (c *TaskController) GetMyTasks(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	parsedUserID := c.parseUserID(userID)
	
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "20"))
	status := ctx.Query("status")

	tasks, total, err := c.service.GetUserTasks(ctx.Request.Context(), parsedUserID, status, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	taskDTOs := make([]dtos.OutputTaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = c.taskToDTO(task)
	}

	ctx.JSON(http.StatusOK, dtos.TaskListResponse{
		Success: true,
		Message: "Your tasks retrieved successfully",
		Tasks:   taskDTOs,
		Page:    page,
		PerPage: perPage,
		Total:   total,
	})
}

// GetOverdueTasks handles GET /api/v1/tasks/overdue - get overdue tasks
func (c *TaskController) GetOverdueTasks(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	parsedUserID := c.parseUserID(userID)

	tasks, err := c.service.GetOverdueTasks(ctx.Request.Context(), &parsedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	taskDTOs := make([]dtos.OutputTaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = c.taskToDTO(task)
	}

	ctx.JSON(http.StatusOK, dtos.TaskListResponse{
		Success: true,
		Message: "Overdue tasks retrieved successfully",
		Tasks:   taskDTOs,
		Page:    1,
		PerPage: len(tasks),
		Total:   len(tasks),
	})
}

// GetUpcomingTasks handles GET /api/v1/tasks/upcoming - get tasks due within 24 hours
func (c *TaskController) GetUpcomingTasks(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	parsedUserID := c.parseUserID(userID)

	tasks, err := c.service.GetUpcomingTasks(ctx.Request.Context(), parsedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	taskDTOs := make([]dtos.OutputTaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = c.taskToDTO(task)
	}

	ctx.JSON(http.StatusOK, dtos.TaskListResponse{
		Success: true,
		Message: "Upcoming tasks retrieved successfully",
		Tasks:   taskDTOs,
		Page:    1,
		PerPage: len(tasks),
		Total:   len(tasks),
	})
}

// AssignTask handles POST /api/v1/tasks/:id/assign - assign task to a user
func (c *TaskController) AssignTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	taskID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var input dtos.AssignTaskDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	assignedBy := c.parseUserID(userID)

	if err := c.service.AssignTask(ctx.Request.Context(), taskID, input.UserID, assignedBy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	task, _ := c.service.GetTask(ctx.Request.Context(), taskID)

	ctx.JSON(http.StatusOK, dtos.TaskResponse{
		Success: true,
		Message: "Task assigned successfully",
		Task:    c.taskToDTO(task),
	})
}

// SetReminder handles POST /api/v1/tasks/:id/reminder - set a reminder for a task
func (c *TaskController) SetReminder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	taskID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var input dtos.SetReminderDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := c.service.SetReminder(ctx.Request.Context(), taskID, input.ReminderAt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	task, _ := c.service.GetTask(ctx.Request.Context(), taskID)

	ctx.JSON(http.StatusOK, dtos.TaskResponse{
		Success: true,
		Message: "Reminder set successfully",
		Task:    c.taskToDTO(task),
	})
}

// GetTaskStats handles GET /api/v1/tasks/stats - get task statistics
func (c *TaskController) GetTaskStats(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	parsedUserID := c.parseUserID(userID)

	stats, err := c.service.GetTaskStats(ctx.Request.Context(), parsedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.TaskStatsResponse{
		Success: true,
		Message: "Task statistics retrieved successfully",
		Stats:   stats,
	})
}

// Helper function to convert domain.Task to OutputTaskDTO
func (c *TaskController) taskToDTO(task *domain.Task) dtos.OutputTaskDTO {
	return dtos.OutputTaskDTO{
		ID:            task.ID,
		Title:         task.Title,
		Description:   task.Description,
		TaskType:      task.TaskType,
		Priority:      task.Priority,
		Status:        task.Status,
		DueDate:       task.DueDate,
		ReminderAt:    task.ReminderAt,
		AssignedTo:    task.AssignedTo,
		CreatedBy:     task.CreatedBy,
		RelatedToType: task.RelatedToType,
		RelatedToID:   task.RelatedToID,
		CompletedAt:   task.CompletedAt,
		CreatedAt:     task.CreatedAt,
		UpdatedAt:     task.UpdatedAt,
	}
}

// Helper function to parse user ID from context
func (c *TaskController) parseUserID(userID interface{}) uuid.UUID {
	if id, ok := userID.(uuid.UUID); ok {
		return id
	}
	if idStr, ok := userID.(string); ok {
		if parsed, err := uuid.Parse(idStr); err == nil {
			return parsed
		}
	}
	return uuid.Nil
}
