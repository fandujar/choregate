package repositories

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// TriggerRepository is a repository that manages triggers.
type TriggerRepository interface {
	// FindAll returns all triggers.
	FindAll(ctx context.Context) ([]*entities.Trigger, error)
	// FindByID returns a trigger by ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Trigger, error)
	// Create creates a new trigger.
	Create(ctx context.Context, trigger *entities.Trigger) error
	// Update updates a trigger.
	Update(ctx context.Context, trigger *entities.Trigger) error
	// Delete deletes a trigger by ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
