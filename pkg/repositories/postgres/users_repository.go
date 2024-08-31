package postgres

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(ctx context.Context) (*PostgresUserRepository, error) {
	pool, err := setupPool(ctx)
	if err != nil {
		return nil, err
	}

	return &PostgresUserRepository{
		pool: pool,
	}, nil
}

func (r *PostgresUserRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id, slug, name, email, system_role FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]*entities.User, 0)
	for rows.Next() {
		user := &entities.User{
			UserConfig: &entities.UserConfig{},
		}
		err := rows.Scan(&user.ID, &user.Slug, &user.Name, &user.Email, &user.SystemRole)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	user := &entities.User{
		UserConfig: &entities.UserConfig{},
	}
	err = conn.QueryRow(ctx, "SELECT id, slug, name, email, system_role, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Slug, &user.Name, &user.Email, &user.SystemRole, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	user := &entities.User{
		UserConfig: &entities.UserConfig{},
	}
	err = conn.QueryRow(ctx, "SELECT id, slug, name, email, system_role, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Slug, &user.Name, &user.Email, &user.SystemRole, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *entities.User) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "INSERT INTO users (id, slug, name, email, system_role, password) VALUES ($1, $2, $3, $4, $5, $6)", user.ID, user.Slug, user.Name, user.Email, user.SystemRole, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *entities.User) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "UPDATE users SET slug = $1, name = $2, email = $3, system_role = $4, password = $5 WHERE id = $6", user.Slug, user.Name, user.Email, user.SystemRole, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
