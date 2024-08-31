package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

type TeamConfig struct {
	ID      uuid.UUID                 `json:"id"`
	Name    string                    `json:"name"`
	Members map[uuid.UUID]*TeamMember `json:"members"`
}

type Team struct {
	*TeamConfig
}

type TeamMember struct {
	ID     uuid.UUID `json:"id"`
	TeamID uuid.UUID `json:"team_id"`
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

func NewTeam(config *TeamConfig) (*Team, error) {
	var err error
	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	if config.Members == nil {
		config.Members = make(map[uuid.UUID]*TeamMember)
	}

	return &Team{
		TeamConfig: config,
	}, nil
}
