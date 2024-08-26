package memory

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// InMemoryUserRepository is an in-memory user repository.
type InMemoryUserRepository struct {
	users map[uuid.UUID]*entities.User
}

// NewInMemoryUserRepository creates a new in-memory user repository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[uuid.UUID]*entities.User),
	}
}

// FindAll returns all users in the repository.
func (r *InMemoryUserRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// FindByID returns the user with the specified ID.
func (r *InMemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, entities.ErrUserNotFound{}
	}
	return user, nil
}

// Create adds a new user to the repository.
func (r *InMemoryUserRepository) Create(ctx context.Context, user *entities.User) error {
	if _, ok := r.users[user.ID]; ok {
		return entities.ErrUserAlreadyExists{}
	}
	r.users[user.ID] = user
	return nil
}

// Update updates an existing user in the repository.
func (r *InMemoryUserRepository) Update(ctx context.Context, user *entities.User) error {
	if _, ok := r.users[user.ID]; !ok {
		return entities.ErrUserNotFound{}
	}
	r.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository.
func (r *InMemoryUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, ok := r.users[id]; !ok {
		return entities.ErrUserNotFound{}
	}
	delete(r.users, id)
	return nil
}
