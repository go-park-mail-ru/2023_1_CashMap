package local_storage

import (
	session "depeche/internal/session/repository"
	"errors"
	"sync"
)

type SessionStorage struct {
	mx      sync.RWMutex
	session map[string]*session.Session
}

func NewMemorySessionStorage() *SessionStorage {
	return &SessionStorage{
		session: map[string]*session.Session{},
	}
}

func (s *SessionStorage) CreateSession(token string, session *session.Session) error {
	s.mx.RLock()
	defer s.mx.RUnlock()
	s.session[token] = session
	return nil
}

func (s *SessionStorage) GetSession(token string) (*session.Session, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	stored := s.session[token]
	if stored == nil {
		return nil, errors.New("session not found")
	}
	return stored, nil
}

func (s *SessionStorage) DeleteSession(token string) error {
	s.mx.RLock()
	defer s.mx.RUnlock()
	stored := s.session[token]
	if stored == nil {
		return errors.New("session not found")
	}
	delete(s.session, token)
	return nil
}
