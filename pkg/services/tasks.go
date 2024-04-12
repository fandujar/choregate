package services

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
)

// TaskService is a service that manages tasks.
type TaskService struct {
	repo repositories.TaskRepository
}

// NewTaskService creates a new TaskService.
func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// FindAll returns all tasks.
func (s *TaskService) FindAll(ctx context.Context) ([]*entities.Task, error) {
	return s.repo.FindAll(ctx)
}

// FindByID returns a task by ID.
func (s *TaskService) FindByID(ctx context.Context, id uuid.UUID) (*entities.Task, error) {
	return s.repo.FindByID(ctx, id)
}

// Create creates a new task.
func (s *TaskService) Create(ctx context.Context, task *entities.Task) error {
	return s.repo.Create(ctx, task)
}

// Update updates a task.
func (s *TaskService) Update(ctx context.Context, task *entities.Task) error {
	return s.repo.Update(ctx, task)
}

// Delete deletes a task by ID.
func (s *TaskService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
