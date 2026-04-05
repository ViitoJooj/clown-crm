package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/repository"
	"github.com/google/uuid"
)

type TaskService struct {
	taskRepo         repository.TaskRepository
	notificationRepo repository.NotificationRepository
}

func NewTaskService(taskRepo repository.TaskRepository, notificationRepo repository.NotificationRepository) *TaskService {
	return &TaskService{
		taskRepo:         taskRepo,
		notificationRepo: notificationRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task *domain.Task, createdBy uuid.UUID) error {
	// Validate
	if err := task.Validate(); err != nil {
		return err
	}
	
	// Set defaults
	task.ID = uuid.New()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.CreatedBy = &createdBy
	
	if task.Status == "" {
		task.Status = "pending"
	}
	
	if task.Priority == "" {
		task.Priority = "medium"
	}
	
	// Create task
	if err := s.taskRepo.Create(ctx, task); err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	
	// Create notification for assigned user if different from creator
	if task.AssignedTo != nil && *task.AssignedTo != createdBy {
		notification := &domain.Notification{
			ID:               uuid.New(),
			UserID:           *task.AssignedTo,
			NotificationType: "task_assigned",
			Title:            "New Task Assigned",
			Message:          strPtr(fmt.Sprintf("You have been assigned a new task: %s", task.Title)),
			RelatedToType:    strPtr("task"),
			RelatedToID:      &task.ID,
			CreatedAt:        time.Now(),
		}
		s.notificationRepo.Create(ctx, notification)
	}
	
	return nil
}

// GetTask retrieves a task by ID
func (s *TaskService) GetTask(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	return s.taskRepo.GetByID(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *domain.Task, updatedBy uuid.UUID) error {
	// Validate
	if err := task.Validate(); err != nil {
		return err
	}
	
	// Get existing to track changes
	existing, err := s.taskRepo.GetByID(ctx, task.ID)
	if err != nil {
		return err
	}
	
	task.UpdatedAt = time.Now()
	
	// Update task
	if err := s.taskRepo.Update(ctx, task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	
	// Notify on status change to completed
	if existing.Status != "completed" && task.Status == "completed" && task.AssignedTo != nil {
		notification := &domain.Notification{
			ID:               uuid.New(),
			UserID:           *task.AssignedTo,
			NotificationType: "task_completed",
			Title:            "Task Completed",
			Message:          strPtr(fmt.Sprintf("Task '%s' has been marked as completed", task.Title)),
			RelatedToType:    strPtr("task"),
			RelatedToID:      &task.ID,
			CreatedAt:        time.Now(),
		}
		s.notificationRepo.Create(ctx, notification)
	}
	
	// Notify on reassignment
	if existing.AssignedTo != task.AssignedTo && task.AssignedTo != nil {
		notification := &domain.Notification{
			ID:               uuid.New(),
			UserID:           *task.AssignedTo,
			NotificationType: "task_assigned",
			Title:            "Task Assigned to You",
			Message:          strPtr(fmt.Sprintf("You have been assigned the task: %s", task.Title)),
			RelatedToType:    strPtr("task"),
			RelatedToID:      &task.ID,
			CreatedAt:        time.Now(),
		}
		s.notificationRepo.Create(ctx, notification)
	}
	
	return nil
}

// CompleteTask marks a task as completed
func (s *TaskService) CompleteTask(ctx context.Context, taskID, completedBy uuid.UUID) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	
	task.Complete()
	task.UpdatedAt = time.Now()
	
	if err := s.taskRepo.Update(ctx, task); err != nil {
		return err
	}
	
	// Notify task creator if different
	if task.CreatedBy != nil && *task.CreatedBy != completedBy {
		notification := &domain.Notification{
			ID:               uuid.New(),
			UserID:           *task.CreatedBy,
			NotificationType: "task_completed",
			Title:            "Task Completed",
			Message:          strPtr(fmt.Sprintf("Task '%s' has been completed", task.Title)),
			RelatedToType:    strPtr("task"),
			RelatedToID:      &taskID,
			CreatedAt:        time.Now(),
		}
		s.notificationRepo.Create(ctx, notification)
	}
	
	return nil
}

// CancelTask marks a task as cancelled
func (s *TaskService) CancelTask(ctx context.Context, taskID, cancelledBy uuid.UUID) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	
	task.Cancel()
	task.UpdatedAt = time.Now()
	
	return s.taskRepo.Update(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return s.taskRepo.Delete(ctx, id)
}

// ListTasks lists tasks with filters
func (s *TaskService) ListTasks(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Task, int, error) {
	return s.taskRepo.List(ctx, filters, page, pageSize)
}

// GetUserTasks retrieves all tasks assigned to a user
func (s *TaskService) GetUserTasks(ctx context.Context, userID uuid.UUID, status string, page, pageSize int) ([]*domain.Task, int, error) {
	filters := map[string]interface{}{
		"assigned_to": userID,
	}
	
	if status != "" {
		filters["status"] = status
	}
	
	return s.taskRepo.List(ctx, filters, page, pageSize)
}

// GetOverdueTasks retrieves overdue tasks
func (s *TaskService) GetOverdueTasks(ctx context.Context, userID *uuid.UUID) ([]*domain.Task, error) {
	filters := map[string]interface{}{
		"status": "pending",
		"overdue": true,
	}
	
	if userID != nil {
		filters["assigned_to"] = *userID
	}
	
	tasks, _, err := s.taskRepo.List(ctx, filters, 1, 1000)
	return tasks, err
}

// GetUpcomingTasks retrieves tasks due soon (within next 24 hours)
func (s *TaskService) GetUpcomingTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	filters := map[string]interface{}{
		"assigned_to": userID,
		"status":      "pending",
		"due_within":  24, // hours
	}
	
	tasks, _, err := s.taskRepo.List(ctx, filters, 1, 1000)
	return tasks, err
}

// AssignTask assigns a task to a user
func (s *TaskService) AssignTask(ctx context.Context, taskID, userID, assignedBy uuid.UUID) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	
	oldAssignedTo := task.AssignedTo
	task.AssignedTo = &userID
	task.UpdatedAt = time.Now()
	
	if err := s.taskRepo.Update(ctx, task); err != nil {
		return err
	}
	
	// Notify new assignee
	notification := &domain.Notification{
		ID:               uuid.New(),
		UserID:           userID,
		NotificationType: "task_assigned",
		Title:            "New Task Assigned",
		Message:          strPtr(fmt.Sprintf("You have been assigned the task: %s", task.Title)),
		RelatedToType:    strPtr("task"),
		RelatedToID:      &taskID,
		CreatedAt:        time.Now(),
	}
	s.notificationRepo.Create(ctx, notification)
	
	// Notify old assignee if exists
	if oldAssignedTo != nil && *oldAssignedTo != userID {
		notification := &domain.Notification{
			ID:               uuid.New(),
			UserID:           *oldAssignedTo,
			NotificationType: "task_assigned",
			Title:            "Task Reassigned",
			Message:          strPtr(fmt.Sprintf("Task '%s' has been reassigned to another user", task.Title)),
			RelatedToType:    strPtr("task"),
			RelatedToID:      &taskID,
			CreatedAt:        time.Now(),
		}
		s.notificationRepo.Create(ctx, notification)
	}
	
	return nil
}

