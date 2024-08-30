package postgres

import (
	"context"
	"os"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresTaskRepository is a postgres task repository.
type PostgresTaskRepository struct {
	pool *pgxpool.Pool
}

// FindAll returns all tasks in the repository.
func (r *PostgresTaskRepository) FindAll(ctx context.Context, scope *entities.TaskScope) ([]*entities.Task, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

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
		var task entities.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Namespace, &task.Description, &task.Timeout, &task.Steps, &task.Params)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// FindByID returns the task with the specified ID.
func (r *PostgresTaskRepository) FindByID(ctx context.Context, id uuid.UUID, scope *entities.TaskScope) (*entities.Task, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	var task entities.Task
	err = conn.QueryRow(ctx, "SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.Namespace, &task.Description, &task.Timeout, &task.Steps, &task.Params)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Create adds a new task to the repository.
func (r *PostgresTaskRepository) Create(ctx context.Context, task *entities.Task) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "INSERT INTO tasks (id, name, namespace, description, timeout, steps, params) VALUES ($1, $2, $3, $4, $5, $6, $7)", task.ID, task.Name, task.Namespace, task.Description, task.Timeout, task.Steps, task.Params)
	return err
}

// Update updates an existing task in the repository.
func (r *PostgresTaskRepository) Update(ctx context.Context, task *entities.Task) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "UPDATE tasks SET name = $2, namespace = $3, description = $4, timeout = $5, steps = $6, params = $7 WHERE id = $1", task.ID, task.Name, task.Namespace, task.Description, task.Timeout, task.Steps, task.Params)
	return err
}

// Delete removes a task from the repository.
func (r *PostgresTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	return err
}

// NewPostgresTaskRepository creates a new instance of PostgresTaskRepository.
func NewPostgresTaskRepository(ctx context.Context) (*PostgresTaskRepository, error) {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DATABASE_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DATABASE_USER")
	if user == "" {
		user = "choregate"
	}

	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		password = "choregate"
	}

	database := os.Getenv("DATABASE_NAME")
	if database == "" {
		database = "choregate"
	}

	databaseURL := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + database
	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)

	return &PostgresTaskRepository{
		pool: pool,
	}, nil
}
