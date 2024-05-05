package memory

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// InMemoryTaskRunRepository is an in-memory task run repository.
type InMemoryTaskRunRepository struct {
	taskRuns map[string]*entities.TaskRun
}

// NewInMemoryTaskRunRepository creates a new instance of InMemoryTaskRunRepository.
func NewInMemoryTaskRunRepository() *InMemoryTaskRunRepository {
	return &InMemoryTaskRunRepository{
		taskRuns: make(map[string]*entities.TaskRun),
	}
}

// FindAll returns all task runs in the repository.
func (r *InMemoryTaskRunRepository) FindAll(ctx context.Context) ([]*entities.TaskRun, error) {
	taskRuns := make([]*entities.TaskRun, 0, len(r.taskRuns))
	for _, taskRun := range r.taskRuns {
		taskRuns = append(taskRuns, taskRun)
	}
	return taskRuns, nil
}

// FindByID returns the task run with the specified ID.
func (r *InMemoryTaskRunRepository) FindByID(ctx context.Context, id string) (*entities.TaskRun, error) {
	taskRun, ok := r.taskRuns[id]
	if !ok {
		return nil, entities.ErrTaskRunNotFound{}
	}
	return taskRun, nil
}

// Create adds a new task run to the repository.
func (r *InMemoryTaskRunRepository) Create(ctx context.Context, taskRun *entities.TaskRun) error {
	if _, ok := r.taskRuns[taskRun.ID]; ok {
		return entities.ErrTaskRunAlreadyExists{}
	}
	r.taskRuns[taskRun.ID] = taskRun
	return nil
}

// Update updates an existing task run in the repository.
func (r *InMemoryTaskRunRepository) Update(ctx context.Context, taskRun *entities.TaskRun) error {
	if _, ok := r.taskRuns[taskRun.ID]; !ok {
		return entities.ErrTaskRunNotFound{}
	}
	r.taskRuns[taskRun.ID] = taskRun
	return nil
}

// Delete removes a task run from the repository.
func (r *InMemoryTaskRunRepository) Delete(ctx context.Context, id string) error {
	if _, ok := r.taskRuns[id]; !ok {
		return entities.ErrTaskRunNotFound{}
	}
	delete(r.taskRuns, id)
	return nil
}

// FindByTaskID returns all task runs for a task.
func (r *InMemoryTaskRunRepository) FindByTaskID(ctx context.Context, taskID uuid.UUID) ([]*entities.TaskRun, error) {
	taskRuns := make([]*entities.TaskRun, 0)
	for _, taskRun := range r.taskRuns {
		if taskRun.TaskID == taskID {
			taskRuns = append(taskRuns, taskRun)
		}
	}
	return taskRuns, nil
}
