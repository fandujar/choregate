package entities

// ErrTeamNotFound is an error that is returned when a team is not found.
type ErrTeamNotFound struct{}

func (e ErrTeamNotFound) Error() string {
	return "team not found"
}

// ErrTeamAlreadyExists is an error that is returned when a team already exists.
type ErrTeamAlreadyExists struct{}

func (e ErrTeamAlreadyExists) Error() string {
	return "team already exists"
}

// ErrTeamInvalid is an error that is returned when a team is invalid.
type ErrTeamInvalid struct{}

func (e ErrTeamInvalid) Error() string {
	return "team is invalid"
}

// ErrMemberNotFound is an error that is returned when a member is not found.
type ErrMemberNotFound struct{}

func (e ErrMemberNotFound) Error() string {
	return "member not found"
}

// ErrMemberAlreadyExists is an error that is returned when a member already exists.
type ErrMemberAlreadyExists struct{}

func (e ErrMemberAlreadyExists) Error() string {
	return "member already exists"
}

// ErrMemberInvalid is an error that is returned when a member is invalid.
type ErrMemberInvalid struct{}

func (e ErrMemberInvalid) Error() string {
	return "member is invalid"
}
