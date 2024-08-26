package services

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

// NewUserService creates a new user service.
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser returns a user by ID.
func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers returns all users.
func (s *UserService) GetUsers(ctx context.Context) ([]*entities.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *entities.User) error {
	err := s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user.
func (s *UserService) UpdateUser(ctx context.Context, user *entities.User) error {
	err := s.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail returns a user by email.
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var user *entities.User

	for _, u := range users {
		if u.Email == email {
			user = u
			break
		}
	}

	if user == nil {
		return nil, entities.ErrUserNotFound{}
	}

	return user, nil
}
