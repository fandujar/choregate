package services

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
)

// TaskService is a service that manages tasks.
type TaskService struct {
	taskRepo repositories.TaskRepository
}

type TaskRunService struct {
	taskRunRepo repositories.TaskRunRepository
}

// NewTaskService creates a new TaskService.
func NewTaskService(taskRepo repositories.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: repo,
	}
}

// NewTaskRunService creates a new TaskRunService.
func NewTaskRunService(taskRunRepo repositories.TaskRunRepository) *TaskRunService {
	return &TaskRunService{
		taskRunRepo: repo,
	}
}

// FindAll returns all tasks.
func (s *TaskService) FindAll(ctx context.Context) ([]*entities.Task, error) {
	return s.taskRepo.FindAll(ctx)
}

// FindByID returns a task by ID.
func (s *TaskService) FindByID(ctx context.Context, id uuid.UUID) (*entities.Task, error) {
	return s.taskRepo.FindByID(ctx, id)
}

// Create creates a new task.
func (s *TaskService) Create(ctx context.Context, task *entities.Task) error {
	return s.taskRepo.Create(ctx, task)
}

// Update updates a task.
func (s *TaskService) Update(ctx context.Context, task *entities.Task) error {
	return s.taskRepo.Update(ctx, task)
}

// Delete deletes a task by ID.
func (s *TaskService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.taskRepo.Delete(ctx, id)
}

// Run runs a task.
func (s *TaskService) Run(ctx context.Context, id uuid.UUID) error {
	return nil
}
