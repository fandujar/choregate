package memory

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// InMemoryTaskRepository is an in-memory task repository.
type InMemoryTaskRepository struct {
	tasks map[uuid.UUID]*entities.Task
}

// FindAll returns all tasks in the repository.
func (r *InMemoryTaskRepository) FindAll(ctx context.Context, scope *entities.TaskScope) ([]*entities.Task, error) {
	// TODO: Implement FindAll method with context
	tasks := make([]*entities.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		if task.TaskScope == nil {
			tasks = append(tasks, task)
		}

		if task.TaskScope != nil && scope.Owner == task.TaskScope.Owner {
			tasks = append(tasks, task)
		}

		if task.TaskScope != nil && scope != nil {
			for _, organization := range task.TaskScope.Organizations {
				for _, org := range scope.Organizations {
					if organization == org {
						tasks = append(tasks, task)
					}
				}
			}

			for _, team := range task.TaskScope.Teams {
				for _, t := range scope.Teams {
					if team == t {
						tasks = append(tasks, task)
					}
				}
			}
		}
	}
	return tasks, nil
}

// FindByID returns the task with the specified ID.
func (r *InMemoryTaskRepository) FindByID(ctx context.Context, id uuid.UUID, scope *entities.TaskScope) (*entities.Task, error) {
	// TODO: Implement FindByID method with context
	task, ok := r.tasks[id]
	if !ok {
		return nil, entities.ErrTaskNotFound{}
	}

	if task.TaskScope == nil {
		return task, nil
	}

	if task.TaskScope != nil && scope != nil && scope.Owner == task.TaskScope.Owner {
		return task, nil
	}

	if task.TaskScope != nil && scope != nil {
		for _, organization := range task.TaskScope.Organizations {
			for _, org := range scope.Organizations {
				if organization == org {
					return task, nil
				}
			}
		}

		for _, team := range task.TaskScope.Teams {
			for _, t := range scope.Teams {
				if team == t {
					return task, nil
				}
			}
		}
	}

	return nil, entities.ErrTaskNotFound{}
}

// Create adds a new task to the repository.
func (r *InMemoryTaskRepository) Create(ctx context.Context, task *entities.Task) error {
	// TODO: Implement Create method with context

	// FIXME: This is a bug, we should check if the task already exists
	// if _, err := r.FindByID(ctx, task.ID, nil); err != nil {
	// 	return entities.ErrTaskAlreadyExists{}
	// }
	r.tasks[task.ID] = task
	return nil
}

// Update updates an existing task in the repository.
func (r *InMemoryTaskRepository) Update(ctx context.Context, task *entities.Task) error {
	// TODO: Implement Update method with context
	if _, ok := r.tasks[task.ID]; !ok {
		return entities.ErrTaskNotFound{}
	}
	r.tasks[task.ID] = task
	return nil
}

// Delete removes a task from the repository.
func (r *InMemoryTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// TODO: Implement Delete method with context
	if _, ok := r.tasks[id]; !ok {
		return entities.ErrTaskNotFound{}
	}
	delete(r.tasks, id)
	return nil
}

// NewInMemoryTaskRepository creates a new instance of InMemoryTaskRepository.
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[uuid.UUID]*entities.Task),
	}
}
