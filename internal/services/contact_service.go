package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/repository"
	"github.com/google/uuid"
)

type ContactService struct {
	contactRepo repository.ContactRepository
	activityRepo repository.ActivityRepository
}

func NewContactService(contactRepo repository.ContactRepository, activityRepo repository.ActivityRepository) *ContactService {
	return &ContactService{
		contactRepo: contactRepo,
		activityRepo: activityRepo,
	}
}

func (s *ContactService) CreateContact(ctx context.Context, contact *domain.Contact, createdBy uuid.UUID) error {
	// Validate
	if err := contact.Validate(); err != nil {
		return err
	}
	
	// Set defaults
	contact.ID = uuid.New()
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	contact.Status = "lead" // Default status
	
	if contact.OwnerID == nil {
		contact.OwnerID = &createdBy
	}
	
	// Create contact
	if err := s.contactRepo.Create(ctx, contact); err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}
	
	// Log activity
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "status_change",
		Title:        "Contact Created",
		Description:  strPtr(fmt.Sprintf("Contact %s created", contact.FullName())),
		ContactID:    &contact.ID,
		UserID:       &createdBy,
		CreatedAt:    time.Now(),
	}
	s.activityRepo.Create(ctx, activity)
	
	return nil
}

// GetContact retrieves a contact by ID
func (s *ContactService) GetContact(ctx context.Context, id uuid.UUID) (*domain.Contact, error) {
	return s.contactRepo.GetByID(ctx, id)
}

func (s *ContactService) UpdateContact(ctx context.Context, contact *domain.Contact, updatedBy uuid.UUID) error {
	// Validate
	if err := contact.Validate(); err != nil {
		return err
	}
	
	// Get existing contact to track changes
	existing, err := s.contactRepo.GetByID(ctx, contact.ID)
	if err != nil {
		return err
	}
	
	// Update timestamp
	contact.UpdatedAt = time.Now()
	
	// Update contact
	if err := s.contactRepo.Update(ctx, contact); err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}
	
	// Log status change if changed
	if existing.Status != contact.Status {
		activity := &domain.Activity{
			ID:           uuid.New(),
			ActivityType: "status_change",
			Title:        "Status Changed",
			Description:  strPtr(fmt.Sprintf("Status changed from %s to %s", existing.Status, contact.Status)),
			ContactID:    &contact.ID,
			UserID:       &updatedBy,
			Metadata: map[string]interface{}{
				"old_status": existing.Status,
				"new_status": contact.Status,
			},
			CreatedAt: time.Now(),
		}
		s.activityRepo.Create(ctx, activity)
	}
	
	return nil
}

func (s *ContactService) DeleteContact(ctx context.Context, id uuid.UUID) error {
	return s.contactRepo.Delete(ctx, id)
}

// ListContacts lists contacts with filters
func (s *ContactService) ListContacts(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Contact, int, error) {
	return s.contactRepo.List(ctx, filters, page, pageSize)
}

// SearchContacts searches contacts by term
func (s *ContactService) SearchContacts(ctx context.Context, searchTerm string, page, pageSize int) ([]*domain.Contact, int, error) {
	return s.contactRepo.Search(ctx, searchTerm, page, pageSize)
}

// AssignContact assigns a contact to a user
func (s *ContactService) AssignContact(ctx context.Context, contactID, userID, assignedBy uuid.UUID) error {
	contact, err := s.contactRepo.GetByID(ctx, contactID)
	if err != nil {
		return err
	}
	
	oldAssignedTo := contact.AssignedTo
	contact.AssignedTo = &userID
	contact.UpdatedAt = time.Now()
	
	if err := s.contactRepo.Update(ctx, contact); err != nil {
		return err
	}
	
	// Log activity
	var desc string
	if oldAssignedTo == nil {
		desc = fmt.Sprintf("Contact assigned to user %s", userID)
	} else {
		desc = fmt.Sprintf("Contact reassigned from %s to %s", *oldAssignedTo, userID)
	}
	
	activity := &domain.Activity{
		ID:           uuid.New(),
		ActivityType: "assignment",
		Title:        "Contact Assigned",
		Description:  &desc,
		ContactID:    &contactID,
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

// AddTagToContact adds a tag to a contact
func (s *ContactService) AddTagToContact(ctx context.Context, contactID uuid.UUID, tag string) error {
	contact, err := s.contactRepo.GetByID(ctx, contactID)
	if err != nil {
		return err
	}
	
	// Check if tag already exists
	for _, existingTag := range contact.Tags {
		if existingTag == tag {
			return nil // Already has this tag
		}
	}
	
	contact.Tags = append(contact.Tags, tag)
	contact.UpdatedAt = time.Now()
	
	return s.contactRepo.Update(ctx, contact)
}

// RemoveTagFromContact removes a tag from a contact
func (s *ContactService) RemoveTagFromContact(ctx context.Context, contactID uuid.UUID, tag string) error {
	contact, err := s.contactRepo.GetByID(ctx, contactID)
	if err != nil {
		return err
	}
	
	// Remove tag
	var newTags []string
	for _, existingTag := range contact.Tags {
		if existingTag != tag {
			newTags = append(newTags, existingTag)
		}
	}
	
	contact.Tags = newTags
	contact.UpdatedAt = time.Now()
	
	return s.contactRepo.Update(ctx, contact)
}

// UpdateLastContact updates the last contact timestamp
func (s *ContactService) UpdateLastContact(ctx context.Context, contactID uuid.UUID) error {
	contact, err := s.contactRepo.GetByID(ctx, contactID)
	if err != nil {
		return err
	}
	
	now := time.Now()
	contact.LastContactAt = &now
	contact.UpdatedAt = now
	
	return s.contactRepo.Update(ctx, contact)
}

// GetContactActivities retrieves activities for a contact
func (s *ContactService) GetContactActivities(ctx context.Context, contactID uuid.UUID, page, pageSize int) ([]*domain.Activity, int, error) {
	return s.activityRepo.GetByContactID(ctx, contactID, page*pageSize, pageSize)
}

// Helper function
func strPtr(s string) *string {
	return &s
}
