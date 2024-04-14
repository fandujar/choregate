package memory

import (
	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// InMemoryUserRepository is an in-memory user repository.
type InMemoryUserRepository struct {
	users map[uuid.UUID]*entities.User
}

// FindAll returns all users in the repository.
func (r *InMemoryUserRepository) FindAll() ([]*entities.User, error) {
	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// FindByID returns the user with the specified ID.
func (r *InMemoryUserRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, entities.ErrUserNotFound{}
	}
	return user, nil
}

// Create adds a new user to the repository.
func (r *InMemoryUserRepository) Create(user *entities.User) error {
	if _, ok := r.users[user.ID]; ok {
		return entities.ErrUserAlreadyExists{}
	}
	r.users[user.ID] = user
	return nil
}

// Update updates an existing user in the repository.
func (r *InMemoryUserRepository) Update(user *entities.User) error {
	if _, ok := r.users[user.ID]; !ok {
		return entities.ErrUserNotFound{}
	}
	r.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository.
func (r *InMemoryUserRepository) Delete(id uuid.UUID) error {
	if _, ok := r.users[id]; !ok {
		return entities.ErrUserNotFound{}
	}
	delete(r.users, id)
	return nil
}
