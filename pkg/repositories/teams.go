package repositories

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// TeamRepository is a repository that manages teams.
type TeamRepository interface {
	// FindAll returns all teams.
	FindAll(ctx context.Context) ([]*entities.Team, error)
	// FindByID returns a team by ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Team, error)
	// Create creates a new team.
	Create(ctx context.Context, team *entities.Team) error
	// Update updates a team.
	Update(ctx context.Context, team *entities.Team) error
	// Delete deletes a team by ID.
	Delete(ctx context.Context, id uuid.UUID) error
	// AddMember adds a member to a team.
	AddMember(ctx context.Context, teamID, userID uuid.UUID, role string) error
	// RemoveMember removes a member from a team.
	RemoveMember(ctx context.Context, teamID, userID uuid.UUID) error
	// UpdateMemberRole updates the role of a member in a team.
	UpdateMemberRole(ctx context.Context, teamID, userID uuid.UUID, role string) error
}
