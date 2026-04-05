package domain

import "fmt"

// ErrInvalidInput represents a validation error
type ErrInvalidInput struct {
	Field   string
	Message string
}

func (e ErrInvalidInput) Error() string {
	return fmt.Sprintf("invalid input for %s: %s", e.Field, e.Message)
}

// ErrNotFound represents a not found error
type ErrNotFound struct {
	Entity string
	ID     string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Entity, e.ID)
}

// ErrUnauthorized represents an unauthorized access error
type ErrUnauthorized struct {
	Message string
}

func (e ErrUnauthorized) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.Message)
}

// ErrConflict represents a conflict error (e.g., duplicate entry)
type ErrConflict struct {
	Message string
}

func (e ErrConflict) Error() string {
	return fmt.Sprintf("conflict: %s", e.Message)
}
