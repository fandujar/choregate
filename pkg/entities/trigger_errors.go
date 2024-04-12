package entities

// ErrTriggerAlreadyExists is an error that is returned when a trigger already exists.
type ErrTriggerAlreadyExists struct{}

func (e ErrTriggerAlreadyExists) Error() string {
	return "trigger already exists"
}

// ErrTriggerNotFound is an error that is returned when a trigger is not found.
type ErrTriggerNotFound struct{}

func (e ErrTriggerNotFound) Error() string {
	return "trigger not found"
}
