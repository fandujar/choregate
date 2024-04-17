package services

import (
	"context"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/google/uuid"
)

type TeamService struct {
	repo repositories.TeamRepository
}

func NewTeamService(repo repositories.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) GetTeam(ctx context.Context, id uuid.UUID) (*entities.Team, error) {
	team, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *TeamService) GetTeams(ctx context.Context) ([]*entities.Team, error) {
	teams, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *TeamService) CreateTeam(ctx context.Context, team *entities.Team) error {
	err := s.repo.Create(ctx, team)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) UpdateTeam(ctx context.Context, team *entities.Team) error {
	err := s.repo.Update(ctx, team)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) AddMember(ctx context.Context, teamID, userID uuid.UUID, role string) error {
	err := s.repo.AddMember(ctx, teamID, userID, role)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) RemoveMember(ctx context.Context, teamID, userID uuid.UUID) error {
	err := s.repo.RemoveMember(ctx, teamID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TeamService) UpdateMemberRole(ctx context.Context, teamID, userID uuid.UUID, role string) error {
	err := s.repo.UpdateMemberRole(ctx, teamID, userID, role)
	if err != nil {
		return err
	}

	return nil
}
