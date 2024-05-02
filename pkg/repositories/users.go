package repositories

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// UserRepository is a repository that manages users.
type UserRepository interface {
	// FindAll returns all users.
	FindAll(ctx context.Context) ([]*entities.User, error)
	// FindByID returns a user by ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	// Create creates a new user.
	Create(ctx context.Context, user *entities.User) error
	// Update updates a user.
	Update(ctx context.Context, user *entities.User) error
	// Delete deletes a user by ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
