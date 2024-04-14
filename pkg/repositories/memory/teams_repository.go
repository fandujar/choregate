package memory

import (
	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// InMemoryTeamRepository is an in-memory repository that manages teams.
type InMemoryTeamRepository struct {
	teams map[uuid.UUID]*entities.Team
}

// FindAll returns all teams in the repository.
func (r *InMemoryTeamRepository) FindAll() ([]*entities.Team, error) {
	teams := make([]*entities.Team, 0, len(r.teams))
	for _, team := range r.teams {
		teams = append(teams, team)
	}
	return teams, nil
}

// FindByID returns the team with the specified ID.
func (r *InMemoryTeamRepository) FindByID(id uuid.UUID) (*entities.Team, error) {
	team, ok := r.teams[id]
	if !ok {
		return nil, entities.ErrTeamNotFound{}
	}
	return team, nil
}

// Create adds a new team to the repository.
func (r *InMemoryTeamRepository) Create(team *entities.Team) error {
	if _, ok := r.teams[team.ID]; ok {
		return entities.ErrTeamAlreadyExists{}
	}
	r.teams[team.ID] = team
	return nil
}

// Update updates an existing team in the repository.
func (r *InMemoryTeamRepository) Update(team *entities.Team) error {
	if _, ok := r.teams[team.ID]; !ok {
		return entities.ErrTeamNotFound{}
	}
	r.teams[team.ID] = team
	return nil
}

// Delete removes a team from the repository.
func (r *InMemoryTeamRepository) Delete(id uuid.UUID) error {
	if _, ok := r.teams[id]; !ok {
		return entities.ErrTeamNotFound{}
	}
	delete(r.teams, id)
	return nil
}

// AddMember adds a member to a team.
func (r *InMemoryTeamRepository) AddMember(teamID, userID uuid.UUID, role string) error {
	team, ok := r.teams[teamID]
	if !ok {
		return entities.ErrTeamNotFound{}
	}

	if _, ok := team.Members[userID]; ok {
		return entities.ErrMemberAlreadyExists{}
	}

	team.Members[userID] = &entities.Member{
		UserID: userID,
		Role:   role,
	}

	return nil
}

// RemoveMember removes a member from a team.
func (r *InMemoryTeamRepository) RemoveMember(teamID, userID uuid.UUID) error {
	team, ok := r.teams[teamID]
	if !ok {
		return entities.ErrTeamNotFound{}
	}

	if _, ok := team.Members[userID]; !ok {
		return entities.ErrMemberNotFound{}
	}

	delete(team.Members, userID)
	return nil
}

// UpdateMemberRole updates the role of a member in a team.
func (r *InMemoryTeamRepository) UpdateMemberRole(teamID, userID uuid.UUID, role string) error {
	team, ok := r.teams[teamID]
	if !ok {
		return entities.ErrTeamNotFound{}
	}

	member, ok := team.Members[userID]
	if !ok {
		return entities.ErrMemberNotFound{}
	}

	member.Role = role
	return nil
}
