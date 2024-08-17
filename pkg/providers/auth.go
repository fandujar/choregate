package providers

import (
	"errors"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

type AuthProvider interface {
	// HandleLogin is a method that will be implemented by the auth provider
	HandleLogin(username, password string) (string, error)
	// ValidateToken is a method that will be implemented by the auth provider
	ValidateToken(token string) (bool, error)
	// ValidateUserPassword is a method that will be implemented by the auth provider
	ValidateUserPassword(username, password string) (role string, valid bool, err error)
	// RefreshToken is a method that will be implemented by the auth provider
	RefreshToken(token string) (string, error)
	// GenerateToken is a method that will be implemented by the auth provider
	GenerateToken(username, role string) (string, error)
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

// HandleLogin is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) HandleLogin(username, password string) (string, error) {
	role, valid, err := a.ValidateUserPassword(username, password)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", errors.New("invalid username or password")
	}
	token, err := a.GenerateToken(username, role)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateToken(token string) (bool, error) {
	_, err := a.NewTokenAuth().Decode(token)
	if err != nil {
		return false, err
	}

	return true, nil
}

// ValidateUserPassword is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateUserPassword(username, password string) (role string, valid bool, err error) {
	if username == "" || password == "" {
		return "", false, nil
	}

	return "user", true, nil
}

// RefreshToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) RefreshToken(token string) (string, error) {
	return "", nil
}

// GetToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) GenerateToken(username, role string) (string, error) {
	_, token, err := a.NewTokenAuth().Encode(
		map[string]interface{}{
			"username": username,
			"role":     role,
		},
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
