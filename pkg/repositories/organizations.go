package repositories

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

type OrganizationRepository interface {
	// FindAll returns all organizations.
	FindAll(ctx context.Context) ([]*entities.Organization, error)
	// FindByID returns an organization by ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Organization, error)
	// Create creates a new organization.
	Create(ctx context.Context, organization *entities.Organization) error
	// Update updates an organization.
	Update(ctx context.Context, organization *entities.Organization) error
	// Delete deletes an organization by ID.
	Delete(ctx context.Context, id uuid.UUID) error
	// AddMember adds a member to an organization.
	AddMember(ctx context.Context, organizationID, userID uuid.UUID, role string) error
	// RemoveMember removes a member from an organization.
	RemoveMember(ctx context.Context, organizationID, userID uuid.UUID) error
	// UpdateMemberRole updates the role of a member in an organization.
	UpdateMemberRole(ctx context.Context, organizationID, userID uuid.UUID, role string) error
	// AddTeam adds a team to an organization.
	AddTeam(ctx context.Context, organizationID, teamID uuid.UUID) error
	// RemoveTeam removes a team from an organization.
	RemoveTeam(ctx context.Context, organizationID, teamID uuid.UUID) error
}
