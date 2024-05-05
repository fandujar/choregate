package repositories

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

type TaskRunRepository interface {
	// FindAll returns all task runs.
	FindAll(ctx context.Context) ([]*entities.TaskRun, error)
	// FindByID returns a task run by ID.
	FindByID(ctx context.Context, id string) (*entities.TaskRun, error)
	// FindByTaskID returns all task runs for a task.
	FindByTaskID(ctx context.Context, taskID uuid.UUID) ([]*entities.TaskRun, error)
	// Create creates a new task run.
	Create(ctx context.Context, taskRun *entities.TaskRun) error
	// Update updates a task run.
	Update(ctx context.Context, taskRun *entities.TaskRun) error
	// Delete deletes a task run by ID.
	Delete(ctx context.Context, id string) error
}
