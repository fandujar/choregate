package repositories

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

type TaskRepository interface {
	// FindAll returns all tasks.
	FindAll(ctx context.Context) ([]*entities.Task, error)
	// FindByID returns a task by ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Task, error)
	// Create creates a new task.
	Create(ctx context.Context, task *entities.Task) error
	// Update updates a task.
	Update(ctx context.Context, task *entities.Task) error
	// Delete deletes a task by ID.
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskRunRepository interface {
	// Create creates a new task run.
	Create(ctx context.Context, taskRun string) error
	// Run runs a task.
	Run(ctx context.Context, task *entities.Task) error
}
