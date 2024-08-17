package entities

import (
	"time"

	"github.com/google/uuid"
)

type SessionConfig struct {
	// ID is the id of the session
	ID uuid.UUID
	// UserID is the id of the user
	UserID string
	// Token is the token of the session
	Token string
	// ExpiresAt is the expiration time of the session
	ExpiresAt time.Duration
}

type Session struct {
	*SessionConfig
}

func NewSession(config *SessionConfig) *Session {
	return &Session{
		SessionConfig: config,
	}
}
