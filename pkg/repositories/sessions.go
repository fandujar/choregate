package repositories

import "time"

type Session struct {
	// ID is the id of the session
	ID string
	// UserID is the id of the user
	UserID string
	// Token is the token of the session
	Token string
	// ExpiresAt is the expiration time of the session
	ExpiresAt time.Duration
}

type ErrSessionNotFound struct{}

func (e ErrSessionNotFound) Error() string {
	return "session not found"
}

// SessionsRepository is an interface that will be implemented by the sessions repository
type SessionsRepository interface {
	// CreateSession is a method that will be implemented by the sessions repository
	CreateSession(session Session) error
	// GetSession is a method that will be implemented by the sessions repository
	GetSession(id string) (Session, error)
	// DeleteSession is a method that will be implemented by the sessions repository
	DeleteSession(id string) error
}