// SetReminder sets or updates a reminder for a task
func (s *TaskService) SetReminder(ctx context.Context, taskID uuid.UUID, reminderAt time.Time) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	
	task.ReminderAt = &reminderAt
	task.UpdatedAt = time.Now()
	
	return s.taskRepo.Update(ctx, task)
}

// CheckAndSendReminders checks for tasks with reminders due and sends notifications
func (s *TaskService) CheckAndSendReminders(ctx context.Context) error {
	// Get tasks with reminders due within the next 5 minutes
	filters := map[string]interface{}{
		"reminder_due": true,
		"status":       "pending",
	}
	
	tasks, _, err := s.taskRepo.List(ctx, filters, 1, 1000)
	if err != nil {
		return err
	}
	
	for _, task := range tasks {
		if task.AssignedTo != nil {
			// Send reminder notification
			notification := &domain.Notification{
				ID:               uuid.New(),
				UserID:           *task.AssignedTo,
				NotificationType: "task_due",
				Title:            "Task Reminder",
				Message:          strPtr(fmt.Sprintf("Reminder: Task '%s' is due soon", task.Title)),
				RelatedToType:    strPtr("task"),
				RelatedToID:      &task.ID,
				CreatedAt:        time.Now(),
			}
			s.notificationRepo.Create(ctx, notification)
			
			// Clear the reminder so we don't send it again
			task.ReminderAt = nil
			task.UpdatedAt = time.Now()
			s.taskRepo.Update(ctx, task)
		}
	}
	
	return nil
}

// GetTaskStats retrieves task statistics for a user
func (s *TaskService) GetTaskStats(ctx context.Context, userID uuid.UUID) (map[string]int, error) {
	stats := make(map[string]int)
	
	// All tasks
	allTasks, totalCount, err := s.taskRepo.List(ctx, map[string]interface{}{"assigned_to": userID}, 1, 10000)
	if err != nil {
		return nil, err
	}
	
	stats["total"] = totalCount
	
	// Count by status
	for _, task := range allTasks {
		switch task.Status {
		case "pending":
			stats["pending"]++
		case "in_progress":
			stats["in_progress"]++
		case "completed":
			stats["completed"]++
		case "cancelled":
			stats["cancelled"]++
		}
		
		// Count overdue
		if task.IsOverdue() {
			stats["overdue"]++
		}
	}
	
	return stats, nil
}
