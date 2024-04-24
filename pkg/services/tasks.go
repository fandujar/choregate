package services

import (
	"context"
	"fmt"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/providers"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TaskService is a service that manages tasks.
type TaskService struct {
	taskRepo     repositories.TaskRepository
	taskRunRepo  repositories.TaskRunRepository
	tektonClient providers.TektonClient
}

// NewTaskService creates a new TaskService.
func NewTaskService(taskRepo repositories.TaskRepository, taskRunRepo repositories.TaskRunRepository, tektonClient providers.TektonClient) *TaskService {
	return &TaskService{
		taskRepo:     taskRepo,
		taskRunRepo:  taskRunRepo,
		tektonClient: tektonClient,
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
	task, err := s.taskRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	taskRun, err := entities.NewTaskRun(
		&entities.TaskRunConfig{
			TaskID: task.ID,
			TaskRun: &tekton.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: fmt.Sprintf("%s-", task.ID.String()),
					Namespace:    task.Namespace,
				},
				Spec: tekton.TaskRunSpec{
					TaskSpec: &tekton.TaskSpec{
						Steps: task.Steps,
					},
				},
			},
		},
	)
	if err != nil {
		return err
	}

	if err := s.taskRunRepo.Create(ctx, taskRun); err != nil {
		return err
	}

	if err := s.tektonClient.RunTask(ctx, taskRun.TaskRun); err != nil {
		log.Error().Err(err).Msg("failed to run task")
		if err := s.taskRunRepo.Update(ctx, taskRun); err != nil {
			return err
		}
		return err
	}

	return nil
}
