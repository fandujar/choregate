package services

import (
	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"

	"context"
)

// TriggerService is a service that manages triggers.
type TriggerService struct {
	repo repositories.TriggerRepository
}

// NewTriggerService creates a new TriggerService.
func NewTriggerService(repo repositories.TriggerRepository) *TriggerService {
	return &TriggerService{
		repo: repo,
	}
}

// FindAll returns all triggers.
func (s *TriggerService) FindAll(ctx context.Context) ([]*entities.Trigger, error) {
	return s.repo.FindAll(ctx)
}

// FindByID returns a trigger by ID.
func (s *TriggerService) FindByID(ctx context.Context, id uuid.UUID) (*entities.Trigger, error) {
	return s.repo.FindByID(ctx, id)
}

// Create creates a new trigger.
func (s *TriggerService) Create(ctx context.Context, trigger *entities.Trigger) error {
	return s.repo.Create(ctx, trigger)
}

// Update updates a trigger.
func (s *TriggerService) Update(ctx context.Context, trigger *entities.Trigger) error {
	return s.repo.Update(ctx, trigger)
}

// Delete deletes a trigger by ID.
func (s *TriggerService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
