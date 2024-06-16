package memory

import (
	"github.com/fandujar/choregate/pkg/entities"
)

type InMemorySessionsRepository struct {
	sessions map[string]entities.Session
}

func NewInMemorySessionsRepository() *InMemorySessionsRepository {
	return &InMemorySessionsRepository{
		sessions: make(map[string]entities.Session),
	}
}

func (r *InMemorySessionsRepository) CreateSession(session entities.Session) error {
	r.sessions[session.ID] = session
	return nil
}

func (r *InMemorySessionsRepository) GetSession(id string) (entities.Session, error) {
	session, ok := r.sessions[id]
	if !ok {
		return entities.Session{}, entities.ErrSessionNotFound{}
	}
	return session, nil
}

func (r *InMemorySessionsRepository) DeleteSession(id string) error {
	delete(r.sessions, id)
	return nil
}
