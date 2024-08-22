package entities

import "github.com/google/uuid"

type Organization struct {
	*OrganizationConfig
}

type OrganizationConfig struct {
	ID      uuid.UUID             `json:"id"`
	Name    string                `json:"name"`
	Teams   map[uuid.UUID]*Team   `json:"teams"`
	Members map[uuid.UUID]*Member `json:"members"`
}
