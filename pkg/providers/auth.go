package providers

import (
	"net/http"
)

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
func NewAuthProvider() (AuthProvider, error) {
	return &AuthProviderImpl{}, nil
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
	if token == "" {
		return false, nil
	}

	if token != a.Token {
		return false, nil
	}

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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
