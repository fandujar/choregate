package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

type UserConfig struct {
	ID    uuid.UUID `json:"id"`
	Slug  string    `json:"slug"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type User struct {
	*UserConfig
}

// NewUser creates a new User with the given configuration and default values.
func NewUser(config *UserConfig) (*User, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	return &User{
		UserConfig: config,
	}, nil
}
