package services

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
)

type OrganizationService struct {
	repo repositories.OrganizationRepository
}

// NewOrganizationService creates a new organization service.
func NewOrganizationService(repo repositories.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

// GetOrganizations returns all organizations.
func (s *OrganizationService) GetOrganizations(ctx context.Context) ([]*entities.Organization, error) {
	return s.repo.FindAll(ctx)
}

// GetOrganization returns an organization by ID.
func (s *OrganizationService) GetOrganization(ctx context.Context, id uuid.UUID) (*entities.Organization, error) {
	return s.repo.FindByID(ctx, id)
}

// CreateOrganization creates a new organization.
func (s *OrganizationService) CreateOrganization(ctx context.Context, organization *entities.Organization) error {
	return s.repo.Create(ctx, organization)
}

// UpdateOrganization updates an organization.
func (s *OrganizationService) UpdateOrganization(ctx context.Context, organization *entities.Organization) error {
	return s.repo.Update(ctx, organization)
}

// DeleteOrganization deletes an organization by ID.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// AddMember adds a member to an organization.
func (s *OrganizationService) AddMember(ctx context.Context, organizationID, userID uuid.UUID, role string) error {
	return s.repo.AddMember(ctx, organizationID, userID, role)
}

// RemoveMember removes a member from an organization.
func (s *OrganizationService) RemoveMember(ctx context.Context, organizationID, userID uuid.UUID) error {
	return s.repo.RemoveMember(ctx, organizationID, userID)
}

// UpdateMemberRole updates the role of a member in an organization.
func (s *OrganizationService) UpdateMemberRole(ctx context.Context, organizationID, userID uuid.UUID, role string) error {
	return s.repo.UpdateMemberRole(ctx, organizationID, userID, role)
}

// AddTeam adds a team to an organization.
func (s *OrganizationService) AddTeam(ctx context.Context, organizationID, teamID uuid.UUID) error {
	return s.repo.AddTeam(ctx, organizationID, teamID)
}

// RemoveTeam removes a team from an organization.
func (s *OrganizationService) RemoveTeam(ctx context.Context, organizationID, teamID uuid.UUID) error {
	return s.repo.RemoveTeam(ctx, organizationID, teamID)
}
