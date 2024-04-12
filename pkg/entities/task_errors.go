package entities

// ErrTaskAlreadyExists is an error that is returned when a task already exists.
type ErrTaskAlreadyExists struct{}

func (e ErrTaskAlreadyExists) Error() string {
	return "task already exists"
}

// ErrTaskNotFound is an error that is returned when a task is not found.
type ErrTaskNotFound struct{}

func (e ErrTaskNotFound) Error() string {
	return "task not found"
}

// ErrTaskInvalid is an error that is returned when a task is invalid.
type ErrTaskInvalid struct{}

func (e ErrTaskInvalid) Error() string {
	return "task is invalid"
}

// ErrTaskTitleRequired is an error that is returned when a task title is required.
type ErrTaskTitleRequired struct{}

func (e ErrTaskTitleRequired) Error() string {
	return "task title is required"
}
