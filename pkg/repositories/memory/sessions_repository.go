package memory

import (
	"github.com/fandujar/choregate/pkg/repositories"
)

type InMemorySessionsRepository struct {
	sessions map[string]repositories.Session
}

func NewInMemorySessionsRepository() *InMemorySessionsRepository {
	return &InMemorySessionsRepository{
		sessions: make(map[string]repositories.Session),
	}
}

func (r *InMemorySessionsRepository) CreateSession(session repositories.Session) error {
	r.sessions[session.ID] = session
	return nil
}

func (r *InMemorySessionsRepository) GetSession(id string) (repositories.Session, error) {
	session, ok := r.sessions[id]
	if !ok {
		return repositories.Session{}, repositories.ErrSessionNotFound{}
	}
	return session, nil
}

func (r *InMemorySessionsRepository) DeleteSession(id string) error {
	delete(r.sessions, id)
	return nil
}
