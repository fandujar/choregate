package services

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
)

type OrganizationService struct {
	organizationRepo repositories.OrganizationRepository
	teamRepo         repositories.TeamRepository
	userRepo         repositories.UserRepository
}

// NewOrganizationService creates a new organization service.
func NewOrganizationService(organizationRepo repositories.OrganizationRepository, teamRepo repositories.TeamRepository, userRepo repositories.UserRepository) *OrganizationService {
	return &OrganizationService{
		organizationRepo: organizationRepo,
		teamRepo:         teamRepo,
		userRepo:         userRepo,
	}
}

// GetOrganizations returns all organizations.
func (s *OrganizationService) GetOrganizations(ctx context.Context) ([]*entities.Organization, error) {
	return s.organizationRepo.FindAll(ctx)
}

// GetOrganization returns an organization by ID.
func (s *OrganizationService) GetOrganization(ctx context.Context, id uuid.UUID) (*entities.Organization, error) {
	return s.organizationRepo.FindByID(ctx, id)
}

// CreateOrganization creates a new organization.
func (s *OrganizationService) CreateOrganization(ctx context.Context, organization *entities.Organization) error {
	return s.organizationRepo.Create(ctx, organization)
}

// UpdateOrganization updates an organization.
func (s *OrganizationService) UpdateOrganization(ctx context.Context, organization *entities.Organization) error {
	return s.organizationRepo.Update(ctx, organization)
}

// DeleteOrganization deletes an organization by ID.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, id uuid.UUID) error {
	return s.organizationRepo.Delete(ctx, id)
}

// AddMember adds a member to an organization.
func (s *OrganizationService) AddOrganizationMember(ctx context.Context, organizationID, userID uuid.UUID, role string) error {
	return s.organizationRepo.AddMember(ctx, organizationID, userID, role)
}

// RemoveMember removes a member from an organization.
func (s *OrganizationService) RemoveOrganizationMember(ctx context.Context, organizationID, userID uuid.UUID) error {
	return s.organizationRepo.RemoveMember(ctx, organizationID, userID)
}

// UpdateMemberRole updates the role of a member in an organization.
func (s *OrganizationService) UpdateOrganizationMemberRole(ctx context.Context, organizationID, userID uuid.UUID, role string) error {
	return s.organizationRepo.UpdateMemberRole(ctx, organizationID, userID, role)
}

// AddTeam adds a team to an organization.
func (s *OrganizationService) AddTeam(ctx context.Context, organizationID, teamID uuid.UUID) error {
	return s.organizationRepo.AddTeam(ctx, organizationID, teamID)
}

// RemoveTeam removes a team from an organization.
func (s *OrganizationService) RemoveTeam(ctx context.Context, organizationID, teamID uuid.UUID) error {
	return s.organizationRepo.RemoveTeam(ctx, organizationID, teamID)
}

// GetTeam returns a team by ID.
func (s *OrganizationService) GetTeam(ctx context.Context, id uuid.UUID) (*entities.Team, error) {
	team, err := s.teamRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return team, nil
}

// GetTeams returns all teams.
func (s *OrganizationService) GetTeams(ctx context.Context) ([]*entities.Team, error) {
	teams, err := s.teamRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// CreateTeam creates a new team.
func (s *OrganizationService) CreateTeam(ctx context.Context, team *entities.Team) error {
	err := s.teamRepo.Create(ctx, team)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTeam updates a team.
func (s *OrganizationService) UpdateTeam(ctx context.Context, team *entities.Team) error {
	err := s.teamRepo.Update(ctx, team)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTeam deletes a team by ID.
func (s *OrganizationService) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	err := s.teamRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// AddMember adds a member to a team.
func (s *OrganizationService) AddTeamMember(ctx context.Context, teamID, userID uuid.UUID, role string) error {
	err := s.teamRepo.AddMember(ctx, teamID, userID, role)
	if err != nil {
		return err
	}

	return nil
}

// RemoveMember removes a member from a team.
func (s *OrganizationService) RemoveTeamMember(ctx context.Context, teamID, userID uuid.UUID) error {
	err := s.teamRepo.RemoveMember(ctx, teamID, userID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMemberRole updates the role of a member in a team.
func (s *OrganizationService) UpdateTeamMemberRole(ctx context.Context, teamID, userID uuid.UUID, role string) error {
	err := s.teamRepo.UpdateMemberRole(ctx, teamID, userID, role)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user by ID.
func (s *OrganizationService) GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers returns all users.
func (s *OrganizationService) GetUsers(ctx context.Context) ([]*entities.User, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser creates a new user.
func (s *OrganizationService) CreateUser(ctx context.Context, user *entities.User) error {
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user.
func (s *OrganizationService) UpdateUser(ctx context.Context, user *entities.User) error {
	err := s.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user by ID.
func (s *OrganizationService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail returns a user by email.
func (s *OrganizationService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserMemberships returns all memberships of a user.
func (s *OrganizationService) GetUserOrganizationsMemberships(ctx context.Context, userID uuid.UUID) ([]*entities.OrganizationMember, error) {
	return s.organizationRepo.FindMemberMemberships(ctx, userID)
}

// GetUserTeamsMemberships returns all memberships of a user.
func (s *OrganizationService) GetUserTeamsMemberships(ctx context.Context, userID uuid.UUID) ([]*entities.TeamMember, error) {
	return s.teamRepo.FindMemberMemberships(ctx, userID)
}
