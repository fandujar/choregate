package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/jwtauth/v5"
)

type AuthProvider interface {
	// HandleLogin is a method that will be implemented by the auth provider
	HandleLogin(ctx context.Context, username, password string) (string, error)
	// ValidateToken is a method that will be implemented by the auth provider
	ValidateToken(ctx context.Context, token string) (bool, error)
	// ValidateUserPassword is a method that will be implemented by the auth provider
	ValidateUserPassword(ctx context.Context, username, password string) (role string, valid bool, err error)
	// RefreshToken is a method that will be implemented by the auth provider
	RefreshToken(ctx context.Context, token string) (string, error)
	// GenerateToken is a method that will be implemented by the auth provider
	GenerateToken(ctx context.Context, username, role string) (string, error)
	// NewTokenAuth is a method that will be implemented by the auth provider
	NewTokenAuth() *jwtauth.JWTAuth
}

// AuthProviderImpl is the default implementation of the AuthProvider interface
type AuthProviderImpl struct {
	JWTSecret   string
	JWTAuth     *jwtauth.JWTAuth
	UserService *services.UserService
}

// NewAuthProvider creates a new AuthProvider
func NewAuthProvider(userService *services.UserService) (AuthProvider, error) {
	secret := os.Getenv("CHOREGATE_JWT_SECRET")
	if secret == "" {
		return nil, errors.New("missing CHOREGATE_JWT_SECRET environment variable")
	}

	return &AuthProviderImpl{
		JWTSecret:   secret,
		UserService: userService,
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
func (a *AuthProviderImpl) HandleLogin(ctx context.Context, username, password string) (token string, err error) {
	role, valid, err := a.ValidateUserPassword(ctx, username, password)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", errors.New("invalid username or password")
	}
	token, err = a.GenerateToken(ctx, username, role)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateToken(ctx context.Context, token string) (bool, error) {
	t, err := a.NewTokenAuth().Decode(token)
	if err != nil {
		return false, err
	}

	if t.Expiration().Before(time.Now()) {
		return false, nil
	}

	return true, nil
}

// ValidateUserPassword is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) ValidateUserPassword(ctx context.Context, username, password string) (role string, valid bool, err error) {
	if username == "" || password == "" {
		return "", false, nil
	}
	superUser, superUserPassword := os.Getenv("CHOREGATE_SUPERUSER"), os.Getenv("CHOREGATE_SUPERUSER_PASSWORD")
	if superUser != "" && superUserPassword != "" {
		if username == superUser && password == superUserPassword {
			return "admin", true, nil
		}
	}

	user, err := a.UserService.GetUserByEmail(ctx, username)
	if err != nil {
		return "", false, err
	}

	if user == nil {
		return "", false, nil
	}

	if user.Password == password {
		return user.SystemRole, true, nil
	}

	return "", false, nil
}

// RefreshToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) RefreshToken(ctx context.Context, token string) (string, error) {
	return "", nil
}

// GetToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) GenerateToken(ctx context.Context, username, systemRole string) (string, error) {
	_, token, err := a.NewTokenAuth().Encode(
		map[string]interface{}{
			"username":    username,
			"email":       username,
			"system_role": systemRole,
			"exp":         time.Now().Add(time.Hour * 24).Unix(),
			"iat":         time.Now().Unix(),
			"iss":         "choregate",
		},
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
