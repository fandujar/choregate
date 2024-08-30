package postgres

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

type PostgresTaskRunRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresTaskRunRepository(ctx context.Context) (*PostgresTaskRunRepository, error) {
	pool, err := setupPool(ctx)
	if err != nil {
		return nil, err
	}

	return &PostgresTaskRunRepository{
		pool: pool,
	}, nil
}

func (r *PostgresTaskRunRepository) FindAll(ctx context.Context) ([]*entities.TaskRun, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// TODO: Implement pagination
	rows, err := conn.Query(ctx, "SELECT * FROM task_runs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskRuns := make([]*entities.TaskRun, 0)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		taskRun := &entities.TaskRun{
			TaskRunConfig: &entities.TaskRunConfig{
				TaskRun: &tekton.TaskRun{},
			},
		}
		err := rows.Scan(&taskRun.ID, &taskRun.TaskID, &taskRun.Status, &taskRun.Spec)
		if err != nil {
			return nil, err
		}

		taskRuns = append(taskRuns, taskRun)
	}

	return taskRuns, nil
}

func (r *PostgresTaskRunRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.TaskRun, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	taskRun := &entities.TaskRun{
		TaskRunConfig: &entities.TaskRunConfig{
			TaskRun: &tekton.TaskRun{},
		},
	}
	err = conn.QueryRow(ctx, "SELECT * FROM task_runs WHERE id = $1", id).Scan(&taskRun.ID, &taskRun.TaskID, &taskRun.Status, &taskRun.Spec)
	if err != nil {
		return nil, err
	}

	return taskRun, nil
}

func (r *PostgresTaskRunRepository) FindByTaskID(ctx context.Context, taskID uuid.UUID) ([]*entities.TaskRun, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT * FROM task_runs WHERE task_id = $1", taskID)
	if err != nil {
		return nil, err
	}

	taskRuns := make([]*entities.TaskRun, 0)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		taskRun := &entities.TaskRun{
			TaskRunConfig: &entities.TaskRunConfig{
				TaskRun: &tekton.TaskRun{},
			},
		}
		err := rows.Scan(&taskRun.ID, &taskRun.TaskID, &taskRun.Status, &taskRun.Spec)
		if err != nil {
			return nil, err
		}

		taskRuns = append(taskRuns, taskRun)
	}

	return taskRuns, nil

}

func (r *PostgresTaskRunRepository) Create(ctx context.Context, taskRun *entities.TaskRun) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "INSERT INTO task_runs (id, task_id, status, spec) VALUES ($1, $2, $3, $4)", taskRun.ID, taskRun.TaskID, taskRun.Status, taskRun.Spec)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresTaskRunRepository) Update(ctx context.Context, taskRun *entities.TaskRun) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "UPDATE task_runs SET status = $2, spec = $3 WHERE id = $1", taskRun.ID, taskRun.Status, taskRun.Spec)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresTaskRunRepository) Delete(ctx context.Context, id uuid.UUID) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "DELETE FROM task_runs WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
