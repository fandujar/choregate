package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

type Organization struct {
	*OrganizationConfig
}

type OrganizationConfig struct {
	ID      uuid.UUID                         `json:"id"`
	Name    string                            `json:"name"`
	Teams   map[uuid.UUID]*OrganizationTeam   `json:"teams"`
	Members map[uuid.UUID]*OrganizationMember `json:"members"`
}

type OrganizationMember struct {
	*OrganizationMemberConfig
}

type OrganizationMemberConfig struct {
	ID     uuid.UUID `json:"id"`
	OrgID  uuid.UUID `json:"org_id"`
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

type OrganizationTeam struct {
	*OrganizationTeamConfig
}

type OrganizationTeamConfig struct {
	ID     uuid.UUID `json:"id"`
	OrgID  uuid.UUID `json:"org_id"`
	TeamID uuid.UUID `json:"team_id"`
}

func NewOrganization(config *OrganizationConfig) (*Organization, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	if config.Teams == nil {
		config.Teams = make(map[uuid.UUID]*OrganizationTeam)
	}

	if config.Members == nil {
		config.Members = make(map[uuid.UUID]*OrganizationMember)
	}

	return &Organization{OrganizationConfig: config}, nil
}

func NewOrganizationMember(config *OrganizationMemberConfig) (*OrganizationMember, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	return &OrganizationMember{OrganizationMemberConfig: config}, nil
}

func NewOrganizationTeam(config *OrganizationTeamConfig) (*OrganizationTeam, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	return &OrganizationTeam{OrganizationTeamConfig: config}, nil
}
