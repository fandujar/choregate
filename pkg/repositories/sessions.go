package repositories

import (
	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// SessionsRepository is an interface that will be implemented by the sessions repository
type SessionsRepository interface {
	// CreateSession is a method that will be implemented by the sessions repository
	CreateSession(session entities.Session) error
	// GetSession is a method that will be implemented by the sessions repository
	GetSession(id uuid.UUID) (entities.Session, error)
	// DeleteSession is a method that will be implemented by the sessions repository
	DeleteSession(id uuid.UUID) error
}
