package services

import (
	"context"
	"fmt"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/providers/tektoncd"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TaskService is a service that manages tasks.
type TaskService struct {
	taskRepo     repositories.TaskRepository
	taskRunRepo  repositories.TaskRunRepository
	tektonClient tektoncd.TektonClient
}

// NewTaskService creates a new TaskService.
func NewTaskService(taskRepo repositories.TaskRepository, taskRunRepo repositories.TaskRunRepository, tektonClient tektoncd.TektonClient) *TaskService {
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
func (s *TaskService) Run(ctx context.Context, taskID uuid.UUID, taskRunID uuid.UUID) error {
	task, err := s.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return err
	}

	if task.Steps == nil || len(task.Steps) == 0 {
		return fmt.Errorf("task %s has no steps", task.ID)
	}

	var taskRun *entities.TaskRun

	if taskRunID == uuid.Nil {
		taskRunID, err := utils.GenerateID()
		if err != nil {
			return err
		}

		taskRun, err = entities.NewTaskRun(
			&entities.TaskRunConfig{
				ID:     taskRunID,
				TaskID: task.ID,
				TaskRun: &tekton.TaskRun{
					ObjectMeta: metav1.ObjectMeta{
						GenerateName: fmt.Sprintf("%s-", taskRunID),
						Namespace:    task.Namespace,
					},
					Spec: tekton.TaskRunSpec{
						TaskSpec: &tekton.TaskSpec{
							Steps:  task.Steps,
							Params: task.Params,
						},
					},
				},
			},
		)
		if err != nil {
			return err
		}
	} else {
		// TODO: Check if taskRun already exists and update it instead of creating a new one.
		// this is used to re-run a taskRun.
		log.Debug().Msgf("taskRunID is not empty %s", taskRunID.String())
		return fmt.Errorf("retry or re-run not implemented")
	}

	if err := s.taskRunRepo.Create(ctx, taskRun); err != nil {
		return err
	}

	if err := s.tektonClient.RunTaskRun(ctx, taskRun.TaskRun); err != nil {
		log.Error().Err(err).Msg("failed to run task")
		if err := s.taskRunRepo.Update(ctx, taskRun); err != nil {
			return err
		}
		return err
	}

	// spin a goroutine to watch the taskRun
	go func() {
		ctx := context.Background()

		event, err := s.tektonClient.WatchTaskRun(ctx, taskRun.TaskRun, taskRun.ID)
		if err != nil {
			log.Error().Err(err).Msg("failed to watch task run")
			return
		}

		for e := range event {
			obj := e.Object.(*tekton.TaskRun)
			taskRun.Status = obj.Status
			s.taskRunRepo.Update(ctx, taskRun)
		}
	}()

	return nil
}

// FindTaskRuns returns all task runs for a task.
func (s *TaskService) FindTaskRuns(ctx context.Context, taskID uuid.UUID) ([]*entities.TaskRun, error) {
	return s.taskRunRepo.FindByTaskID(ctx, taskID)
}

// FindTaskRunByID returns a task run by ID.
func (s *TaskService) FindTaskRunByID(ctx context.Context, taskID uuid.UUID, taskRunID uuid.UUID) (*entities.TaskRun, error) {
	return s.taskRunRepo.FindByID(ctx, taskRunID)
}

// FindTaskRunLogs returns a stream of logs for a task run.
func (s *TaskService) FindTaskRunLogs(ctx context.Context, taskID uuid.UUID, taskRunID uuid.UUID) (string, error) {
	taskRun, err := s.FindTaskRunByID(ctx, taskID, taskRunID)
	if err != nil {
		return "", err
	}

	return s.tektonClient.GetTaskRunLogs(ctx, taskRun.TaskRun)
}

// FindTaskRunStatus returns the status of a task run.
func (s *TaskService) FindTaskRunStatus(ctx context.Context, taskID uuid.UUID, taskRunID uuid.UUID) (tekton.TaskRunStatus, error) {
	taskRun, err := s.FindTaskRunByID(ctx, taskID, taskRunID)
	if err != nil {
		return taskRun.Status, err
	}

	return taskRun.Status, nil
}
