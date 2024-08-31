package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserConfig struct {
	ID         uuid.UUID `json:"id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	SystemRole string    `json:"system_role"`
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

	if config.Password != "" {
		passwordBytes, err := bcrypt.GenerateFromPassword([]byte(config.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		config.Password = string(passwordBytes)
	}

	return &User{
		UserConfig: config,
	}, nil
}
