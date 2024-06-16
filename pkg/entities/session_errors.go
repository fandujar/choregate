package entities

// ErrSessionNotFound is an error that is returned when a session is not found
type ErrSessionNotFound struct{}

func (e ErrSessionNotFound) Error() string {
	return "session not found"
}

// ErrInvalidSession is an error that is returned when a session is invalid
type ErrInvalidSession struct{}

func (e ErrInvalidSession) Error() string {
	return "invalid session"
}
