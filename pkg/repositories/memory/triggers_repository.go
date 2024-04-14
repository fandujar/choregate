package memory

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// TriggerRepository is a repository that manages triggers in memory.
type TriggerRepository struct {
	triggers map[uuid.UUID]*entities.Trigger
}

// FindAll returns all triggers.
func (r *TriggerRepository) FindAll(ctx context.Context) ([]*entities.Trigger, error) {
	triggers := make([]*entities.Trigger, 0, len(r.triggers))
	for _, trigger := range r.triggers {
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}

// FindByID returns a trigger by ID.
func (r *TriggerRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Trigger, error) {
	trigger, ok := r.triggers[id]
	if !ok {
		return nil, entities.ErrTriggerNotFound{}
	}
	return trigger, nil
}

// Create creates a new trigger.
func (r *TriggerRepository) Create(ctx context.Context, trigger *entities.Trigger) error {
	if _, ok := r.triggers[trigger.ID]; ok {
		return entities.ErrTriggerAlreadyExists{}
	}
	r.triggers[trigger.ID] = trigger
	return nil
}

// Update updates a trigger.
func (r *TriggerRepository) Update(ctx context.Context, trigger *entities.Trigger) error {
	if _, ok := r.triggers[trigger.ID]; !ok {
		return entities.ErrTriggerNotFound{}
	}
	r.triggers[trigger.ID] = trigger
	return nil
}

// Delete deletes a trigger by ID.
func (r *TriggerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, ok := r.triggers[id]; !ok {
		return entities.ErrTriggerNotFound{}
	}
	delete(r.triggers, id)
	return nil
}

// NewInMemoryTriggerRepository creates a new in-memory trigger repository.
func NewInMemoryTriggerRepository() *TriggerRepository {
	return &TriggerRepository{
		triggers: make(map[uuid.UUID]*entities.Trigger),
	}
}
