package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

type TeamConfig struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Team struct {
	*TeamConfig
	Members map[uuid.UUID]*Member `json:"members"`
}

type Member struct {
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

	return &Team{
		TeamConfig: config,
		Members:    make(map[uuid.UUID]*Member),
	}, nil
}
