package memory

import (
	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

type InMemorySessionsRepository struct {
	sessions map[uuid.UUID]entities.Session
}

func NewInMemorySessionsRepository() *InMemorySessionsRepository {
	return &InMemorySessionsRepository{
		sessions: make(map[uuid.UUID]entities.Session),
	}
}

func (r *InMemorySessionsRepository) CreateSession(session entities.Session) error {
	r.sessions[session.ID] = session
	return nil
}

func (r *InMemorySessionsRepository) GetSession(id uuid.UUID) (entities.Session, error) {
	session, ok := r.sessions[id]
	if !ok {
		return entities.Session{}, entities.ErrSessionNotFound{}
	}
	return session, nil
}

func (r *InMemorySessionsRepository) DeleteSession(id uuid.UUID) error {
	delete(r.sessions, id)
	return nil
}
