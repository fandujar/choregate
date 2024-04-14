package entities

// ErrUserNotFound is an error that is returned when a user is not found.
type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "user not found"
}

// ErrUserAlreadyExists is an error that is returned when a user already exists.
type ErrUserAlreadyExists struct{}

func (e ErrUserAlreadyExists) Error() string {
	return "user already exists"
}

// ErrUserInvalid is an error that is returned when a user is invalid.
type ErrUserInvalid struct{}

func (e ErrUserInvalid) Error() string {
	return "user is invalid"
}
