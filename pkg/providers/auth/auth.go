package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/jwtauth/v5"
)

type AuthProvider interface {
	// HandleLogin is a method that will be implemented by the auth provider
	HandleLogin(ctx context.Context, username, password string) (*entities.User, string, error)
	// ValidateToken is a method that will be implemented by the auth provider
	ValidateToken(ctx context.Context, token string) (bool, error)
	// ValidateUserPassword is a method that will be implemented by the auth provider
	ValidateUserPassword(ctx context.Context, username, password string) (user *entities.User, valid bool, err error)
	// RefreshToken is a method that will be implemented by the auth provider
	RefreshToken(ctx context.Context, token string) (string, error)
	// GenerateToken is a method that will be implemented by the auth provider
	GenerateToken(ctx context.Context, user *entities.User) (string, error)
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
func (a *AuthProviderImpl) HandleLogin(ctx context.Context, username, password string) (user *entities.User, token string, err error) {
	user, valid, err := a.ValidateUserPassword(ctx, username, password)
	if err != nil {
		return nil, "", err
	}
	if !valid {
		return nil, "", errors.New("invalid username or password")
	}
	token, err = a.GenerateToken(ctx, user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
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
func (a *AuthProviderImpl) ValidateUserPassword(ctx context.Context, username, password string) (user *entities.User, valid bool, err error) {
	if username == "" || password == "" {
		return nil, false, nil
	}
	superUser, superUserPassword := os.Getenv("CHOREGATE_SUPERUSER"), os.Getenv("CHOREGATE_SUPERUSER_PASSWORD")
	if superUser != "" && superUserPassword != "" {
		if username == superUser && password == superUserPassword {
			user, _ = entities.NewUser(
				&entities.UserConfig{
					Slug:       superUser,
					Name:       superUser,
					Email:      superUser,
					Password:   superUserPassword,
					SystemRole: "admin",
				},
			)
			return user, true, nil
		}
	}

	user, err = a.UserService.GetUserByEmail(ctx, username)
	if err != nil {
		return nil, false, err
	}

	if user == nil {
		return nil, false, nil
	}

	if user.Password == password {
		return user, true, nil
	}

	return nil, false, nil
}

// RefreshToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) RefreshToken(ctx context.Context, token string) (string, error) {
	return "", nil
}

// GetToken is a method that will be implemented by the auth provider
func (a *AuthProviderImpl) GenerateToken(ctx context.Context, user *entities.User) (string, error) {
	_, token, err := a.NewTokenAuth().Encode(
		map[string]interface{}{
			"username":    user.Email,
			"user_id":     user.ID.String(),
			"email":       user.Email,
			"system_role": user.SystemRole,
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
