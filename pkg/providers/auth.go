package providers

import (
	"errors"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

type AuthProvider interface {
	// ValidateToken is a method that will be implemented by the auth provider
	ValidateToken(token string) (bool, error)
	// ValidateUserPassword is a method that will be implemented by the auth provider
	ValidateUserPassword(username, password string) (bool, error)
	// RefreshToken is a method that will be implemented by the auth provider
	RefreshToken(token string) (string, error)
	// GenerateToken is a method that will be implemented by the auth provider
	GenerateToken(username string) (string, error)
	// NewTokenAuth is a method that will be implemented by the auth provider
	NewTokenAuth() *jwtauth.JWTAuth
}

// AuthProviderImpl is the default implementation of the AuthProvider interface
type AuthProviderImpl struct {
	JWTSecret string
	JWTAuth   *jwtauth.JWTAuth
}

// NewAuthProvider creates a new AuthProvider
func NewAuthProvider() (AuthProvider, error) {
	secret := os.Getenv("CHOREGATE_JWT_SECRET")
	if secret == "" {
		return nil, errors.New("missing CHOREGATE_JWT_SECRET environment variable")
	}

	return &AuthProviderImpl{
		JWTSecret: secret,
	}, nil
}

// NewTokenAuth is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) NewTokenAuth() *jwtauth.JWTAuth {
	if a.JWTAuth == nil {
		a.JWTAuth = jwtauth.New("HS256", []byte(a.JWTSecret), nil)
	}
	return a.JWTAuth
}

// ValidateToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateToken(token string) (bool, error) {
	return true, nil
}

// ValidateUserPassword is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateUserPassword(username, password string) (bool, error) {
	if username == "" || password == "" {
		return false, nil
	}
	return true, nil
}

// RefreshToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) RefreshToken(token string) (string, error) {
	return "", nil
}

// GetToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) GenerateToken(username string) (string, error) {
	_, token, err := a.NewTokenAuth().Encode(map[string]interface{}{"username": username})
	if err != nil {
		return "", err
	}

	return token, nil
}
