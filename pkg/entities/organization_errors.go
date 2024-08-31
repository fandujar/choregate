package entities

// ErrOrganizationNotFound is an error that is returned when a organization is not found
type ErrOrganizationNotFound struct{}

func (e ErrOrganizationNotFound) Error() string {
	return "organization not found"
}

// ErrInvalidOrganization is an error that is returned when a organization is invalid
type ErrInvalidOrganization struct{}

func (e ErrInvalidOrganization) Error() string {
	return "invalid organization"
}

// ErrOrganizationAlreadyExists is an error that is returned when a organization already exists
type ErrOrganizationAlreadyExists struct{}

func (e ErrOrganizationAlreadyExists) Error() string {
	return "organization already exists"
}

// ErrOrganizationMemberNotFound is an error that is returned when a organization member is not found
type ErrOrganizationMemberNotFound struct{}

func (e ErrOrganizationMemberNotFound) Error() string {
	return "organization member not found"
}
