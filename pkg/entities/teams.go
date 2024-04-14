package entities

import "github.com/google/uuid"

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
