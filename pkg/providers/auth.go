package providers

type AuthProvider interface {
	// Login is a method that will be implemented by the auth provider
	Login(username, password string) (string, error)
	// Logout is a method that will be implemented by the auth provider
	Logout(token string) error
	// ValidateToken is a method that will be implemented by the auth provider
	ValidateToken(token string) (bool, error)
	// RefreshToken is a method that will be implemented by the auth provider
	RefreshToken(token string) (string, error)
	// GetToken is a method that will be implemented by the auth provider
	GetToken() string
}

// AuthProviderImpl is the default implementation of the AuthProvider interface
type AuthProviderImpl struct {
	// Username is the username of the user
	Username string
	// Password is the password of the user
	Password string
	// Token is the token of the user
	Token string
}

// NewAuthProvider creates a new AuthProvider
func NewAuthProvider(username, password string) AuthProvider {
	return &AuthProviderImpl{
		Username: username,
		Password: password,
	}
}

// Login is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) Login(username, password string) (string, error) {
	return "", nil
}

// Logout is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) Logout(token string) error {
	return nil
}

// ValidateToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateToken(token string) (bool, error) {
	return true, nil
}

// RefreshToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) RefreshToken(token string) (string, error) {
	return "", nil
}

// GetToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) GetToken() string {
	return ""
}
