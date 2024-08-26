package memory

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

type InMemoryOrganizationRepository struct {
	organizations map[uuid.UUID]*entities.Organization
}

func NewInMemoryOrganizationRepository() *InMemoryOrganizationRepository {
	return &InMemoryOrganizationRepository{
		organizations: make(map[uuid.UUID]*entities.Organization),
	}
}

func (r *InMemoryOrganizationRepository) FindAll(ctx context.Context) ([]*entities.Organization, error) {
	organizations := make([]*entities.Organization, 0, len(r.organizations))
	for _, organization := range r.organizations {
		organizations = append(organizations, organization)
	}
	return organizations, nil
}

func (r *InMemoryOrganizationRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Organization, error) {
	organization, ok := r.organizations[id]
	if !ok {
		return nil, entities.ErrOrganizationNotFound{}
	}
	return organization, nil
}

func (r *InMemoryOrganizationRepository) Create(ctx context.Context, organization *entities.Organization) error {
	if _, ok := r.organizations[organization.ID]; ok {
		return entities.ErrOrganizationAlreadyExists{}
	}
	r.organizations[organization.ID] = organization
	return nil
}

func (r *InMemoryOrganizationRepository) Update(ctx context.Context, organization *entities.Organization) error {
	if _, ok := r.organizations[organization.ID]; !ok {
		return entities.ErrOrganizationNotFound{}
	}
	r.organizations[organization.ID] = organization
	return nil
}

func (r *InMemoryOrganizationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, ok := r.organizations[id]; !ok {
		return entities.ErrOrganizationNotFound{}
	}
	delete(r.organizations, id)
	return nil
}

func (r *InMemoryOrganizationRepository) AddMember(ctx context.Context, organizationID, userID uuid.UUID, role string) error {
	var err error
	organization, ok := r.organizations[organizationID]
	if !ok {
		return entities.ErrOrganizationNotFound{}
	}
	organization.Members[userID], err = entities.NewOrganizationMember(
		&entities.OrganizationMemberConfig{
			OrgID:  organizationID,
			UserID: userID,
			Role:   role,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *InMemoryOrganizationRepository) RemoveMember(ctx context.Context, organizationID, userID uuid.UUID) error {
	organization, ok := r.organizations[organizationID]
	if !ok {
		return entities.ErrOrganizationNotFound{}
	}
	delete(organization.Members, userID)
	return nil
}

func (r *InMemoryOrganizationRepository) FindMembers(ctx context.Context, organizationID uuid.UUID) ([]*entities.OrganizationMember, error) {
	organization, ok := r.organizations[organizationID]
	if !ok {
		return nil, entities.ErrOrganizationNotFound{}
	}
	members := make([]*entities.OrganizationMember, 0, len(organization.Members))
	for _, member := range organization.Members {
		members = append(members, member)
	}
	return members, nil
}

func (r *InMemoryOrganizationRepository) FindMember(ctx context.Context, organizationID, userID uuid.UUID) (*entities.OrganizationMember, error) {
	organization, ok := r.organizations[organizationID]
	if !ok {
		return nil, entities.ErrOrganizationNotFound{}
	}
	member, ok := organization.Members[userID]
	if !ok {
		return nil, entities.ErrOrganizationMemberNotFound{}
	}
	return member, nil
}

func (r *InMemoryOrganizationRepository) FindMemberMemberships(ctx context.Context, userID uuid.UUID) ([]*entities.OrganizationMember, error) {
	memberships := make([]*entities.OrganizationMember, 0)
	for _, organization := range r.organizations {
		member, ok := organization.Members[userID]
		if ok {
			memberships = append(memberships, member)
		}
	}
	return memberships, nil
}

func (r *InMemoryOrganizationRepository) UpdateMemberRole(ctx context.Context, organizationID, userID uuid.UUID, role string) error {
	organization, ok := r.organizations[organizationID]
	if !ok {
		return entities.ErrOrganizationNotFound{}
	}
	member, ok := organization.Members[userID]
	if !ok {
		return entities.ErrOrganizationMemberNotFound{}
	}
	member.Role = role
	return nil
}

func (r *InMemoryOrganizationRepository) AddTeam(ctx context.Context, organizationID, teamID uuid.UUID) error {
	var err error
	organization, ok := r.organizations[organizationID]
	if !ok {
		return entities.ErrOrganizationNotFound{}
	}

	organization.Teams[teamID], err = entities.NewOrganizationTeam(
		&entities.OrganizationTeamConfig{
			OrgID:  organizationID,
			TeamID: teamID,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *InMemoryOrganizationRepository) RemoveTeam(ctx context.Context, organizationID, teamID uuid.UUID) error {
	organization, ok := r.organizations[organizationID]
	if !ok {
		return entities.ErrOrganizationNotFound{}
	}
	delete(organization.Teams, teamID)
	return nil
}
