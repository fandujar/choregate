package providers

import (
	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

type SessionRequest struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Token    string    `json:"token"`
}

type SessionProvider interface {
	// CreateSession is a method that will be implemented by the session provider
	CreateSession(request SessionRequest) error
	// GetSession is a method that will be implemented by the session provider
	GetSession(id string) (entities.Session, error)
	// DeleteSession is a method that will be implemented by the session provider
	DeleteSession(id string) error
}

type SessionProviderImpl struct {
	Repository repositories.SessionsRepository
}

func NewSessionProvider(repository repositories.SessionsRepository) (SessionProvider, error) {
	return &SessionProviderImpl{
		Repository: repository,
	}, nil
}

func (s *SessionProviderImpl) CreateSession(request SessionRequest) error {
	var session entities.Session
	var err error
	if request.ID == uuid.Nil {
		session.ID, err = utils.GenerateID()
		if err != nil {
			return err
		}
	}

	if request.Username == "" || request.Password == "" {
		return entities.ErrInvalidSession{}
	}

	auth, err := NewAuthProvider(
		request.Username,
		request.Password,
		request.Token,
	)

	if err != nil {
		return err
	}

	valid, err := auth.ValidateUserPassword(request.Username, request.Password)
	if err != nil {
		return err
	}

	if !valid {
		return entities.ErrInvalidSession{}
	}

	return nil
}

func (s *SessionProviderImpl) GetSession(id string) (entities.Session, error) {
	uid, err := uuid.Parse(id)

	if err != nil {
		return entities.Session{}, err
	}

	return s.Repository.GetSession(uid)
}

func (s *SessionProviderImpl) DeleteSession(id string) error {
	return s.Repository.DeleteSession(id)
}
