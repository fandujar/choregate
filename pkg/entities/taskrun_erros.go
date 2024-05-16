package entities

// ErrTaskRunAlreadyExists is an error that is returned when a task run already exists.
type ErrTaskRunAlreadyExists struct{}

func (e ErrTaskRunAlreadyExists) Error() string {
	return "task run already exists"
}

// ErrTaskRunNotFound is an error that is returned when a task run is not found.
type ErrTaskRunNotFound struct{}

func (e ErrTaskRunNotFound) Error() string {
	return "task run not found"
}

// ErrTaskRunInvalid is an error that is returned when a task run is invalid.
type ErrTaskRunInvalid struct{}

func (e ErrTaskRunInvalid) Error() string {
	return "task run is invalid"
}

// ErrTaskRunFailed is an error that is returned when a task run has failed.
type ErrTaskRunFailed struct{}

func (e ErrTaskRunFailed) Error() string {
	return "task run has failed"
}
