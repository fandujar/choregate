package postgres

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

// PostgresTaskRepository is a postgres task repository.
type PostgresTaskRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresTaskRepository creates a new instance of PostgresTaskRepository.
func NewPostgresTaskRepository(ctx context.Context) (*PostgresTaskRepository, error) {
	pool, err := setupPool(ctx)
	if err != nil {
		return nil, err
	}

	return &PostgresTaskRepository{
		pool: pool,
	}, nil
}

// FindAll returns all tasks in the repository.
func (r *PostgresTaskRepository) FindAll(ctx context.Context, scope *entities.TaskScope) ([]*entities.Task, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*entities.Task, 0)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		task := &entities.Task{
			TaskConfig: &entities.TaskConfig{
				TaskScope: &entities.TaskScope{},
				TaskSpec:  &tekton.TaskSpec{},
			},
		}
		err := rows.Scan(&task.ID, &task.Name, &task.Namespace, &task.Description, &task.Timeout, &task.TaskScope, &task.TaskSpec)
		if err != nil {
			return nil, err
		}

		if task.TaskScope == nil {
			tasks = append(tasks, task)
			continue
		}

		if task.TaskScope.Owner == scope.Owner {
			tasks = append(tasks, task)
			continue
		}

		// TODO: spike if this filter should be done in the database
		for _, organization := range task.TaskScope.Organizations {
			for _, org := range scope.Organizations {
				if organization.ID == org.ID {
					tasks = append(tasks, task)
					continue
				}
			}
		}

		// TODO: spike if this filter should be done in the database
		for _, team := range task.TaskScope.Teams {
			for _, t := range scope.Teams {
				if team.ID == t.ID {
					tasks = append(tasks, task)
					continue
				}
			}
		}

	}

	return tasks, nil
}

// FindByID returns the task with the specified ID.
func (r *PostgresTaskRepository) FindByID(ctx context.Context, id uuid.UUID, scope *entities.TaskScope) (*entities.Task, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	task := &entities.Task{
		TaskConfig: &entities.TaskConfig{
			TaskScope: &entities.TaskScope{},
			TaskSpec:  &tekton.TaskSpec{},
		},
	}
	err = conn.QueryRow(ctx, "SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.Namespace, &task.Description, &task.Timeout, &task.TaskScope, &task.TaskSpec)
	if err != nil {
		return nil, err
	}

	if task.TaskScope == nil {
		return task, nil
	}

	if task.TaskScope.Owner == scope.Owner {
		return task, nil
	}

	// TODO: spike if this filter should be done in the database
	for _, organization := range task.TaskScope.Organizations {
		for _, org := range scope.Organizations {
			if organization.ID == org.ID {
				return task, nil
			}
		}
	}

	// TODO: spike if this filter should be done in the database
	for _, team := range task.TaskScope.Teams {
		for _, t := range scope.Teams {
			if team.ID == t.ID {
				return task, nil
			}
		}
	}

	return nil, entities.ErrTaskNotFound{}
}

// Create adds a new task to the repository.
func (r *PostgresTaskRepository) Create(ctx context.Context, task *entities.Task) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO tasks (id, name, namespace, description, timeout, task_scope, task_spec) VALUES ($1, $2, $3, $4, $5, $6, $7)", task.ID, task.Name, task.Namespace, task.Description, task.Timeout, task.TaskScope, task.TaskSpec)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing task in the repository.
func (r *PostgresTaskRepository) Update(ctx context.Context, task *entities.Task) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE tasks SET name = $2, namespace = $3, description = $4, timeout = $5, task_scope = $6, task_spec = $7 WHERE id = $1", task.ID, task.Name, task.Namespace, task.Description, task.Timeout, task.TaskScope, task.TaskSpec)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a task from the repository.
func (r *PostgresTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return err
}
