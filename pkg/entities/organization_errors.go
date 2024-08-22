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
